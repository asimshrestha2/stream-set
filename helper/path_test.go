package helper

import (
	"testing"
)

func TestGetGameClient(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		wantClient string
	}{
		// TODO: Add test cases.
		{"Steam", args{"D:\\SteamLibrary\\steamapps\\common\\Tricky Towers\\TrickyTowers.exe"}, "Steam"},
		{"Not Steam", args{"D:\\SteamLibrary\\steamapp\\common\\Tricky Towers\\TrickyTowers.exe"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotClient := GetGameClient(tt.args.path)
			if gotClient != tt.wantClient {
				t.Errorf("GetGameClient() gotClient = %v, want %v", gotClient, tt.wantClient)
			}
		})
	}
}

func TestGetGameNameWithGameClient(t *testing.T) {
	type args struct {
		path   string
		client string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Steam Path 1", args{"D:\\SteamLibrary\\steamapps\\common\\Tricky Towers\\TrickyTowers.exe", "Steam"}, "Tricky Towers"},
		{"Steam Path 2", args{"D:\\Program Files (x86)\\Steam\\steamapps\\common\\PUBG\\TslGame\\Binaries\\Win64\\TslGame.exe", "Steam"}, "PUBG"},
		{"Not Steam Path", args{"D:\\SteamLibrary\\steamapp\\common\\Tricky Towers\\TrickyTowers.exe", "Steam"}, ""},
		{"No Client Path", args{"D:\\SteamLibrary\\steamapps\\common\\Tricky Towers\\TrickyTowers.exe", ""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetGameNameWithGameClient(tt.args.path, tt.args.client); got != tt.want {
				t.Errorf("GetGameNameWithGameClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
