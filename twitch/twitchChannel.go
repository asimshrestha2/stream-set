package twitch

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/lxn/walk"

	"github.com/asimshrestha2/stream-set/guicontroller"
	"github.com/asimshrestha2/stream-set/save"
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

type ChannelG struct {
	ChannelA GameC `json:"channel"`
}

type GameC struct {
	GameA string `json:"game"`
}

func SetTwitchChannel() {
	UserChannel = GetChannelInfo()
	save.Image(UserChannel.Logo)
	img, err := walk.NewImageFromFile(save.ImagePathFromURL(UserChannel.Logo))
	if err != nil {
		log.Println(err)
	}
	guicontroller.MW.TwitchImage.SetImage(img)
	guicontroller.MW.TwitchUsername.SetText(UserChannel.DisplayName)
	guicontroller.MW.TwitchGame.SetText(UserChannel.Game)
}

func GetChannelInfo() Channel {
	body, err := TwitchRequest("GET", TwitchAPIURL+"/channel", nil, true, false)
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
		resC := &ChannelG{
			ChannelA: GameC{
				GameA: game,
			},
		}

		res2B, _ := json.Marshal(resC)

		if _, err := TwitchRequest("PUT", TwitchAPIURL+"/channels/"+UserChannel.ID, bytes.NewBuffer(res2B), true, true); err != nil {
			log.Panicf("%s\n", err)
		} else {
			SetTwitchChannel()
		}

	}
}
