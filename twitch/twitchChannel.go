package twitch

import (
	"bytes"
	"encoding/json"
	"log"
)

type Channel struct {
	Mature                       bool   `json:"mature"`
	Status                       string `json:"status"`
	BroadcasterLanguage          string `json:"broadcaster_language"`
	DisplayName                  string `json:"display_name"`
	Game                         string `json:"game"`
	Language                     string `json:"language"`
	ID                           string `json:"_id"`
	Name                         string `json:"name"`
	CreatedAt                    string `json:"created_at"`
	UpdatedAt                    string `json:"updated_at"`
	Partner                      bool   `json:"partner"`
	Logo                         string `json:"logo"`
	VideoBanner                  string `json:"video_banner"`
	ProfileBanner                string `json:"profile_banner"`
	ProfileBannerBackgroundColor string `json:"profile_banner_background_color"`
	URL                          string `json:"url"`
	Views                        int64  `json:"views"`
	Followers                    int64  `json:"followers"`
	BroadcasterType              string `json:"broadcaster_type"`
	StreamKey                    string `json:"stream_key"`
	Email                        string `json:"email"`
}

func GetChannelInfo() Channel {
	body, err := TwitchRequest("GET", TwitchAPIURL+"/channel", nil, true)
	if err != nil {
		log.Panicf("%s\n", err)
	}
	ch := Channel{}
	if err := json.Unmarshal([]byte(body), &ch); err != nil {
		log.Panicf("%s\n", err)
	}
	return ch
}

func UpdateChannelGame(game string) {
	if game != "" {
		var jsonStr = []byte(`{"channel":{"game": ` + game + `}}`)
		if _, err := TwitchRequest("PUT", TwitchAPIURL+"/channel/"+UserChannel.ID, bytes.NewBuffer(jsonStr), true); err != nil {
			log.Panicf("%s\n", err)
		}
	}
}
