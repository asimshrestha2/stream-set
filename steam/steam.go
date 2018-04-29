package steam

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"time"
)

//https://api.steampowered.com/ISteamApps/GetAppList/v2/

type (
	SteamResponse struct {
		AppList AppList `json:"applist"`
	}

	AppList struct {
		Apps []App `json:"apps"`
	}

	App struct {
		AppID int64  `json:"appid"`
		Name  string `json:"name"`
	}
)

const (
	steamgamesfile = "steamgames.json"
)

func Request(url string) (string, error) {
	fmt.Println("Steam Request: ", url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	rbody, _ := ioutil.ReadAll(resp.Body)

	return string(rbody), err
}

func SaveSteamGames(steamresponse interface{}) {
	_, err := os.Stat(steamgamesfile)
	if os.IsNotExist(err) {
		f, err := os.Create(steamgamesfile)
		if err != nil {
			log.Fatal(err)
			return
		}

		jf, _ := json.Marshal(steamresponse)

		if _, err := f.Write(jf); err != nil {
			log.Fatal(err)
			return
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		jf, _ := json.Marshal(steamresponse)
		err := ioutil.WriteFile(steamgamesfile, jf, 666)
		if err != nil {
			return
		}
	}
}

func GetApps() (SteamResponse, error) {
	sr := SteamResponse{}
	if stat, err := os.Stat(steamgamesfile); os.IsNotExist(err) || time.Now().Sub(stat.ModTime()).Hours() >= 24 {
		res, err := Request("https://api.steampowered.com/ISteamApps/GetAppList/v2/")
		if err != nil {
			return sr, err
		}
		if err := json.Unmarshal([]byte(res), &sr); err != nil {
			return sr, err
		}
		go SaveSteamGames(sr)
		return sr, nil
	}

	file, e := ioutil.ReadFile(steamgamesfile)
	if e != nil {
		return sr, e
	}
	if err := json.Unmarshal(file, &sr); err != nil {
		return sr, err
	}
	return sr, nil
}

func FindGameName(id int) string {
	sr, err := GetApps()
	if err != nil {
		return ""
	}

	i := sort.Search(len(sr.AppList.Apps), func(i int) bool { return int(sr.AppList.Apps[i].AppID) >= id })
	if i < len(sr.AppList.Apps) && int(sr.AppList.Apps[i].AppID) == id {
		fmt.Printf("found %d at index %d\n", id, i)
		return sr.AppList.Apps[i].Name
	}

	return ""
}
