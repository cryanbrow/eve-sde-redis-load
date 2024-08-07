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

var sdeAgents agent

type agent map[int]struct {
	AgentTypeID   int  `yaml:"agentTypeID" json:"agentTypeID"`
	CorporationID int  `yaml:"corporationID" json:"corporationID"`
	DivisionID    int  `yaml:"divisionID" json:"divisionID"`
	IsLocator     bool `yaml:"isLocator" json:"isLocator"`
	Level         int  `yaml:"level" json:"level"`
	LocationID    int  `yaml:"locationID" json:"locationID"`
	ID            int  `json:"id"`
}

func LoadRedisAgents(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeAgents)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeAgents {
		singleAgent := sdeAgents[k]
		singleAgent.ID = k
		singleAgentJSON, _ := json.Marshal(singleAgent)
		redisKey := "agent:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleAgentJSON, cache.NoExpiration)
	}
}
