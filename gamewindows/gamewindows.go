package gamewindows

import (
	"fmt"
	"strings"
	"syscall"
	"time"
	"unsafe"

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

func doesProccessExist(pid int) bool {
	p, err := ps.FindProcess(pid)

	if err != nil {
		fmt.Println("Error : ", err)
		return false
	}
	if p != nil {
		return true
	}

	return false
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
					gameInList := contains(twitch.GameNameList, trimedText)
					currentPID := GetWindowThreadProcessID(HWND(hwnd))
					fmt.Println(t, "Updated: Current Window: ", text, " Last Window: ", lastTitle, " Game in List: ", gameInList)
					fmt.Println("Pid: ", currentPID, " #hwnd: ", hwnd)
					if twitch.Token != "" && !doesProccessExist(currentGame.pid) && currentGame.name != trimedText && currentGame.pid != currentPID && gameInList {
						currentGame.name = trimedText
						currentGame.hwnd = hwnd
						currentGame.pid = currentPID
						fmt.Println("Game Updated To: " + trimedText)
						twitch.UpdateChannelGame(trimedText)
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
