package server

import (
	"flag"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/julienschmidt/httprouter"
)

func StartServer() {
	var dir string
	flag.StringVar(&dir, "dir", "./server/static/", "the directory to serve files from. Defaults to the static folder")
	GetDataFromFile()
	flag.Parse()

	r := httprouter.New()
	r.GET("/", Index)
	// r.GET("/notification/", Notification)
	// r.GET("/settings/", Settings)
	// r.GET("/api/", APIIndex)
	// r.GET("/api/getfilepath", APIGetFilePath)
	// r.POST("/api/setsettings", APISettings)

	// Twitch API
	r.GET("/twitch/token", TwitchTokenAPI)
	r.GET("/twitch/gamelist", TwitchGameListAPI)

	// Gamelist
	r.GET("/gamelist", Gamelist)
	r.POST("/gamelist", GamelistPost)

	// r.GET("/fileimage", FileImage)
	// r.GET("/ws", WSHandler)
	r.ServeFiles("/static/*filepath", http.Dir(dir))

	// n := negroni.Classic() // Includes some default middlewares
	// n.UseHandler(r)

	n := negroni.New()
	n.Use(negroni.HandlerFunc(SetXPoweredBy))
	n.UseHandler(r)

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
	}
	log.Fatal(srv.ListenAndServe())
}

func SetXPoweredBy(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rw.Header().Set("x-powered-by", "Potato")
	next(rw, r)
	// do some stuff after
}
