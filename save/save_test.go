package save

import (
	"testing"
)

func Test_LoadSettings(t *testing.T) {
	LoadSettings()
	if CFG == nil {
		t.Errorf("Couldn't Find Any Section")
	}
}

func Test_LoadGameList(t *testing.T) {
	var db []struct {
		TwitchName       string   `json:"twitchName"`
		FileName         string   `json:"fileName"`
		AlternativeNames []string `json:"alternativeName"`
	}
	LoadGameList(&db)
	if len(db) <= 0 {
		t.Errorf("Couldn't Load GameList")
	}
}
