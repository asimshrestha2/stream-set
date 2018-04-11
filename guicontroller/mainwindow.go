package guicontroller

import "github.com/lxn/walk"

type MyMainWindow struct {
	*walk.MainWindow
	LL             *walk.LinkLabel
	TwitchUsername *walk.Label
	TwitchImage    *walk.ImageView
}

var MW = &MyMainWindow{}
