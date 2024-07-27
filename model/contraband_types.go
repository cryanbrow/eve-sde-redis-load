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

var sdeContrabandTypes contrabandType

type contrabandType map[int]struct {
	Factions map[int]struct {
		AttackMinSec     float32 `yaml:"attackMinSec" json:"attackMinSec"`
		ConfiscateMinSec float32 `yaml:"confiscateMinSec" json:"confiscateMinSec"`
		FineByValue      float32 `yaml:"fineByValue" json:"fineByValue"`
		StandingLoss     float32 `yaml:"standingLoss" json:"standingLoss"`
	}
	ID int `json:"id"`
}

func LoadRediscontrabandTypes(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeContrabandTypes)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeContrabandTypes {
		singleContrabandType := sdeContrabandTypes[k]
		singleContrabandType.ID = k
		singleContrabandTypeJSON, _ := json.Marshal(singleContrabandType)
		redisKey := "contrabandType:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleContrabandTypeJSON, cache.NoExpiration)
	}
}
