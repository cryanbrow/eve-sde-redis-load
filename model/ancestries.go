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

var sdeAncestries ancestry

type ancestry map[int]struct {
	BloodlineID   int `yaml:"bloodlineID" json:"bloodlineID"`
	Charisma      int `yaml:"charisma" json:"charisma"`
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
	Perception       int    `yaml:"perception" json:"perception"`
	ShortDescription string `yaml:"shortDescription" json:"shortDescription"`
	Willpower        int    `yaml:"willpower" json:"willpower"`
	ID               int    `json:"id"`
}

func LoadRedisAncestries(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeAncestries)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeAncestries {
		singleAncestry := sdeAncestries[k]
		singleAncestry.ID = k
		singleAncestryJSON, _ := json.Marshal(singleAncestry)
		redisKey := "ancestry:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleAncestryJSON, cache.NoExpiration)
	}
}
