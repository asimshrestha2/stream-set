package save

import (
	"fmt"

	"gopkg.in/ini.v1"
)

const (
	configFile = "settings.ini"
)

var CFG *ini.File

func LoadSettings() {
	var err error
	CFG, err = ini.Load(configFile)
	if err != nil {
		fmt.Println("Error: ", err)
		SetUpSettings()
	}
}

func SetUpSettings() {
	cfg := ini.Empty()
	cfg.Section("list").Key("ignore").SetValue("Firefox, Google Chrome, Discord, Steam, Blizzard Battle.net, Epic Games Launcher, Stream Set")
	cfg.Section("twitch").Key("defaultGame").SetValue("IRL")
	cfg.Section("twitch").Key("waitToReset").SetValue("500")
	cfg.SaveTo(configFile)
	CFG = cfg
}

func GetTwitchDefaultGame() string {
	temp := CFG.Section("twitch").Key("defaultGame").String()
	if temp != "" {
		return temp
	}
	return "IRL"
}

func GetResetTime() float64 {
	temp, err := CFG.Section("twitch").Key("waitToReset").Float64()
	if err != nil {
		fmt.Println("Error: ", err)
		return 500
	}
	return temp
}
