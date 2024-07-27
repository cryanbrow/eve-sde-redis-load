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

var sdeGroupIDs groupID

type groupID map[int]struct {
	Anchorable           bool `yaml:"anchorable" json:"anchorable"`
	Anchored             bool `yaml:"anchored" json:"anchored"`
	CategoryID           int  `yaml:"categoryID" json:"categoryID"`
	FittableNonSingleton bool `yaml:"fittableNonSingleton" json:"fittableNonSingleton"`
	Name                 struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"name" json:"name"`
	Published    bool `yaml:"published" json:"published"`
	UseBasePrice bool `yaml:"useBasePrice" json:"useBasePrice"`
	ID           int  `json:"id"`
}

func LoadRedisGroupIDs(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeGroupIDs)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeGroupIDs {
		singleGroupID := sdeGroupIDs[k]
		singleGroupID.ID = k
		singleGroupIDJSON, _ := json.Marshal(singleGroupID)
		redisKey := "groupID:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleGroupIDJSON, 0)
	}
}
