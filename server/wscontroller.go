package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
)

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Millisecond

	// Time allowed to read the next pong message from the client.
	pongWait = 60 * time.Second

	// Send pings to client with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Poll file for changes with this period.
	filePeriod = 10 * time.Millisecond
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func readFileIfModified(lastMod time.Time) ([]byte, time.Time, error) {
	fileInfo, err := os.Stat(sf.FileLocation)
	if err != nil {
		return nil, lastMod, err
	}
	if !fileInfo.ModTime().After(lastMod) {
		return nil, lastMod, nil
	}
	fileData, err := ioutil.ReadFile(sf.FileLocation)
	if err != nil {
		return nil, fileInfo.ModTime(), err
	}
	return fileData, fileInfo.ModTime(), nil
}

func reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(512)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error {
		ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}

}

func writer(ws *websocket.Conn, lastMod time.Time) {
	lastError := ""
	pingTicker := time.NewTicker(pingPeriod)
	fileTicker := time.NewTicker(filePeriod)
	defer func() {
		pingTicker.Stop()
		fileTicker.Stop()
		ws.Close()
	}()

	for {
		select {
		case <-fileTicker.C:
			var data []byte
			var err error

			data, lastMod, err = readFileIfModified(lastMod)
			if err != nil {
				if s := err.Error(); s != lastError {
					lastError = s
					data = []byte(lastError)
				}
			} else {
				lastError = ""
			}

			if data != nil {
				ws.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.WriteMessage(websocket.TextMessage, data); err != nil {
					return
				}
			}
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func WSHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	var lastMod time.Time
	if n, err := strconv.ParseInt(r.FormValue("lastMod"), 16, 64); err == nil {
		lastMod = time.Unix(0, n)
	}

	go writer(ws, lastMod)
	reader(ws)
}
