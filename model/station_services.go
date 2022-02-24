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

var sdeStationServices stationService

type stationService map[int]struct {
	ServiceNameID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"serviceNameID" json:"serviceNameID"`
	ID int `json:"id"`
}

func LoadStationServices(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeStationServices)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeStationServices {
		singleStationService := sdeStationServices[k]
		singleStationService.ID = k
		singleStationServiceJSON, _ := json.Marshal(singleStationService)
		redisKey := "stationService:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, singleStationServiceJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(singleStationServiceJSON))
	}
}
