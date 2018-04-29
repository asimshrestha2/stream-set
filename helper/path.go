package helper

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/asimshrestha2/stream-set/steam"
)

var (
	GameClients = map[string]string{
		"Steam": "\\steamapps\\common\\",
	}
)

// GetGameClient Checks if the path is a game library path
func GetGameClient(path string) (client string) {
	for c, gcpath := range GameClients {
		if strings.Contains(path, gcpath) {
			return c
		}
	}
	return ""
}

func GetGameNameWithGameClient(filepath string, client string) string {
	gcpath, ok := GameClients[client]
	if ok && strings.Contains(filepath, gcpath) {
		fileSplit := strings.Split(filepath, gcpath)
		tmpS := strings.Split(fileSplit[1], "\\")
		appidpath := path.Join(fileSplit[0], gcpath, tmpS[0], "steam_appid.txt")
		if _, err := os.Stat(appidpath); os.IsNotExist(err) {
			return tmpS[0]
		}

		f, err := ioutil.ReadFile(appidpath)
		if err != nil {
			return tmpS[0]
		}
		appid, err := strconv.Atoi(string(f))
		if err != nil {
			return tmpS[0]
		}

		name := steam.FindGameName(appid)
		if name == "" {
			return tmpS[0]
		}
		fmt.Println("Appid: ", appid, " Name: ", name)
		return name
	}
	return ""
}
