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

var sdeIconIDS iconID

type iconID map[int]struct {
	Description string `yaml:"description" json:"description"`
	IconFile    string `yaml:"iconFile" json:"iconFile"`
	ID          int    `json:"id"`
}

func LoadRedisIconIDS(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeIconIDS)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeIconIDS {
		singleIconID := sdeIconIDS[k]
		singleIconID.ID = k
		singleIconIDJSON, _ := json.Marshal(singleIconID)
		redisKey := "iconID:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleIconIDJSON, 0)
	}
}
