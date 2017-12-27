package controller

import (
	"encoding/json"
	"fmt"
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
)

func SaveDataToFile() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr.HomeDir)

	applocal := path.Join(usr.HomeDir, "AppData\\Local\\stream-set\\data.json")
	sfj, _ := json.Marshal(sf)

	if _, err := os.Stat(applocal); err == nil {
		fmt.Println("File Exist")
		ioutil.WriteFile(applocal, sfj, 0644)
	} else {
		fmt.Println("File Doesn't Exist")
		err := os.MkdirAll(path.Join(usr.HomeDir, "AppData\\Local\\stream-set"), 0755)
		if err == nil {
			ioutil.WriteFile(applocal, sfj, 0644)
		} else {
			log.Fatal(err)
		}
	}

}

func GetDataFromFile() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr.HomeDir)
	applocal := path.Join(usr.HomeDir, "AppData\\Local\\stream-set\\data.json")

	if _, err := os.Stat(applocal); err == nil {
		fmt.Println("File Exist")
		data, err := ioutil.ReadFile(applocal)
		if err != nil {
			log.Fatal(err)
		}

		json.Unmarshal(data, &sf)
		fmt.Println(sf)
	}
}
