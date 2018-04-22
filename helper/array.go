package helper

import (
	"strings"

	"github.com/asimshrestha2/stream-set/twitch"
)

func ContainsInDB(slice []twitch.DBGame, item string) int {
	for i, s := range slice {
		if s.TwitchName == item {
			return i
		}

		if len(s.AlternativeNames) > 0 {
			if Contains(s.AlternativeNames, item) > -1 {
				return i
			}
		}
	}
	return -1
}

func Contains(slice []string, item string) int {
	for i, s := range slice {
		if s == item {
			return i
		}
	}

	return -1
}

func ContainsText(slice []string, item string) int {
	for i, s := range slice {
		if strings.Contains(s, item) {
			return i
		}
	}

	return -1
}
