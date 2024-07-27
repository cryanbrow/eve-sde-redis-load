package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cryanbrow/eve-sde-redis-load/data"
	"gopkg.in/yaml.v3"
)

var sdeInvNames []invName

type invName struct {
	ItemID   int    `yaml:"itemID" json:"item_id"`
	ItemName string `yaml:"itemName" json:"item_name"`
}

var invNames = make(map[int]string)
var invIds = make(map[string]int)

func LoadInvNames(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeInvNames)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, element := range sdeInvNames {
		invNames[element.ItemID] = element.ItemName

		invIds[element.ItemName] = element.ItemID

		redisKey := "invName:" + element.ItemName
		data.NonExpiringCache.Set(redisKey, element.ItemName, 0)

		redisKey = "invID:" + strconv.Itoa(element.ItemID)
		data.NonExpiringCache.Set(redisKey, element.ItemName, 0)
	}

}
