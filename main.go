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
		MaxSize:  Size{450, 350},
		Layout:   VBox{MarginsZero: true},
		Children: []Widget{
			HSplitter{
				MaxSize:     Size{400, 128},
				HandleWidth: 0,
				Children: []Widget{
					ImageView{
						AssignTo: &guicontroller.MW.TwitchImage,
						Image:    "Asim_Ymir.png",
						Margin:   10,
						Mode:     ImageViewModeShrink,
					},
					VSplitter{
						Children: []Widget{
							Label{
								AssignTo: &guicontroller.MW.TwitchUsername,
								Text:     "<UserName>",
							},
							Label{
								AssignTo: &guicontroller.MW.TwitchGame,
								Text:     "<CurrentGame>",
							},
						},
					},
				},
			},
			LinkLabel{
				AssignTo: &guicontroller.MW.LL,
				Text:     `<a id="twitch_token" href="` + twitch.RequestTokenURL + `">Twitch Login</a>`,
				OnLinkActivated: func(link *walk.LinkLabelLink) {
					guicontroller.MW.LL.SetText("Loading...")
					// log.Printf("id: '%s', url: '%s'\n", link.Id(), link.URL())
					browser.OpenURL(link.URL())
				},
			},
		},
		StatusBarItems: []StatusBarItem{
			StatusBarItem{
				AssignTo:    &guicontroller.MW.CurrentWindow,
				Text:        "Current Window: <window>",
				ToolTipText: "Current Active Window",
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}
}
