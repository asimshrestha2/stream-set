package save

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	ignorelistfile = "ignorelist.txt"
)

func GetIgnoreList() []string {
	if _, err := os.Stat(ignorelistfile); os.IsNotExist(err) {
		log.Printf("No %s found.\n", ignorelistfile)
		return nil
	}

	file, err := ioutil.ReadFile(ignorelistfile)
	if err != nil {
		log.Println("Can't read ingnore list")
		return nil
	}

	return strings.Split(string(file), "\n")
}
