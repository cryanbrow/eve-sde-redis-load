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

var sdeResearchAgents researchAgent

type researchAgent map[int]struct {
	Skills []struct {
		TypeID int `yaml:"typeID" json:"typeID"`
	} `yaml:"skils" json:"skills"`
	ID int `json:"id"`
}

func LoadResearchAgents(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeResearchAgents)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeResearchAgents {
		singleResearchAgent := sdeResearchAgents[k]
		singleResearchAgent.ID = k
		singleResearchAgentJSON, _ := json.Marshal(singleResearchAgent)
		redisKey := "researchAgent:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, singleResearchAgentJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(singleResearchAgentJSON))
	}
}
