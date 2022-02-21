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

var sdeCharacterAttributes characterAttribute

type characterAttribute map[int]struct {
	Description string `yaml:"description" json:"description"`
	IconID      int    `yaml:"iconID" json:"iconID"`
	NameID      struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	Notes            string `yaml:"notes" json:"notes"`
	ShortDescription string `yaml:"shortDescription" json:"shortdescription"`
	ID               int    `json:"id"`
}

func LoadRedisCharacterAttributes(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeCharacterAttributes)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeCharacterAttributes {
		singleCharacterAttribute := sdeCharacterAttributes[k]
		singleCharacterAttribute.ID = k
		sdeCharacterAttributesJSON, _ := json.Marshal(singleCharacterAttribute)
		redisKey := "characterAttribute:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, sdeCharacterAttributesJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(sdeCharacterAttributesJSON))
	}
}
