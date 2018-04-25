package helper

import (
	"strings"

	"github.com/asimshrestha2/stream-set/save"

	"github.com/asimshrestha2/stream-set/twitch"
)

//UpdateInDB updates the DataBase Game List
func UpdateInDB(data twitch.DBGame) {
	for i, d := range twitch.GameDB {
		if d.TwitchName == data.TwitchName {
			twitch.GameDB[i].FileName = data.FileName
			twitch.GameDB[i].AlternativeNames = data.AlternativeNames
		}
	}
	go save.SaveGameList(twitch.GameDB)
}

//ContainsInDB Checks if the title, filename of the window contains in the database
func ContainsInDB(slice []twitch.DBGame, item string, filename string) int {
	for i, s := range slice {
		if s.TwitchName == item {
			return i
		}

		if filename != "" && s.FileName == filename {
			return i
		}

		if len(s.AlternativeNames) > 0 && Contains(s.AlternativeNames, item) > -1 {
			return i
		}
	}
	return -1
}

//Contains returns index in from the array for the matching item in the array with the item
func Contains(slice []string, item string) int {
	for i, s := range slice {
		if s == item {
			return i
		}
	}

	return -1
}

//ContainsText returns index in from the array for the matching item in the array with a part in item (string)
func ContainsText(slice []string, item string) int {
	for i, s := range slice {
		if strings.Contains(item, s) {
			return i
		}
	}

	return -1
}
