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

var sdeTypeIDs typeID

type typeID map[int]struct {
	BasePrice   float64 `yaml:"de" json:"de"`
	Capacity    float32 `yaml:"capacity" json:"capacity"`
	Description struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"description" json:"description"`
	GraphicID     int           `yaml:"GraphicID" json:"GraphicID"`
	GroupID       int           `yaml:"GroupID" json:"GroupID"`
	IconID        int           `yaml:"IconID" json:"IconID"`
	MarketGroupID int           `yaml:"MarketGroupID" json:"MarketGroupID"`
	Masteries     map[int][]int `yaml:"Masteries" json:"Masteries"`
	MetaGroupID   int           `yaml:"MetaGroupID" json:"MetaGroupID"`
	Mass          string        `yaml:"Mass" json:"Mass"`
	Name          struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"name" json:"name"`
	PortionSize    int     `yaml:"PortionSize" json:"PortionSize"`
	Published      bool    `yaml:"Published" json:"Published"`
	RaceID         int     `yaml:"RaceID" json:"RaceID"`
	Radius         float32 `yaml:"Radius" json:"Radius"`
	SofFactionName string  `yaml:"SofFactionName" json:"SofFactionName"`
	SoundID        int     `yaml:"SoundID" json:"SoundID"`
	Traits         trait   `yaml:"Traits" json:"Traits"`
	Volume         float64 `yaml:"Volume" json:"Volume"`
	ID             int     `json:"id"`
}

type trait struct {
	IconID      int                 `yaml:"IconID" json:"IconID"`
	MiscBonuses []bonus             `yaml:"MiscBonuses" json:"MiscBonuses"`
	RoleBonuses []bonus             `yaml:"RoleBonuses" json:"RoleBonuses"`
	TraitTypes  map[int][]traitType `yaml:"TraitTypes" json:"TraitTypes"`
}

type bonus struct {
	BonusAmmount float32 `yaml:"bonus" json:"bonus"`
	BonusText    struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"bonusText" json:"bonusText"`
	Importance int  `yaml:"Importance" json:"Importance"`
	IsPositive bool `yaml:"IsPositive" json:"IsPositive"`
	UnitID     int  `yaml:"UnitID" json:"UnitID"`
}

type traitType struct {
	BonusAmount float32 `yaml:"bonus" json:"bonus"`
	BonusText   struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"bonusText" json:"bonusText"`
	Importance int `yaml:"Importance" json:"Importance"`
	UnitID     int `yaml:"UnitID" json:"UnitID"`
}

func LoadTypeIDs(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeTypeIDs)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeTypeIDs {
		singleTypeID := sdeTypeIDs[k]
		singleTypeID.ID = k
		singleTypeIDJSON, _ := json.Marshal(singleTypeID)
		redisKey := "typeID:" + strconv.Itoa(k)
		data.Rdb.Set(context.Background(), redisKey, singleTypeIDJSON, 0)

	}
}
