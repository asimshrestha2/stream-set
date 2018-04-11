package gamewindows

import (
	"fmt"
	"log"
	"syscall"
	"time"
	"unsafe"

	ps "github.com/mitchellh/go-ps"
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
	ticker := time.NewTicker(1 * time.Second)

	IgnoreList := []string{"svchost.exe", "explorer.exe", "System", "[System Process]", "winlogon.exe"}

	go func() {
		for t := range ticker.C {

			// if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {
			// 	text := GetWindowText(HWND(hwnd))
			// 	fmt.Println("window :", text, "# hwnd:", hwnd)
			// }
			fmt.Println("Tick at", t)
			p, err := ps.Processes()
			if err != nil {
				log.Fatalf("err: %s", err)
			}

			if len(p) <= 0 {
				log.Fatal("should have processes")
			}

			for _, p1 := range p {
				if !contains(IgnoreList, p1.Executable()) {
					log.Print(p1)
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
