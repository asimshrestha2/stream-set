package save

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const (
	gamelistfile = "gamelist.json"
)

func SaveGameList(savestruct interface{}) {
	if !GameListExist() {
		f, err := os.Create(gamelistfile)
		if err != nil {
			log.Fatal(err)
			return
		}

		jf, _ := json.Marshal(savestruct)

		if _, err := f.Write(jf); err != nil {
			log.Fatal(err)
			return
		}
		if err := f.Close(); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		jf, _ := json.Marshal(savestruct)
		err := ioutil.WriteFile(gamelistfile, jf, 666)
		if err != nil {
			return
		}
	}
}

func RemoveGameList() error {
	if GameListExist() {
		err := os.Remove(gamelistfile)
		if err != nil {
			return err
		}
	}
	return nil
}

func LoadGameList(t interface{}) error {
	if GameListExist() {
		f, err := ioutil.ReadFile(gamelistfile)
		if err != nil {
			return err
		}

		json.Unmarshal(f, &t)
	}
	return nil
}

func GameListExist() bool {
	if _, err := os.Stat(gamelistfile); os.IsNotExist(err) {
		return false
	}
	return true
}
