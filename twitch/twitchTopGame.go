package twitch

import (
	"encoding/json"
	"log"
)

type TopGamesResponse struct {
	Total int64        `json:"_total"`
	Top   []TwitchGame `json:"top"`
}

type TwitchGame struct {
	Channels int64 `json:"channels"`
	Viewers  int64 `json:"viewers"`
	Game     Game  `json:"game"`
}

type Game struct {
	ID          int64      `json:"_id"`
	Box         GameImages `json:"box"`
	GiantbombID int64      `json:"giantbomb_id"`
	Logo        GameImages `json:"logo"`
	Name        string     `json:"name"`
	Popularity  int64      `json:"popularity"`
}

type GameImages struct {
	Large    string `json:"large"`
	Medium   string `json:"medium"`
	Small    string `json:"small"`
	Template string `json:"template"`
}

func GetTopGames() TopGamesResponse {
	body, err := TwitchRequest("GET", TwitchAPIURL+"/games/top?limit=100", nil, false, false)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	topGamesResponse := TopGamesResponse{}
	if err := json.Unmarshal([]byte(body), &topGamesResponse); err != nil {
		log.Panicf("%s\n", err)
	}
	return topGamesResponse
}

func GetTopGamesNames() {
	tgr := GetTopGames()
	for _, g := range tgr.Top {
		GameNameList = append(GameNameList, g.Game.Name)
	}
}
