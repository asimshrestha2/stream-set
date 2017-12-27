package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/urfave/negroni"

	"github.com/asimshrestha2/stream-set/controller"
	"github.com/julienschmidt/httprouter"
)

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//     fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

func main() {
	var dir string
	flag.StringVar(&dir, "dir", "./static/", "the directory to serve files from. Defaults to the static folder")
	controller.GetDataFromFile()
	flag.Parse()

	r := httprouter.New()
	r.GET("/", controller.Index)
	r.GET("/notification/", controller.Notification)
	r.GET("/settings/", controller.Settings)
	r.GET("/api/", controller.APIIndex)
	r.GET("/api/getfilepath", controller.APIGetFilePath)
	r.POST("/api/setsettings", controller.APISettings)

	// r.GET("/hello/:name", Hello)
	r.GET("/ws", controller.WSHandler)
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
