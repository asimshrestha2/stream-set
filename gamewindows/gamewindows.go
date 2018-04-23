package gamewindows

import (
	"fmt"
	"log"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/asimshrestha2/stream-set/guicontroller"
	"github.com/asimshrestha2/stream-set/helper"
	"github.com/asimshrestha2/stream-set/save"
	"github.com/asimshrestha2/stream-set/twitch"
	ps "github.com/mitchellh/go-ps"
)

var (
	lastWindowChange time.Time
	currentGame      = &game{
		name: "",
		hwnd: 0,
		pid:  -1,
	}

	user32                       = syscall.NewLazyDLL("user32.dll")
	procGetWindowText            = user32.NewProc("GetWindowTextW")
	procGetWindowTextLength      = user32.NewProc("GetWindowTextLengthW")
	procGetWindowThreadProcessID = user32.NewProc("GetWindowThreadProcessId")
)

type (
	HANDLE uintptr
	HWND   HANDLE
	game   struct {
		name string
		hwnd uintptr
		pid  int
	}
)

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(uintptr(hwnd))

	return int(ret)
}

func GetWindowThreadProcessID(hwnd HWND) int {
	var pid int
	procGetWindowThreadProcessID.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&pid)))

	return pid
}

func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1
	buf := make([]uint16, textLen)
	procGetWindowText.Call(
		uintptr(hwnd),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func getWindow(funcName string) uintptr {
	proc := user32.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
}

func GetWindows() {
	var lastTitle = ""
	ticker := time.NewTicker(1 * time.Second)
	gameIndex := -1
	IgnoreList = save.GetIgnoreList()
	log.Println(WaitToReset)
	go func() {
		for t := range ticker.C {
			if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {

				text := GetWindowText(HWND(hwnd))
				trimedText := strings.TrimSpace(text)
				guicontroller.MW.CurrentWindow.SetText("Current Window: " + trimedText)

				lastGameProcess, _ := ps.FindProcess(currentGame.pid)

				if lastTitle != text {
					lastWindowChange = time.Now()
					lastTitle = text

					if twitch.GameDB == nil {
						twitch.GetTopGamesNames()
					}

					currentPID := GetWindowThreadProcessID(HWND(hwnd))
					currentProcess, _ := ps.FindProcess(currentPID)
					gameIndex = helper.ContainsInDB(twitch.GameDB, trimedText, currentProcess.Executable())

					fmt.Println(t, "Updated: Current Window: ", text, " Last Window: ", lastTitle, " GameDB Index: ", gameIndex)
					fmt.Println("Pid: ", currentPID, " #hwnd: ", hwnd)

					if twitch.Token != "" && currentGame.name != trimedText && lastGameProcess == nil &&
						currentGame.pid != currentPID && gameIndex > -1 {

						currentGame.name = trimedText
						currentGame.hwnd = hwnd
						currentGame.pid = currentPID

						log.Println("GameDB: ", twitch.GameDB[gameIndex])

						if helper.ContainsText(IgnoreList, trimedText) <= -1 &&
							twitch.UserChannel.Game != twitch.GameDB[gameIndex].TwitchName {
							fmt.Println("Game Updated To: " + trimedText)
							twitch.UpdateChannelGame(trimedText)
						}

						if twitch.GameDB[gameIndex].FileName == "" {
							twitch.GameDB[gameIndex].FileName = currentProcess.Executable()
							go save.SaveGameList(twitch.GameDB)
						}
						continue
					}
				}

				if gameIndex <= -1 && twitch.Token != "" && twitch.UserChannel.Game != DefaultGame && lastGameProcess == nil &&
					time.Now().Sub(lastWindowChange).Seconds() >= WaitToReset {

					fmt.Println("Game Updated To: " + DefaultGame)
					twitch.UpdateChannelGame(DefaultGame)
				}
			}
		}
	}()
}
