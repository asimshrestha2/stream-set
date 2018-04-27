package guicontroller

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func OpenWebWindow(title string, url string) {
	var wv *walk.WebView

	MainWindow{
		Title:   title,
		MinSize: Size{800, 600},
		Layout:  VBox{MarginsZero: true},
		Children: []Widget{
			WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      url,
			},
		},
	}.Run()
}
