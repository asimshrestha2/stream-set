package gamewindows

import (
	"fmt"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"github.com/asimshrestha2/stream-set/twitch"
)

var (
	mod                     = syscall.NewLazyDLL("user32.dll")
	procGetWindowText       = mod.NewProc("GetWindowTextW")
	procGetWindowTextLength = mod.NewProc("GetWindowTextLengthW")
)

type (
	HANDLE uintptr
	HWND   HANDLE
)

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := procGetWindowTextLength.Call(uintptr(hwnd))

	return int(ret)
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
	proc := mod.NewProc(funcName)
	hwnd, _, _ := proc.Call()
	return hwnd
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
					gameInList := contains(twitch.GameNameList, strings.TrimSpace(text))
					fmt.Println(t, "Updated: Current Window: ", text, " #hwnd: ", hwnd, " Last Window: ", lastTitle, " Game in List: ", gameInList)
					if twitch.Token != "" && gameInList {
						fmt.Println("Game Updated To: " + strings.TrimSpace(text))
						twitch.UpdateChannelGame(strings.TrimSpace(text))
					}
				}
			}
		}
	}()

	// time.Sleep(1600 * time.Millisecond)
	// ticker.Stop()
	// fmt.Println("Ticker stopped")
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
