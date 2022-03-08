package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cryanbrow/eve-sde-redis-load/helpers"
	"gopkg.in/yaml.v3"
)

type region struct {
	Center          []float64 `yaml:"center" json:"-"`
	Constellations  []int     `yaml:"-" json:"constellations"`
	DescriptionID   int       `yaml:"descriptionID" json:"description_id"`
	MaxArray        []float64 `yaml:"max" json:"-"`
	Max             position  `yaml:"-" json:"max"`
	MinArray        []float64 `yaml:"min" json:"-"`
	Min             position  `yaml:"-" json:"min"`
	Name            string    `yaml:"-" json:"name"`
	NameID          int       `yaml:"nameID" json:"name_id"`
	Nebula          int       `yaml:"nebula" json:"nebula"`
	RegionID        int       `yaml:"regionID" json:"region_id"`
	Position        position  `yaml:"-" json:"position"`
	WormholeClassID int       `yaml:"wormholeClassID" json:"wormhole_class_id"`
}

func LoadRegion(path string) {
	var file *os.File
	var err error
	var sdeRegion region
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeRegion)
	if err != nil {
		fmt.Println(err.Error())
	}

	directoryTraveralPath := strings.Replace(path, "region.staticdata", "", -1)

	constallationNames := helpers.ReturnDirNames(directoryTraveralPath)

	for _, constellationName := range constallationNames {
		sdeRegion.Constellations = append(sdeRegion.Constellations, ids[constellationName].ItemID)
	}

	sdeRegion.Max.X = sdeRegion.MaxArray[0]
	sdeRegion.Max.Y = sdeRegion.MaxArray[1]
	sdeRegion.Max.Z = sdeRegion.MaxArray[2]

	sdeRegion.Min.X = sdeRegion.MinArray[0]
	sdeRegion.Min.Y = sdeRegion.MinArray[1]
	sdeRegion.Min.Z = sdeRegion.MinArray[2]

	sdeRegion.Position.X = sdeRegion.Center[0]
	sdeRegion.Position.Y = sdeRegion.Center[1]
	sdeRegion.Position.Z = sdeRegion.Center[2]

	sdeRegion.Name = names[sdeRegion.RegionID].ItemName

}
