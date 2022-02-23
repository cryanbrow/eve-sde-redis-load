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

var sdeFactions faction

type faction map[int]struct {
	CorporationID int `yaml:"corporationID" json:"corporationID"`
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	IconID               int   `yaml:"iconID" json:"iconID"`
	MemberRaces          []int `yaml:"memberRaces" json:"memberRaces"`
	MilitiaCorporationID int   `yaml:"militiaCorporationID" json:"militiaCorporationID"`
	NameID               struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	ShortDescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"shortDescriptionID" json:"shortDescriptionID"`
	SizeFactor    float32 `yaml:"sizeFactor" json:"sizeFactor"`
	SolarSystemID int     `yaml:"solarSystemID" json:"solarSystemID"`
	UniqueName    bool    `yaml:"uniqueName" json:"uniqueName"`
	ID            int     `json:"id"`
}

func LoadRedisFactions(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeFactions)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeFactions {
		singleFaction := sdeFactions[k]
		singleFaction.ID = k
		singleFactionJSON, _ := json.Marshal(singleFaction)
		redisKey := "faction:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, singleFactionJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(singleFactionJSON))
	}
}
