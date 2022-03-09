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

var sdeStationOperations stationOperation

type stationOperation map[int]struct {
	ActivityID    int     `yaml:"activityID" json:"activityID"`
	Border        float32 `yaml:"border" json:"border"`
	Corridor      float32 `yaml:"corridor" json:"corridor"`
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	Fringe              float32 `yaml:"fringe" json:"fringe"`
	Hub                 float32 `yaml:"hub" json:"hub"`
	ManufacturingFactor float32 `yaml:"manufacturingFactor" json:"manufacturingFactor"`
	OperationNameID     struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"operationNameID" json:"operationNameID"`
	Ratio          float32     `yaml:"ratio" json:"ratio"`
	ResearchFactor float32     `yaml:"researchFactor" json:"researchFactor"`
	Services       []int       `yaml:"services" json:"services"`
	StationTypes   map[int]int `yaml:"stationTypes" json:"stationTypes"`
	ID             int         `json:"id"`
}

func LoadStationOperations(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeStationOperations)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeStationOperations {
		singleStationOperation := sdeStationOperations[k]
		singleStationOperation.ID = k
		singleStationOperationJSON, _ := json.Marshal(singleStationOperation)
		redisKey := "stationOperation:" + strconv.Itoa(k)
		data.Rdb.Set(context.Background(), redisKey, singleStationOperationJSON, 0)

	}
}
