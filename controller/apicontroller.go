package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type (
	Return struct {
		Err  bool   `json:"err"`
		Data string `json:"data"`
	}
)

// APIIndex for Website
func APIIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	u := Return{
		Err:  false,
		Data: "Hi",
	}

	uj, _ := json.Marshal(u)

	fmt.Fprintf(w, "%s", uj)
}

func APIGetFilePath(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	// reg := regexp.MustCompile(`\\`)
	// ret := reg.ReplaceAllString(sf.FileLocation, "\\\\")
	retj, _ := json.Marshal(sf)
	fmt.Fprintf(w, "%s", retj)
}

//data: "circlecolor=" + color_input.value + "&timersetting=" + time_input.value + "&logopath=" + logopath_input.value,

func APISettings(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	if r.FormValue("filepath") != "" {
		sf.FileLocation = r.FormValue("filepath")
	}
	if r.FormValue("circlecolor") != "" {
		sf.CircleColor = r.FormValue("circlecolor")
	}
	if r.FormValue("timersetting") != "" {
		t, err := strconv.ParseInt(r.FormValue("timersetting"), 10, 64)
		sf.TimerSetting = t
		if err != nil {
			sf.TimerSetting = 60000
		}
	}
	if r.FormValue("logopath") != "" {
		sf.LogoPath = r.FormValue("logopath")
	}
	SaveDataToFile()
}
