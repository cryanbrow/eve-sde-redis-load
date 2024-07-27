package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cryanbrow/eve-sde-redis-load/data"
	"github.com/patrickmn/go-cache"
	"gopkg.in/yaml.v3"
)

var sdeBloodlines bloodline

type bloodline map[int]struct {
	Charisma      int `yaml:"charisma" json:"charisma"`
	CorporationID int `yaml:"corporationID" json:"corporationID"`
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	IconID       int `yaml:"iconID" json:"iconID"`
	Intelligence int `yaml:"intelligence" json:"intelligence"`
	Memory       int `yaml:"memory" json:"memory"`
	NameID       struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	Perception int `yaml:"perception" json:"perception"`
	RaceID     int `yaml:"raceID" json:"raceID"`
	Willpower  int `yaml:"willpower" json:"willpower"`
	ID         int `json:"id"`
}

func LoadRedisBloodlines(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeBloodlines)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeBloodlines {
		singleBloodline := sdeBloodlines[k]
		singleBloodline.ID = k
		singleBloodlineJSON, _ := json.Marshal(singleBloodline)
		redisKey := "agent:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleBloodlineJSON, cache.NoExpiration)
	}
}
