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
	gameChange  time.Time
	currentGame = &game{
		name: "",
		hwnd: 0,
		pid:  -1,
	}

	user32                       = syscall.NewLazyDLL("user32.dll")
	procGetWindowText            = user32.NewProc("GetWindowTextW")
	procGetWindowTextLength      = user32.NewProc("GetWindowTextLengthW")
	procGetWindowThreadProcessID = user32.NewProc("GetWindowThreadProcessId")
	procGetWindowModuleFileName  = user32.NewProc("GetWindowModuleFileNameW")

	kernel32                      = syscall.NewLazyDLL("kernel32.dll")
	procOpenProcess               = kernel32.NewProc("OpenProcess")
	procCloseHandle               = kernel32.NewProc("CloseHandle")
	procGetVolumePathName         = kernel32.NewProc("GetVolumePathNameW")
	procQueryFullProcessImageName = kernel32.NewProc("QueryFullProcessImageNameW")

	psapi                       = syscall.NewLazyDLL("Psapi.dll")
	procGetProcessImageFileName = psapi.NewProc("GetProcessImageFileNameW")
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

func GetWindowModuleFileName(hwnd HWND) string {
	textLen := 2000
	buf := make([]uint16, textLen)
	procGetWindowModuleFileName.Call(
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

func OpenProcess(desiredAccess uint32, inheritHandle bool, processId uint32) (handle HANDLE, err error) {
	inherit := 0
	if inheritHandle {
		inherit = 1
	}

	ret, _, err := procOpenProcess.Call(
		uintptr(desiredAccess),
		uintptr(inherit),
		uintptr(processId))
	if err != nil && err.Error() == "The operation completed successfully." {
		err = nil
	}
	handle = HANDLE(ret)
	return
}

func CloseHandle(object HANDLE) bool {
	ret, _, _ := procCloseHandle.Call(
		uintptr(object))
	return ret != 0
}

func GetVolumePathName(path string) string {
	textLen := 256
	inpath := syscall.StringToUTF16(path)
	buf := make([]uint16, textLen)
	procGetVolumePathName.Call(
		uintptr(unsafe.Pointer(&inpath[0])),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func QueryFullProcessImageName(handle HANDLE) string {
	textLen := 256
	buf := make([]uint16, textLen)
	procQueryFullProcessImageName.Call(
		uintptr(handle),
		uintptr(0),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(unsafe.Pointer(&textLen)))

	return syscall.UTF16ToString(buf)
}

func GetProcessImageFileName(handle HANDLE) string {
	textLen := 256
	buf := make([]uint16, textLen)
	procGetProcessImageFileName.Call(
		uintptr(handle),
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(textLen))

	return syscall.UTF16ToString(buf)
}

func GetWindows() {
	var lastTitle = ""
	ticker := time.NewTicker(1 * time.Second)
	gameIndex := -1
	IgnoreList = save.GetIgnoreList()

	go func() {
		for t := range ticker.C {
			if hwnd := getWindow("GetForegroundWindow"); hwnd != 0 {

				text := GetWindowText(HWND(hwnd))

				trimedText := strings.TrimSpace(text)
				guicontroller.MW.CurrentWindow.SetText("Current Window: " + trimedText)

				lastGameProcess, _ := ps.FindProcess(currentGame.pid)

				if lastTitle != text {
					lastTitle = text

					if twitch.GameDB == nil {
						twitch.GetTopGamesNames()
					}

					currentPID := GetWindowThreadProcessID(HWND(hwnd))
					handler, err := OpenProcess(1024|16, true, uint32(currentPID))
					if err != nil {
						fmt.Println(err)
					}

					filepath := QueryFullProcessImageName(handler)
					CloseHandle(handler)

					gameIndex = helper.ContainsInDB(twitch.GameDB, trimedText, filepath)
					gameClient := helper.GetGameClient(filepath)

					fmt.Printf("%v - Current Window: %s (Index: %d, PID: %d, Path: %s)\n", t, trimedText, gameIndex, currentPID, filepath)

					if gameIndex <= -1 && gameClient != "" {
						gdb, err := twitch.SearchGames(helper.GetGameNameWithGameClient(filepath, gameClient))
						if err != nil {
							log.Println(err)
						} else {
							gdb.FilePath = filepath
							twitch.GameDB = append(twitch.GameDB, gdb)
							gameIndex = len(twitch.GameDB) - 1
							go save.SaveGameList(twitch.GameDB)
						}
					}

					if twitch.Token != "" && currentGame.name != trimedText && lastGameProcess == nil &&
						currentGame.pid != currentPID && gameIndex > -1 {
						gameChange = time.Now()

						currentGame.name = trimedText
						currentGame.hwnd = hwnd
						currentGame.pid = currentPID

						log.Println("GameDB: ", twitch.GameDB[gameIndex])

						if helper.ContainsText(IgnoreList, trimedText) <= -1 &&
							twitch.UserChannel.Game != twitch.GameDB[gameIndex].TwitchName {
							fmt.Println("Game Updated To: " + twitch.GameDB[gameIndex].TwitchName)
							twitch.UpdateChannelGame(twitch.GameDB[gameIndex].TwitchName)
						}

						if twitch.GameDB[gameIndex].FilePath == "" {
							twitch.GameDB[gameIndex].FilePath = filepath
							go save.SaveGameList(twitch.GameDB)
						}
						continue
					}
				}

				if gameIndex <= -1 && twitch.Token != "" && twitch.UserChannel.Game != DefaultGame && lastGameProcess == nil &&
					time.Now().Sub(gameChange).Seconds() >= WaitToReset {

					fmt.Println("Game Updated To: " + DefaultGame)
					twitch.UpdateChannelGame(DefaultGame)
				}
			}
		}
	}()
}
