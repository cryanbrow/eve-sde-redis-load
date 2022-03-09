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

var sdeMetaGroups metaGroup

type metaGroup map[int]struct {
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	IconID     int    `yaml:"iconID" json:"iconID"`
	IconSuffix string `yaml:"iconSuffix" json:"iconSuffix"`
	NameID     struct {
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

func LoadMetaGroups(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeMetaGroups)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeMetaGroups {
		singleMetaGroup := sdeMetaGroups[k]
		singleMetaGroup.ID = k
		singleMetaGroupJSON, _ := json.Marshal(singleMetaGroup)
		redisKey := "metaGroup:" + strconv.Itoa(k)
		data.Rdb.Set(context.Background(), redisKey, singleMetaGroupJSON, 0)

	}
}
