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

var sdeDogmaEffects dogmaEffect

type dogmaEffect map[int]struct {
	DescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	DisallowAutoRepeat bool `yaml:"disallowAutoRepeat" json:"disallowAutoRepeat"`
	DisplayNameID      struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"displayNameID" json:"displayNameID"`
	DischargeAttributeID     int        `yaml:"dischargeAttributeID" json:"dischargeAttributeID"`
	Distribution             int        `yaml:"distribution" json:"distribution"`
	DurationAttributeID      int        `yaml:"durationAttributeID" json:"durationAttributeID"`
	EffectCategory           int        `yaml:"effectCategory" json:"effectCategory"`
	EffectID                 int        `yaml:"effectID" json:"effectID"`
	EffectName               string     `yaml:"effectName" json:"effectName"`
	ElectronicChance         bool       `yaml:"electronicChance" json:"electronicChance"`
	FalloffAttributeID       int        `yaml:"falloffAttributeID" json:"falloffAttributeID"`
	Guid                     string     `yaml:"guid" json:"guid"`
	IconID                   int        `yaml:"iconID" json:"iconID"`
	IsAssistance             bool       `yaml:"isAssistance" json:"isAssistance"`
	IsOffensive              bool       `yaml:"isOffensive" json:"isOffensive"`
	IsWarpSafe               bool       `yaml:"isWarpSafe" json:"isWarpSafe"`
	ModifierInfo             []modifier `yaml:"modifierInfo" json:"modifierInfo"`
	PropulsionChance         bool       `yaml:"propulsionChance" json:"propulsionChance"`
	Published                bool       `yaml:"published" json:"published"`
	RangeAttributeID         int        `yaml:"rangeAttributeID" json:"rangeAttributeID"`
	RangeChance              bool       `yaml:"rangeChance" json:"rangeChance"`
	TrackingSpeedAttributeID int        `yaml:"trackingSpeedAttributeID" json:"trackingSpeedAttributeID"`
	ID                       int        `json:"id"`
}

type modifier struct {
	Domain               string `yaml:"domain" json:"domain"`
	Function             string `yaml:"function" json:"function"`
	ModifiedAttributeID  int    `yaml:"modifiedAttributeID" json:"modifiedAttributeID"`
	ModifyingAttributeID int    `yaml:"modifyingAttributeID" json:"modifyingAttributeID"`
	Operation            int    `yaml:"operation" json:"operation"`
}

func LoadRedissdeDogmaEffects(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeDogmaEffects)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeDogmaEffects {
		singleDogmaEffect := sdeDogmaEffects[k]
		singleDogmaEffect.ID = k
		singleDogmaEffectJSON, _ := json.Marshal(singleDogmaEffect)
		redisKey := "dogmaEffect:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, singleDogmaEffectJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(singleDogmaEffectJSON))
	}
}
