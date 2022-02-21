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

var sdeCertificates certificate

type certificate map[int]struct {
	Description    string `yaml:"description" json:"description"`
	GroupID        string `yaml:"groupID" json:"groupID"`
	Name           string `yaml:"name" json:"name"`
	RecommendedFor []int  `yaml:"recommendedFor" json:"recommendedFor"`
	SkillTypes     struct {
		SkillType map[int]struct {
			Advanced int `yaml:"advanced" json:"advanced"`
			Basic    int `yaml:"basic" json:"basic"`
			Elite    int `yaml:"elite" json:"elite"`
			Improved int `yaml:"improved" json:"improved"`
			Standard int `yaml:"standard" json:"standard"`
		}
	}
	ID int `json:"id"`
}

func LoadRedisCertificates(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeCertificates)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeCertificates {
		singleCertificate := sdeCertificates[k]
		singleCertificate.ID = k
		singleCertificateJSON, _ := json.Marshal(singleCertificate)
		redisKey := "certificate:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, singleCertificateJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(singleCertificateJSON))
	}
}
