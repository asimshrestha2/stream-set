package gamewindows

import (
	"fmt"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/asimshrestha2/stream-set/guicontroller"
	"github.com/asimshrestha2/stream-set/save"
	"github.com/asimshrestha2/stream-set/twitch"
	ps "github.com/mitchellh/go-ps"
)

var (
	currentGame = &game{
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

func doesProccessExist(pid int) ps.Process {
	p, err := ps.FindProcess(pid)

	if err != nil {
		fmt.Println("Error : ", err)
		return nil
	}

	return p
}

func GetWindows() {
	var lastTitle = ""
	ticker := time.NewTicker(1 * time.Second)

	// IgnoreList := []string{"svchost.exe", "explorer.exe", "System", "[System Process]", "winlogon.exe"}

	go func() {
		for t := range ticker.C {
			if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {
				text := GetWindowText(HWND(hwnd))
				if lastTitle != text {
					lastTitle = text
					trimedText := strings.TrimSpace(text)
					gameIndex := containsInDB(twitch.GameDB, trimedText)
					currentPID := GetWindowThreadProcessID(HWND(hwnd))
					currentProcces, _ := ps.FindProcess(currentGame.pid)
					fmt.Println(t, "Updated: Current Window: ", text, " Last Window: ", lastTitle, " Game in List: ", gameIndex)
					fmt.Println("Pid: ", currentPID, " #hwnd: ", hwnd, twitch.Token != "" && currentProcces != nil && currentGame.name != trimedText && currentGame.pid != currentPID && gameIndex > -1)
					guicontroller.MW.CurrentWindow.SetText("Current Window: " + trimedText)
					if twitch.Token != "" && currentProcces != nil && currentGame.name != trimedText && currentGame.pid != currentPID && gameIndex > -1 {
						currentGame.name = trimedText
						currentGame.hwnd = hwnd
						currentGame.pid = currentPID
						fmt.Println("Game Updated To: " + trimedText)
						twitch.UpdateChannelGame(trimedText)

						if twitch.GameDB[gameIndex].FileName == "" {
							twitch.GameDB[gameIndex].FileName = currentProcces.Executable()
							go save.SaveGameList(twitch.GameDB)
						}

					}
				}
			}
		}
	}()

	// time.Sleep(1600 * time.Millisecond)
	// ticker.Stop()
	// fmt.Println("Ticker stopped")
}

func containsInDB(slice []twitch.DBGame, item string) int {
	for i, s := range slice {
		if s.TwitchName == item {
			return i
		}
	}
	return -1
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
