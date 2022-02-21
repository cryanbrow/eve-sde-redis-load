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

var sdeCategoryIDs categoryID

type categoryID map[int]struct {
	Name struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"name" json:"name"`
	Published bool `yaml:"published" json:"published"`
	ID        int  `json:"id"`
}

func LoadRedisCategoryIDs(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeCategoryIDs)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeCategoryIDs {
		singleCategoryID := sdeCategoryIDs[k]
		singleCategoryID.ID = k
		sdeCategoryIDsJSON, _ := json.Marshal(singleCategoryID)
		redisKey := "categoryID:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, sdeCategoryIDsJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(sdeCategoryIDsJSON))
	}
}
