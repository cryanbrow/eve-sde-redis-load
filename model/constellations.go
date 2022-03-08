package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cryanbrow/eve-sde-redis-load/helpers"
	"gopkg.in/yaml.v3"
)

type constellation struct {
	Center          []float64 `yaml:"center" json:"-"`
	ConstellationID int       `yaml:"constellationID" json:"constellation_id"`
	Systems         []int     `yaml:"-" json:"systems"`
	MaxArray        []float64 `yaml:"max" json:"-"`
	Max             position  `yaml:"-" json:"max"`
	MinArray        []float64 `yaml:"min" json:"-"`
	Min             position  `yaml:"-" json:"min"`
	Name            string    `yaml:"-" json:"name"`
	NameID          int       `yaml:"nameID" json:"name_id"`
	RegionID        int       `yaml:"-" json:"region_id"`
	Position        position  `yaml:"-" json:"position"`
}

func LoadConstellation(path string) {
	var file *os.File
	var err error
	var sdeConstellation constellation
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeConstellation)
	if err != nil {
		fmt.Println(err.Error())
	}

	directoryTraveralPath := strings.Replace(path, "constellation.staticdata", "", -1)

	systemNames := helpers.ReturnDirNames(directoryTraveralPath)

	for _, systemName := range systemNames {
		sdeConstellation.Systems = append(sdeConstellation.Systems, ids[systemName].ItemID)
	}

	sdeConstellation.Max.X = sdeConstellation.MaxArray[0]
	sdeConstellation.Max.Y = sdeConstellation.MaxArray[1]
	sdeConstellation.Max.Z = sdeConstellation.MaxArray[2]

	sdeConstellation.Min.X = sdeConstellation.MinArray[0]
	sdeConstellation.Min.Y = sdeConstellation.MinArray[1]
	sdeConstellation.Min.Z = sdeConstellation.MinArray[2]

	sdeConstellation.Position.X = sdeConstellation.Center[0]
	sdeConstellation.Position.Y = sdeConstellation.Center[1]
	sdeConstellation.Position.Z = sdeConstellation.Center[2]

	sdeConstellation.RegionID = ids[determineRegionNameFromFilePath(path)].ItemID

	sdeConstellation.Name = names[sdeConstellation.ConstellationID].ItemName

	singleConstellationJSON, _ := json.MarshalIndent(sdeConstellation, "", "  ")
	singleConstellationJSONString := string(singleConstellationJSON[:])
	fmt.Println(singleConstellationJSONString)

}
