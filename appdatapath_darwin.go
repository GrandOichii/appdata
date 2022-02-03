package appdata

import (
	"os"
	"path"
)

var appDataPath string = "/Library/Application Support"

func init() {
	dirname, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	appDataPath = path.Join(dirname, appDataPath)
}
