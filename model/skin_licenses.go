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
		data.NonExpiringCache.Set(redisKey, singleSkinLicenseJSON, 0)
	}
}
