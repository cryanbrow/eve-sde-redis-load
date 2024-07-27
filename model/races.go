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

var sdeRaces race

type race map[int]struct {
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	IconID int `yaml:"iconID" json:"iconID"`
	NameID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	ShipTypeID int         `yaml:"shipTypeID" json:"shipTypeID"`
	Skills     map[int]int `yaml:"skills" json:"skills"`
	ID         int         `json:"id"`
}

func LoadRaces(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeRaces)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeRaces {
		singleRace := sdeRaces[k]
		singleRace.ID = k
		singleRaceJSON, _ := json.Marshal(singleRace)
		redisKey := "race:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleRaceJSON, 0)
	}
}
