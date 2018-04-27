package twitch

import (
	"reflect"
	"testing"
)

func TestSearchGames(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    DBGame
		wantErr bool
	}{
		{"Search Game that Exist", args{"Rocket League"}, DBGame{TwitchName: "Rocket League"}, false},
		{"Search Game that Does not exist", args{"Rocket League 2"}, DBGame{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SearchGames(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchGames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchGames() = %v, want %v", got, tt.want)
			}
		})
	}
}
