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

var sdeSkins skin

type skin map[int]struct {
	AllowCCPDevs       bool   `yaml:"allowCCPDevs" json:"allowCCPDevs"`
	InternalName       string `yaml:"internalName" json:"internalName"`
	IsStructureSkin    bool   `yaml:"isStructureSkin" json:"isStructureSkin"`
	SkinID             int    `yaml:"skinID" json:"skinID"`
	SkinMaterialID     int    `yaml:"skinMaterialID" json:"skinMaterialID"`
	Types              []int  `yaml:"types" json:"types"`
	VisibleSerenity    bool   `yaml:"visibleSerenity" json:"visibleSerenity"`
	VisibleTranquility bool   `yaml:"visibleTranquility" json:"visibleTranquility"`
	ID                 int    `json:"id"`
}

func LoadSkins(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeSkins)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeSkins {
		singleSkin := sdeSkins[k]
		singleSkin.ID = k
		singleSkinJSON, _ := json.Marshal(singleSkin)
		redisKey := "skin:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, singleSkinJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(singleSkinJSON))
	}
}
