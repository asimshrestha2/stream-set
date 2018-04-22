package save

import (
	"strings"
)

// const (
// 	ignorelistfile = "ignorelist.txt"
// )

func GetIgnoreList() []string {
	return strings.Split(CFG.Section("list").Key("ignore").String(), ", ")
}
