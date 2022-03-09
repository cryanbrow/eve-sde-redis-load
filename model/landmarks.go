package model

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cryanbrow/eve-sde-redis-load/data"
	"gopkg.in/yaml.v3"
)

var sdeLandmarks landmark

type landmark map[int]struct {
	DescriptionID  int       `yaml:"descriptionID" json:"descriptionID"`
	IconID         int       `yaml:"iconID" json:"iconID"`
	LandmarkNameID int       `yaml:"landmarkNameID" json:"landmarkNameID"`
	LocationID     int       `yaml:"locationID" json:"locationID"`
	Position       []float64 `yaml:"position" json:"position"`
	ID             int       `json:"id"`
}

func LoadLandmarks(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeLandmarks)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeLandmarks {
		singleLandmark := sdeLandmarks[k]
		singleLandmark.ID = k
		singleLandmarkJSON, _ := json.Marshal(singleLandmark)
		redisKey := "landmark:" + strconv.Itoa(k)
		data.Rdb.Set(context.Background(), redisKey, singleLandmarkJSON, 0)

	}
}
