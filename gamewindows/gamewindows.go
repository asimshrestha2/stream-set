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
	gameIndex        = -1
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

	IgnoreList = save.GetIgnoreList()

	go func() {
		for t := range ticker.C {
			if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {
				text := GetWindowText(HWND(hwnd))
				lastGameProcess, _ := ps.FindProcess(currentGame.pid)
				if lastTitle != text {
					lastWindowChange = time.Now()
					lastTitle = text
					trimedText := strings.TrimSpace(text)

					if twitch.GameDB == nil {
						twitch.GetTopGamesNames()
					}

					gameIndex = helper.ContainsInDB(twitch.GameDB, trimedText)
					currentPID := GetWindowThreadProcessID(HWND(hwnd))

					fmt.Println(t, "Updated: Current Window: ", text, " Last Window: ", lastTitle, " GameDB Index: ", gameIndex)
					fmt.Println("Pid: ", currentPID, " #hwnd: ", hwnd)
					guicontroller.MW.CurrentWindow.SetText("Current Window: " + trimedText)

					if twitch.Token != "" && lastGameProcess == nil && currentGame.name != trimedText &&
						currentGame.pid != currentPID && gameIndex > -1 {

						currentGame.name = trimedText
						currentGame.hwnd = hwnd
						currentGame.pid = currentPID

						gameProcess, _ := ps.FindProcess(currentPID)

						log.Println("GameDB: ", twitch.GameDB[gameIndex])

						if helper.ContainsText(IgnoreList, trimedText) <= -1 &&
							twitch.UserChannel.Game != twitch.GameDB[gameIndex].TwitchName {
							fmt.Println("Game Updated To: " + trimedText)
							twitch.UpdateChannelGame(trimedText)
						}

						if twitch.GameDB[gameIndex].FileName == "" {
							twitch.GameDB[gameIndex].FileName = gameProcess.Executable()
							go save.SaveGameList(twitch.GameDB)
						}
						continue
					}
				}

				if gameIndex <= -1 && twitch.Token != "" && lastGameProcess == nil &&
					twitch.UserChannel.Game != DefaultGame && time.Now().Sub(lastWindowChange).Seconds() >= WaitToReset {

					fmt.Println("Game Updated To: " + DefaultGame)
					twitch.UpdateChannelGame(DefaultGame)
				}
			}
		}
	}()
}
