package helper

import (
	"strings"
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

func GetGameNameWithGameClient(path string, client string) string {
	gcpath, ok := GameClients[client]
	if ok && strings.Contains(path, gcpath) {
		tmpS := strings.Split(strings.Split(path, gcpath)[1], "\\")
		return tmpS[0]
	}
	return ""
}
