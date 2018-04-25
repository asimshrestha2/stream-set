package helper

import (
	"testing"

	"github.com/asimshrestha2/stream-set/twitch"
)

func TestContainsInDB(t *testing.T) {
	type args struct {
		slice    []twitch.DBGame
		item     string
		filename string
	}

	db := []twitch.DBGame{
		{TwitchName: "Game 1", FileName: "Game1.exe", AlternativeNames: nil},
		{TwitchName: "Game 2", FileName: "Game2.exe", AlternativeNames: []string{"Game 2.0", "Game 2.1"}},
		{TwitchName: "Game 3", FileName: "Game3.exe", AlternativeNames: nil},
		{TwitchName: "Game 4", FileName: "Game4.exe", AlternativeNames: nil},
		{TwitchName: "Game 5", FileName: "Game5.exe", AlternativeNames: nil},
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"When item not found",
			args{slice: db, item: "apple", filename: "apple.exe"},
			-1,
		},
		{
			"When item found without filename",
			args{slice: db, item: "Game 1", filename: ""},
			0,
		},
		{
			"When item found with filename",
			args{slice: db, item: "Aasdf", filename: "Game3.exe"},
			2,
		},
		{
			"When item found with AltName",
			args{slice: db, item: "Game 2.0", filename: ""},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInDB(tt.args.slice, tt.args.item, tt.args.filename); got != tt.want {
				t.Errorf("ContainsInDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsText(t *testing.T) {
	type args struct {
		slice []string
		item  string
	}

	list := []string{"Firefox", "Chrome", "Apple"}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"Not found in list", args{list, "Fire Name"}, -1},
		{"Found full name", args{list, "Apple"}, 2},
		{"Found in Middle", args{list, "ASDF Firefox Name"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsText(tt.args.slice, tt.args.item); got != tt.want {
				t.Errorf("ContainsText() = %v, want %v", got, tt.want)
			}
		})
	}
}
