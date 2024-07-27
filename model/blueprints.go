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

var sdeBlueprints blueprint

type blueprint map[int]struct {
	Activites struct {
		Copying struct {
			Time int `yaml:"time" json:"time"`
		} `yaml:"copying" json:"copying"`
		Invention struct {
			Materials []material `yaml:"materials" json:"materials"`
			Products  []product  `yaml:"products" json:"products"`
			Skills    []skill    `yaml:"skills" json:"skills"`
			Time      int        `yaml:"time" json:"time"`
		} `yaml:"invention" json:"invention"`
		Manufacturing struct {
			Materials []material `yaml:"materials" json:"materials"`
			Products  []product  `yaml:"products" json:"products"`
			Skills    []skill    `yaml:"skills" json:"skills"`
			Time      int        `yaml:"time" json:"time"`
		} `yaml:"manufacturing" json:"manufacturing"`
		ResearchMaterial struct {
			Time int `yaml:"time" json:"time"`
		} `yaml:"research_material" json:"research_material"`
		ResearchTime struct {
			Time int `yaml:"time" json:"time"`
		} `yaml:"research_time" json:"research_time"`
	} `yaml:"activities" json:"activities"`
	BlueprintTypeID    int `yaml:"blueprintTypeID" json:"blueprintTypeID"`
	MaxProductionLimit int `yaml:"maxProductionLimit" json:"maxProductionLimit"`
	ID                 int `json:"id"`
}

type material struct {
	Quantity int `yaml:"quantity" json:"quantity"`
	TypeID   int `yaml:"typeID" json:"typeID"`
}

type product struct {
	Probability float32 `yaml:"probability" json:"probability"`
	Quantity    int     `yaml:"quantity" json:"quantity"`
	TypeID      int     `yaml:"typeID" json:"typeID"`
}

type skill struct {
	Level  int `yaml:"level" json:"level"`
	TypeID int `yaml:"typeID" json:"typeID"`
}

func LoadRedisBlueprints(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeBlueprints)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeBlueprints {
		singleBlueprint := sdeBlueprints[k]
		singleBlueprint.ID = k
		sdeBlueprintsJSON, _ := json.Marshal(singleBlueprint)
		redisKey := "blueprint:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, sdeBlueprintsJSON, cache.NoExpiration)
	}
}
