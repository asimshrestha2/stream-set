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

	mw := MainWindow{
		AssignTo:   &guicontroller.MW.MainWindow,
		Title:      "Stream Set",
		MinSize:    Size{500, 240},
		Size:       Size{500, 300},
		Icon:       "icon.ico",
		Background: SolidColorBrush{Color: walk.RGB(29, 37, 44)},
		Layout:     VBox{MarginsZero: true},
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						Text: "&Edit GameList",
						OnTriggered: func() {
							url := "http://localhost:8000/gamelist"
							browser.OpenURL(url)
						},
					},
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
						Image:      "account.png",
						Margin:     10,
						Mode:       ImageViewModeShrink,
					},
					VSplitter{
						Children: []Widget{
							VSpacer{},
							Label{
								AssignTo:  &guicontroller.MW.TwitchUsername,
								Font:      guicontroller.FontTitle,
								Text:      "Not Logged In",
								TextColor: walk.RGB(225, 225, 225),
							},
							VSpacer{},
							Label{
								Font:      guicontroller.FontSSubTitle,
								Text:      "Current Game: ",
								TextColor: walk.RGB(225, 225, 225),
							},
							Label{
								AssignTo:  &guicontroller.MW.TwitchGame,
								Font:      guicontroller.FontSubTitle,
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
	}

	icon, err := walk.Resources.Icon("icon.ico")
	if err != nil {
		log.Fatal(err)
	}

	ni, err := walk.NewNotifyIcon()
	if err != nil {
		log.Fatal(err)
	}
	defer ni.Dispose()

	if err := ni.SetIcon(icon); err != nil {
		log.Fatal(err)
	}
	if err := ni.SetToolTip("Click for info or use the context menu to exit."); err != nil {
		log.Fatal(err)
	}

	// When the left mouse button is pressed, bring up our balloon.
	ni.MouseDown().Attach(func(x, y int, button walk.MouseButton) {
		if button != walk.LeftButton {
			return
		}

		guicontroller.MW.MainWindow.Show()
	})

	hideAction := walk.NewAction()
	if err := hideAction.SetText("Hide Stream Set"); err != nil {
		log.Fatal(err)
	}
	hideAction.Triggered().Attach(func() { guicontroller.MW.MainWindow.Hide() })
	if err := ni.ContextMenu().Actions().Add(hideAction); err != nil {
		log.Fatal(err)
	}
	// We put an exit action into the context menu.
	exitAction := walk.NewAction()
	if err := exitAction.SetText("E&xit"); err != nil {
		log.Fatal(err)
	}
	exitAction.Triggered().Attach(func() { walk.App().Exit(0) })
	if err := ni.ContextMenu().Actions().Add(exitAction); err != nil {
		log.Fatal(err)
	}

	// The notify icon is hidden initially, so we have to make it visible.
	if err := ni.SetVisible(true); err != nil {
		log.Fatal(err)
	}

	// Now that the icon is visible, we can bring up an info balloon.
	if err := ni.ShowInfo("Walk NotifyIcon Example", "Click the icon to show again."); err != nil {
		log.Fatal(err)
	}

	if _, err := mw.Run(); err != nil {
		log.Fatal(err)
	}

	guicontroller.MW.MainWindow.SetIcon(icon)
}
