package save

import (
	"encoding/json"
	"log"
	"os"
)

var (
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
	}
}

func LoadGameList() {

}

func GameListExist() bool {
	if _, err := os.Stat(gamelistfile); os.IsNotExist(err) {
		return false
	}
	return true
}
