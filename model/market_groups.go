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

var sdeMarketGroups marketGroup

type marketGroup map[int]struct {
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	HasTypes bool `yaml:"hasTypes" json:"hasTypes"`
	IconID   int  `yaml:"iconID" json:"iconID"`
	NameID   struct {
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

func LoadMarketGroups(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeMarketGroups)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeMarketGroups {
		singleMarketGroup := sdeMarketGroups[k]
		singleMarketGroup.ID = k
		singleMarketGroupJSON, _ := json.Marshal(singleMarketGroup)
		redisKey := "marketGroup:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleMarketGroupJSON, 0)
	}
}
