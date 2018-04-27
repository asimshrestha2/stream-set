package twitch

import (
	"encoding/json"
	"errors"
	"net/url"
)

// SearchRespose - Search Response
type SearchRespose struct {
	Games []SGame `json:"games"`
}

// SGame - Search Game
type SGame struct {
	Name          string     `json:"name"`
	Popularity    int64      `json:"popularity"`
	ID            int64      `json:"_id"`
	GiantbombID   int64      `json:"giantbomb_id"`
	Box           GameImages `json:"box"`
	Logo          GameImages `json:"logo"`
	LocalizedName string     `json:"localized_name"`
	Locale        string     `json:"locale"`
}

// SearchGames sends a request to twitch and checks if there is a game based on query
func SearchGames(query string) (DBGame, error) {
	quertP, _ := url.Parse(query)
	uri := TwitchAPIURL + "/search/games?query=" + quertP.String()
	body, err := Request("GET", uri, nil, false, false)
	if err != nil {
		return DBGame{}, err
	}
	sr := SearchRespose{}
	if err := json.Unmarshal([]byte(body), &sr); err != nil {
		return DBGame{}, err
	}
	if len(sr.Games) > 0 {
		return DBGame{TwitchName: sr.Games[0].Name}, nil
	}
	return DBGame{}, errors.New("Not Found")
}
