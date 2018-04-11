package main

import (
	"log"

	"github.com/asimshrestha2/stream-set/gamewindows"
	"github.com/asimshrestha2/stream-set/guicontroller"
	"github.com/asimshrestha2/stream-set/server"
	"github.com/asimshrestha2/stream-set/twitch"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/pkg/browser"
)

func main() {
	walk.Resources.SetRootDirPath("./img")

	go server.StartServer()
	go gamewindows.GetWindows()

	if _, err := (MainWindow{
		AssignTo: &guicontroller.MW.MainWindow,
		Title:    "Stream Set",
		MinSize:  Size{320, 240},
		Size:     Size{400, 300},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				MaxSize: Size{400, 128},
				Children: []Widget{
					ImageView{
						AssignTo: &guicontroller.MW.TwitchImage,
						Image:    "Asim_Ymir.png",
						Margin:   10,
						Mode:     ImageViewModeShrink,
					},
					Label{
						AssignTo: &guicontroller.MW.TwitchUsername,
						Text:     "<username>",
					},
				},
			},
			LinkLabel{
				AssignTo: &guicontroller.MW.LL,
				Text:     `<a id="twitch_token" href="` + twitch.RequestTokenURL + `">Twitch Login</a>`,
				OnLinkActivated: func(link *walk.LinkLabelLink) {
					guicontroller.MW.LL.SetText("Loading...")
					log.Printf("id: '%s', url: '%s'\n", link.Id(), link.URL())
					browser.OpenURL(link.URL())
				},
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
