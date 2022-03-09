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

var sdeTypeMaterials typeMaterial

type typeMaterial map[int]struct {
	Materials []struct {
		MaterialTypeID int `yaml:"materialTypeID" json:"materialTypeID"`
		Quantity       int `yaml:"quantity" json:"quantity"`
	} `yaml:"materials" json:"materials"`
	ID int `json:"id"`
}

func LoadTypeMaterials(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeTypeMaterials)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeTypeMaterials {
		singleTypeMaterial := sdeTypeMaterials[k]
		singleTypeMaterial.ID = k
		singleTypeMaterialJSON, _ := json.Marshal(singleTypeMaterial)
		redisKey := "typeMaterial:" + strconv.Itoa(k)
		data.Rdb.Set(context.Background(), redisKey, singleTypeMaterialJSON, 0)

	}
}
