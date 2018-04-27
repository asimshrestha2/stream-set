package guicontroller

import (
	"log"

	dec "github.com/lxn/walk/declarative"
	"github.com/lxn/win"

	"github.com/asimshrestha2/stream-set/save"
	"github.com/lxn/walk"
)

type MyMainWindow struct {
	*walk.MainWindow
	LL             *walk.LinkLabel
	TwitchUsername *walk.Label
	TwitchGame     *walk.Label
	TwitchImage    *walk.ImageView
	CurrentWindow  *walk.StatusBarItem
	Cleargamelist  *walk.Action
}

var (
	MW = &MyMainWindow{}

	FontTitle = dec.Font{
		Family:    "Arial",
		PointSize: 16,
	}

	FontSubTitle = dec.Font{
		Family:    "Arial",
		PointSize: 14,
	}

	FontSSubTitle = dec.Font{
		Family:    "Arial",
		PointSize: 10,
	}
)

func (mw *MyMainWindow) Cleargamelist_Triggered() {
	i := walk.MsgBox(mw, "Clear Gamelist", "Are you sure you want to delete your Gamelist", walk.MsgBoxIconWarning|walk.MsgBoxYesNo)
	log.Print(i == win.IDYES)
	if i == win.IDYES {
		save.RemoveGameList()
	}
	// twitch.GetTopGamesNames()
}
