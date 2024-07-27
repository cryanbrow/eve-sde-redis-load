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

var sdeUniqueNames []uniqueNames

type uniqueNames struct {
	GroupID  int    `yaml:"groupID" json:"groupID"`
	ItemID   int    `yaml:"itemID" json:"itemID"`
	ItemName string `yaml:"itemName" json:"itemName"`
}

var names = make(map[int]nameForID)
var ids = make(map[string]idForName)

type nameForID struct {
	GroupID  int    `yaml:"groupID" json:"groupID"`
	ItemName string `yaml:"itemName" json:"itemName"`
}

type idForName struct {
	GroupID int `yaml:"groupID" json:"groupID"`
	ItemID  int `yaml:"itemID" json:"itemID"`
}

func LoadUniqueNames(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeUniqueNames)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, element := range sdeUniqueNames {
		var localname nameForID
		localname.GroupID = element.GroupID
		localname.ItemName = element.ItemName
		names[element.ItemID] = localname

		var localID idForName
		localID.GroupID = element.GroupID
		localID.ItemID = element.ItemID
		ids[element.ItemName] = localID

		nameJSON, _ := json.Marshal(localname)
		redisKey := "name:" + strconv.Itoa(element.ItemID)
		data.NonExpiringCache.Set(redisKey, nameJSON, 0)

		idJSON, _ := json.Marshal(localID)
		redisKey = "id:" + strconv.Itoa(element.ItemID)
		data.NonExpiringCache.Set(redisKey, idJSON, 0)
	}

}
