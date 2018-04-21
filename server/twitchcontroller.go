package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/asimshrestha2/stream-set/guicontroller"
	"github.com/asimshrestha2/stream-set/twitch"

	"github.com/julienschmidt/httprouter"
)

var (
	setChannel = false
)

func TwitchTokenAPI(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	q := r.URL.Query()
	if twitch.Token == "" && q.Get("access_token") == "" {
		fmt.Fprintf(w, `
			<script>
				(function(){
					if(window.location.hash !== ""){
						window.location.href = "//" + window.location.host + window.location.pathname + "?" + window.location.hash.replace("#", '');
					}
				})()
			</script>
			<a id="twitch_token" href="%s">Twitch Login</a>
		`, twitch.RequestTokenURL)
	} else {
		twitch.Token = q.Get("access_token")
		go func() {
			if !setChannel {
				twitch.SetTwitchChannel()
				twitch.GetTopGamesNames()
				setChannel = true
			}
		}()
		// guicontroller.MW.LL.SetText("Logged In")
		guicontroller.MW.LL.SetVisible(false)
		fmt.Fprintf(w, "<html><body><div>You can close this window now. :D</div></body></html>")
	}
}

func TwitchGameListAPI(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	retj, _ := json.Marshal(twitch.GameDB)
	fmt.Fprintf(w, "%s", retj)
}
