package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

type (
	SaveFile struct {
		FileLocation string `json:"filelocation"`
		CircleColor  string `json:"circlecolor"`
		TimerSetting int64  `json:"timersetting"`
		LogoPath     string `json:"logopath"`
	}
)

var (
	sf = SaveFile{
		FileLocation: "",
		CircleColor:  "#212121",
		TimerSetting: 60000,
		LogoPath:     "",
	}
	userHomeDir string
	appLocalDir string
	appLocal    string
)

func SaveDataToFile() {
	if userHomeDir == "" {
		getHomeDirectory()
	}
	sfj, _ := json.Marshal(sf)

	if _, err := os.Stat(appLocal); err == nil {
		ioutil.WriteFile(appLocal, sfj, 0644)
		log.Println("File Saved")
	} else {
		err := os.MkdirAll(appLocalDir, 0755)
		if err == nil {
			ioutil.WriteFile(appLocal, sfj, 0644)
			log.Println("File Saved")
		} else {
			log.Fatal(err)
		}
	}

}

func GetDataFromFile() {
	if userHomeDir == "" {
		getHomeDirectory()
	}

	if _, err := os.Stat(appLocal); err == nil {
		log.Println("Saved File Found")
		data, err := ioutil.ReadFile(appLocal)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(data, &sf)
		log.Printf("Last Saved Information: %v", sf)
	}
}

func getHomeDirectory() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	userHomeDir = usr.HomeDir
	appLocalDir = path.Join(userHomeDir, "AppData\\Local\\stream-set")
	appLocal = path.Join(userHomeDir, "AppData\\Local\\stream-set\\data.json")
}
