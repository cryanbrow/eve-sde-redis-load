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

var sdeSkinMaterials skinMaterial

type skinMaterial map[int]struct {
	DisplayNameID  int `yaml:"displayNameID" json:"displayNameID"`
	M3aterialSetID int `yaml:"materialSetID" json:"materialSetID"`
	SkinMaterialID int `yaml:"skinMaterialID" json:"skinMaterialID"`
	ID             int `json:"id"`
}

func LoadSkinMaterials(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeSkinMaterials)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeSkinMaterials {
		singleSkinMaterial := sdeSkinMaterials[k]
		singleSkinMaterial.ID = k
		singleSkinMaterialJSON, _ := json.Marshal(singleSkinMaterial)
		redisKey := "skinMaterial:" + strconv.Itoa(k)
		data.NonExpiringCache.Set(redisKey, singleSkinMaterialJSON, 0)
	}
}
