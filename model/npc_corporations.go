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

var sdeNPCCorporations npcCorporation

type npcCorporation map[int]struct {
	AllowedMemberRaces []int           `yaml:"allowedMemberRaces" json:"allowedMemberRaces"`
	CeoID              int             `yaml:"ceoID" json:"ceoID"`
	CorporationTrades  map[int]float32 `yaml:"corporationTrades" json:"corporationTrades"`
	Deleted            bool            `yaml:"deleted" json:"deleted"`
	DescriptionID      struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"descriptionID" json:"descriptionID"`
	Divisions                 map[int]division `yaml:"divisions" json:"divisions"`
	EnemyID                   int              `yaml:"enemyID" json:"enemyID"`
	Extent                    string           `yaml:"extent" json:"extent"`
	FactionID                 int              `yaml:"factionID" json:"factionID"`
	FriendID                  int              `yaml:"friendID" json:"friendID"`
	HasPlayerPersonnelManager bool             `yaml:"hasPlayerPersonnelManager" json:"hasPlayerPersonnelManager"`
	IconID                    int              `yaml:"iconID" json:"iconID"`
	InitialPrice              int              `yaml:"initialPrice" json:"initialPrice"`
	Investors                 map[int]int      `yaml:"investors" json:"investors"`
	LPOfferTables             []int            `yaml:"lpOfferTables" json:"lpOfferTables"`
	MainActivityID            int              `yaml:"mainActivityID" json:"mainActivityID"`
	MemberLimit               int              `yaml:"memberLimit" json:"memberLimit"`
	MinSecurity               float32          `yaml:"minSecurity" json:"minSecurity"`
	MinimumJoinStanding       int              `yaml:"minimumJoinStanding" json:"minimumJoinStanding"`
	NameID                    struct {
		DE string `yaml:"de" json:"de"`
		EN string `yaml:"en" json:"en"`
		FR string `yaml:"fr" json:"fr"`
		JA string `yaml:"ja" json:"ja"`
		KO string `yaml:"ko" json:"ko"`
		RU string `yaml:"ru" json:"ru"`
		ZH string `yaml:"zh" json:"zh"`
	} `yaml:"nameID" json:"nameID"`
	PublicShares               int     `yaml:"publicShares" json:"publicShares"`
	RaceID                     int     `yaml:"raceID" json:"raceID"`
	SendCharTerminationMessage bool    `yaml:"sendCharTerminationMessage" json:"sendCharTerminationMessage"`
	Shares                     int     `yaml:"shares" json:"shares"`
	Size                       string  `yaml:"size" json:"size"`
	SizeFactor                 float32 `yaml:"sizeFactor" json:"sizeFactor"`
	SolarSystemID              int     `yaml:"solarSystemID" json:"solarSystemID"`
	StationID                  int     `yaml:"stationID" json:"stationID"`
	TaxRate                    float32 `yaml:"taxRate" json:"taxRate"`
	TickerName                 string  `yaml:"tickerName" json:"tickerName"`
	UniqueName                 bool    `yaml:"uniqueName" json:"uniqueName"`
	ID                         int     `json:"id"`
}

type division struct {
	DivisionNumber int `yaml:"divisionNumber" json:"divisionNumber"`
	LeaderID       int `yaml:"leaderID" json:"leaderID"`
	Size           int `yaml:"size" json:"size"`
}

func LoadNPCCorporations(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeNPCCorporations)
	if err != nil {
		fmt.Println(err.Error())
	}

	for k := range sdeNPCCorporations {
		singleNPCCorporation := sdeNPCCorporations[k]
		singleNPCCorporation.ID = k
		singleNPCCorporationJSON, _ := json.Marshal(singleNPCCorporation)
		redisKey := "npcCorporation:" + strconv.Itoa(k)
		status := data.Rdb.Set(context.Background(), redisKey, singleNPCCorporationJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(singleNPCCorporationJSON))
	}
}
