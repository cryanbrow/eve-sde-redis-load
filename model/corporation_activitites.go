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

var sdeCorporationActivities corporationActivity

type corporationActivity map[int]struct {
	NameID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	ID int `json:"id"`
}

func LoadRedisCorporationActivities(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeCorporationActivities)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeCorporationActivities {
		singleCorporationActivity := sdeCorporationActivities[k]
		singleCorporationActivity.ID = k
		sdeCorporationActivitiesJSON, _ := json.Marshal(singleCorporationActivity)
		redisKey := "corporationActivity:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, sdeCorporationActivitiesJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(sdeCorporationActivitiesJSON))
	}
}
