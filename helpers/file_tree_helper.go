package helpers

import (
	"io/ioutil"
	"log"
)

func ReturnDirNames(path string) []string {
	files, err := ioutil.ReadDir(path)
	var dirNameArray []string

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			dirNameArray = append(dirNameArray, file.Name())
		}
	}
	return dirNameArray
}
