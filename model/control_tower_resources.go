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

var sdeControlTowerAttributes controlTowerAttribute

type controlTowerAttribute map[int]struct {
	Resources []resource `yaml:"resource" json:"resource"`
	ID        int        `json:"id"`
}

type resource struct {
	FactionID        int     `yaml:"factionID" json:"factionID"`
	MinSecurityLevel float32 `yaml:"minSecurityLevel" json:"minSecurityLevel"`
	Purpose          int     `yaml:"purpose" json:"purpose"`
	Quantity         int     `yaml:"quantity" json:"quantity"`
	ResourceTypeID   int     `yaml:"resourceTypeID" json:"resourceTypeID"`
}

func LoadRedisControlTowerAttributes(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeControlTowerAttributes)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeControlTowerAttributes {
		singleControlTowerAttribute := sdeControlTowerAttributes[k]
		singleControlTowerAttribute.ID = k
		sdeControlTowerAttributesJSON, _ := json.Marshal(singleControlTowerAttribute)
		redisKey := "controlTowerAttribute:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, sdeControlTowerAttributesJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(sdeControlTowerAttributesJSON))
	}
}
