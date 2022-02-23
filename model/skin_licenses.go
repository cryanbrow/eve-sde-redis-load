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

var sdeSkinLicenses skinLicense

type skinLicense map[int]struct {
	Duration      int `yaml:"duration" json:"duration"`
	LicenseTypeID int `yaml:"licenseTypeID" json:"licenseTypeID"`
	SkinID        int `yaml:"skinID" json:"skinID"`
	ID            int `json:"id"`
}

func LoadSkinLicenses(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeSkinLicenses)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeSkinLicenses {
		singleSkinLicense := sdeSkinLicenses[k]
		singleSkinLicense.ID = k
		singleSkinLicenseJSON, _ := json.Marshal(singleSkinLicense)
		redisKey := "skinLicense:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, singleSkinLicenseJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(singleSkinLicenseJSON))
	}
}
