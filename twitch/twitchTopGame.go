package twitch

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/asimshrestha2/stream-set/save"
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

func GetTopGames(limit int, offset int) TopGamesResponse {
	var url = TwitchAPIURL + "/games/top?limit=" + strconv.Itoa(limit)
	if offset > 0 {
		url += "&offset=" + strconv.Itoa(offset)
	}
	body, err := Request("GET", url, nil, false, false)
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
	if save.GameListExist() {
		err := save.LoadGameList(&GameDB)
		if err != nil {
			log.Fatalln(err)
		}
		// log.Println(GameDB)
	} else {
		tgr := GetTopGames(100, 0)
		time.Sleep(500 * time.Millisecond)
		tgr1 := GetTopGames(100, 100)
		tgr.Top = append(tgr.Top, tgr1.Top...)
		for _, g := range tgr.Top {
			tempGame := DBGame{
				TwitchName: g.Game.Name,
			}
			GameDB = append(GameDB, tempGame)
		}

		go save.SaveGameList(GameDB)
	}
}
