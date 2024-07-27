package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cryanbrow/eve-sde-redis-load/data"
	"gopkg.in/yaml.v3"
)

var sdeDogmaAttributeCategories dogmaAttributeCategory

type dogmaAttributeCategory map[int]struct {
	Description string `yaml:"description" json:"description"`
	Name        string `yaml:"name" json:"name"`
	ID          int    `json:"id"`
}

func LoadRedisDogmaAttributeCategories(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeDogmaAttributeCategories)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeDogmaAttributeCategories {
		singleDogmaAttributeCategory := sdeDogmaAttributeCategories[k]
		singleDogmaAttributeCategory.ID = k
		singleDogmaAttributeCategoryJSON, _ := json.Marshal(singleDogmaAttributeCategory)
		redisKey := "dogmaAttributeCategory:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleDogmaAttributeCategoryJSON, 0)
	}
}
