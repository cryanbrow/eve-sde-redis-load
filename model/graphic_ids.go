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

var sdeGraphicIDs graphicID

type graphicID map[int]struct {
	AttributeID          int     `yaml:"attributeID" json:"attributeID"`
	CategoryID           int     `yaml:"categoryID" json:"categoryID"`
	ChargeRechargeTimeID int     `yaml:"chargeRechargeTimeID" json:"chargeRechargeTimeID"`
	DataType             int     `yaml:"dataType" json:"dataType"`
	DefaultValue         float32 `yaml:"defaultValue" json:"defaultValue"`
	Description          string  `yaml:"description" json:"description"`
	DisplayNameID        struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"displayNameID" json:"displayNameID"`
	HighIsGood           bool   `yaml:"highIsGood" json:"highIsGood"`
	IconID               int    `yaml:"iconID" json:"iconID"`
	MaxAttributeID       int    `yaml:"maxAttributeID" json:"maxAttributeID"`
	Name                 string `yaml:"name" json:"name"`
	Published            bool   `yaml:"published" json:"published"`
	Stackable            bool   `yaml:"stackable" json:"stackable"`
	TooltipDescriptionID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"tooltipDescriptionID" json:"tooltipDescriptionID"`
	TooltipTitleID struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"tooltipTitleID" json:"tooltipTitleID"`
	UnitID int `yaml:"unitID" json:"unitID"`
	ID     int `json:"id"`
}

func LoadRedisGraphicIDs(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeGraphicIDs)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeGraphicIDs {
		singleGraphicID := sdeGraphicIDs[k]
		singleGraphicID.ID = k
		singleGraphicIDJSON, _ := json.Marshal(singleGraphicID)
		redisKey := "graphicID:" + strconv.Itoa(k)
		data.Rdb.Set(context.Background(), redisKey, singleGraphicIDJSON, 0)

	}
}
