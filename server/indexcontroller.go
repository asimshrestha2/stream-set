package server

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/Joker/jade"

	"github.com/julienschmidt/httprouter"
)

// Index for Website
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<html><body>Hi</body></html>")
}

// Notification for Website
func Notification(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notificationTemp, _ := template.ParseFiles("./template/notification.html")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	p, lastMod, err := readFileIfModified(time.Time{})
	if err != nil {
		p = []byte(err.Error())
		lastMod = time.Unix(0, 0)
	}
	var v = struct {
		Host         string
		Data         string
		LastMod      string
		CircleColor  string
		TimerSetting int64
		LogoPath     string
	}{
		r.Host,
		string(p),
		strconv.FormatInt(lastMod.UnixNano(), 16),
		sf.CircleColor,
		sf.TimerSetting,
		sf.LogoPath,
	}
	notificationTemp.Execute(w, &v)
}

// Settings for Website
func Settings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data, err := ioutil.ReadFile("./template/settings.pug")
	if err != nil {
		fmt.Fprintf(w, "File Error: %v", err)
		return
	}
	if data != nil {
		t, err := jade.Parse("settings", string(data))
		if err != nil {
			fmt.Fprintf(w, "Parse Error: %v", err)
			return
		}
		fmt.Fprintf(w, t)
	}
}

//FileImage: to convert image location threw server
func FileImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	filepath := r.FormValue("path")
	http.ServeFile(w, r, filepath)
}
