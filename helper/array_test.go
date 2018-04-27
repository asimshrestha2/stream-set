package helper

import (
	"testing"

	"github.com/asimshrestha2/stream-set/twitch"
)

func TestContainsInDB(t *testing.T) {
	type args struct {
		slice    []twitch.DBGame
		item     string
		filepath string
	}

	db := []twitch.DBGame{
		{TwitchName: "Game 1", FilePath: "Game1.exe", AlternativeNames: nil},
		{TwitchName: "Game 2", FilePath: "Game2.exe", AlternativeNames: []string{"Game 2.0", "Game 2.1"}},
		{TwitchName: "Game 3", FilePath: "Game3.exe", AlternativeNames: nil},
		{TwitchName: "Game 4", FilePath: "Game4.exe", AlternativeNames: nil},
		{TwitchName: "Game 5", FilePath: "Game5.exe", AlternativeNames: nil},
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			"When item not found",
			args{db, "apple", "apple.exe"},
			-1,
		},
		{
			"When item found without filename",
			args{db, "Game 1", ""},
			0,
		},
		{
			"When item found with filename",
			args{db, "Aasdf", "Game3.exe"},
			2,
		},
		{
			"When item found with AltName",
			args{db, "Game 2.0", ""},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInDB(tt.args.slice, tt.args.item, tt.args.filepath); got != tt.want {
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
