package main

import (
	"log"

	"github.com/asimshrestha2/stream-set/save"

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
	go func() {
		save.LoadSettings()
		gamewindows.DefaultGame = save.GetTwitchDefaultGame()
		gamewindows.WaitToReset = save.GetResetTime()
		gamewindows.GetWindows()
	}()

	fontTitle := Font{
		Family:    "Arial",
		PointSize: 16,
	}

	fontSubTitle := Font{
		Family:    "Arial",
		PointSize: 14,
	}

	fontSSubTitle := Font{
		Family:    "Arial",
		PointSize: 10,
	}

	if _, err := (MainWindow{
		AssignTo:   &guicontroller.MW.MainWindow,
		Title:      "Stream Set",
		MinSize:    Size{500, 240},
		Size:       Size{500, 300},
		Background: SolidColorBrush{Color: walk.RGB(29, 37, 44)},
		Layout:     VBox{MarginsZero: true},
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						AssignTo:    &guicontroller.MW.Cleargamelist,
						Text:        "&Clear GameList",
						OnTriggered: guicontroller.MW.Cleargamelist_Triggered,
					},
					Separator{},
					Action{
						Text:        "Exit",
						OnTriggered: func() { guicontroller.MW.Close() },
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						Text: "&About",
						OnTriggered: func() {
							url := "https://github.com/asimshrestha2/stream-set#stream-set"
							browser.OpenURL(url)
						},
					},
				},
			},
		},
		Children: []Widget{
			HSplitter{
				HandleWidth: 0,
				Children: []Widget{
					ImageView{
						AssignTo:   &guicontroller.MW.TwitchImage,
						MinSize:    Size{90, 90},
						MaxSize:    Size{90, 90},
						Background: SolidColorBrush{Color: walk.RGB(29, 37, 44)},
						Image:      "Asim_Ymir.png",
						Margin:     10,
						Mode:       ImageViewModeShrink,
					},
					VSplitter{
						Children: []Widget{
							VSpacer{},
							Label{
								AssignTo:  &guicontroller.MW.TwitchUsername,
								Font:      fontTitle,
								Text:      "Not Logged In",
								TextColor: walk.RGB(225, 225, 225),
							},
							VSpacer{},
							Label{
								Font:      fontSSubTitle,
								Text:      "Current Game: ",
								TextColor: walk.RGB(225, 225, 225),
							},
							Label{
								AssignTo:  &guicontroller.MW.TwitchGame,
								Font:      fontSubTitle,
								Text:      "Unknown",
								TextColor: walk.RGB(225, 225, 225),
							},
							VSpacer{},
						},
					},
				},
			},
			LinkLabel{
				MinSize:  Size{380, 30},
				AssignTo: &guicontroller.MW.LL,
				Text:     `<a id="twitch_token" href="` + twitch.RequestTokenURL + `">Twitch Login</a>`,
				OnLinkActivated: func(link *walk.LinkLabelLink) {
					// log.Printf("id: '%s', url: '%s'\n", link.Id(), link.URL())
					browser.OpenURL(link.URL())
				},
			},
		},
		StatusBarItems: []StatusBarItem{
			StatusBarItem{
				AssignTo:    &guicontroller.MW.CurrentWindow,
				Text:        "Current Window: ",
				ToolTipText: "Current Active Window",
			},
		},
	}.Run()); err != nil {
		log.Fatal(err)
	}

	guicontroller.MW.SetBounds(walk.Rectangle{
		Height: 300,
		Width:  500,
	})
}
