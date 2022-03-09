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

var sdeAgentsInSpace agentInSpace

type agentInSpace map[int]struct {
	DungeonID     int `yaml:"dungeonID" json:"dungeonID"`
	SolarSystemID int `yaml:"solarSystemID" json:"solarSystemID"`
	SpawnPointID  int `yaml:"spawnPointID" json:"spawnPointID"`
	TypeID        int `yaml:"typeID" json:"typeID"`
	ID            int `json:"id"`
}

func LoadRedisAgentsInSpace(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeAgentsInSpace)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeAgentsInSpace {
		singleAgent := sdeAgentsInSpace[k]
		singleAgent.ID = k
		singleAgentInSpaceJSON, _ := json.Marshal(singleAgent)
		redisKey := "agentInSpace:" + strconv.Itoa(k)
		data.Rdb.Set(context.Background(), redisKey, singleAgentInSpaceJSON, 0)
	}
}
