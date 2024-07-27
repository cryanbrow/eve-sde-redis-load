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

var sdePlanetSchematics planetSchematic

type planetSchematic map[int]struct {
	CycleTime int `yaml:"cycleTime" json:"cycleTime"`
	NameID    struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	Pins  []int                `yaml:"pins" json:"pins"`
	Types map[int]resourceType `yaml:"types" json:"types"`
	ID    int                  `json:"id"`
}

type resourceType struct {
	IsInput  bool `yaml:"isInput" json:"isInput"`
	Quantity int  `yaml:"quantity" json:"quantity"`
}

func LoadPlanetSchematics(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdePlanetSchematics)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdePlanetSchematics {
		singlePlanetSchematic := sdePlanetSchematics[k]
		singlePlanetSchematic.ID = k
		singlePlanetSchematicJSON, _ := json.Marshal(singlePlanetSchematic)
		redisKey := "planetSchematic:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singlePlanetSchematicJSON, 0)
	}
}
