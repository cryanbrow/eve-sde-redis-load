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

var sdeTypeDogmas typeDogma

type typeDogma map[int]struct {
	DogmaAttributes []struct {
		AttributeID int     `yaml:"attributeID" json:"attributeID"`
		Value       float32 `yaml:"value" json:"value"`
	} `yaml:"dogmaAttributes" json:"dogmaAttributes"`
	DogmaEffects []struct {
		EffectID  int  `yaml:"effectID" json:"effectID"`
		IsDefault bool `yaml:"isDefault" json:"isDefault"`
	} `yaml:"dogmaEffects" json:"dogmaEffects"`
	ID int `json:"id"`
}

func LoadTypeDogmas(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeTypeDogmas)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeTypeDogmas {
		singletypeDogma := sdeTypeDogmas[k]
		singletypeDogma.ID = k
		singletypeDogmaJSON, _ := json.Marshal(singletypeDogma)
		redisKey := "typeDogma:" + strconv.Itoa(k)
		data.Rdb.Set(context.Background(), redisKey, singletypeDogmaJSON, 0)

	}
}
