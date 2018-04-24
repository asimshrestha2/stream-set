package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/asimshrestha2/stream-set/helper"
	"github.com/asimshrestha2/stream-set/twitch"

	"github.com/Joker/jade"

	"github.com/julienschmidt/httprouter"
)

// Index for Website
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<html><body>Hi</body></html>")
}

// Gamelist for App
func Gamelist(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	data, err := ioutil.ReadFile("./server/template/gamelist.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, string(data))
}

// GamelistPost : To update the Game that was updated
func GamelistPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var dbData twitch.DBGame
	err := decoder.Decode(&dbData)
	if err != nil {
		fmt.Fprintf(w, "")
	}
	defer r.Body.Close()
	// log.Println(dbData)
	helper.UpdateInDB(dbData)
	fmt.Fprintf(w, "")
}

// Notification for Website
func Notification(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notificationTemp, _ := template.ParseFiles("./server/template/notification.html")
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
	data, err := ioutil.ReadFile("./server/template/settings.pug")
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

//FileImage : to convert image location threw server
func FileImage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	filepath := r.FormValue("path")
	http.ServeFile(w, r, filepath)
}
