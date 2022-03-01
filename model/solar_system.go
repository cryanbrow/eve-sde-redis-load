package model

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var sdeSolarSystem solarSystem

type solarSystem map[int]struct {
	Border                     bool             `yaml:"border" json:"border"`
	Center                     []float64        `yaml:"center"`
	ConstellationID            int              `json:"constellationID"`
	Corridor                   bool             `yaml:"corridor" json:"corridor"`
	DisallowedAnchorCategories []int            `yaml:"disallowedAnchorCategories" json:"disallowedAnchorCategories"`
	DisallowedAnchorGroups     []int            `yaml:"disallowedAnchorGroups" json:"disallowedAnchorGroups"`
	Fringe                     bool             `yaml:"fringe" json:"fringe"`
	Hub                        bool             `yaml:"hub" json:"hub"`
	International              bool             `yaml:"international" json:"international"`
	Luminosity                 float64          `yaml:"luminosity" json:"luminosity"`
	MaxArray                   []float64        `yaml:"max"`
	Max                        position         `json:"max"`
	MinArray                   []float64        `yaml:"min"`
	Min                        position         `json:"min"`
	Planets                    map[int]planet   `yaml:"planets"`
	Position                   position         `json:"position"`
	Radius                     float64          `yaml:"radius" json:"radius"`
	Regional                   bool             `yaml:"regional" json:"regional"`
	RegionID                   int              `json:"regionID"`
	Security                   float64          `yaml:"security" json:"security"`
	SecurityClass              string           `yaml:"securityClass" json:"securityClass"`
	SolarSystemID              int              `yaml:"solarSystemID" json:"solarSystemID"`
	SolarSystemNameID          int              `yaml:"solarSystemNameID" json:"solarSystemNameID"`
	Star                       starType         `yaml:"star"`
	StarID                     int              `json:"star_id"`
	Stargates                  map[int]stargate `yaml:"stargates"`
	StargateIDs                []int            `json:"stargates"`
	SystemPlanets              []systemPlanet   `json:"planets"`
	SunTypeID                  int              `yaml:"sunTypeID" json:"sunTypeID"`
}

type asteroidBelt struct {
	Name          string         `json:"name"`
	PositionArray []float64      `yaml:"position"`
	Position      position       `json:"position"`
	Statistics    statisticsType `yaml:"statistics" json:"statistics"`
	SystemID      int            `json:"system_id"`
	TypeID        int            `yaml:"typeID" json:"typeID"`
}

type statisticsType struct {
	Density        float32 `yaml:"density" json:"density"`
	Eccentricity   float64 `yaml:"eccentricity" json:"eccentricity"`
	EscapeVelocity float64 `yaml:"escapeVelocity" json:"escapeVelocity"`
	Fragmented     bool    `yaml:"fragmented" json:"fragmented"`
	Life           float64 `yaml:"life" json:"life"`
	Locked         bool    `yaml:"locked" json:"locked"`
	MassDust       float64 `yaml:"massDust" json:"massDust"`
	MassGas        float64 `yaml:"massGas" json:"massGas"`
	OrbitPeriod    float64 `yaml:"orbitPeriod" json:"orbitPeriod"`
	OrbitRadius    float64 `yaml:"orbitRadius" json:"orbitRadius"`
	Pressure       float64 `yaml:"pressure" json:"pressure"`
	Radius         float64 `yaml:"radius" json:"radius"`
	RotationRate   float64 `yaml:"rotationRate" json:"rotationRate"`
	SpectralClass  string  `yaml:"spectralClass" json:"spectralClass"`
	SurfaceGravity float64 `yaml:"surfaceGravity" json:"surfaceGravity"`
	Temperature    float64 `yaml:"temperature" json:"temperature"`
}

type starStatisticsType struct {
	Age           float64 `yaml:"age" json:"age"`
	Life          float64 `yaml:"life" json:"life"`
	Locked        bool    `yaml:"locked" json:"locked"`
	Luminosity    float64 `yaml:"luminosity" json:"luminosity"`
	Radius        float64 `yaml:"radius" json:"radius"`
	SpectralClass string  `yaml:"spectralClass" json:"spectralClass"`
	Temperature   float64 `yaml:"temperature" json:"temperature"`
}

type planet struct {
	AsteroidBelts    map[int]asteroidBelt `yaml:"asteroidBelts" json:"asteroidBelts"`
	CelestialIndex   int                  `yaml:"celestialIndex" json:"celestialIndex"`
	Moons            map[int]moon         `yaml:"moons" json:"moons"`
	PlanetAttributes planetAttributesType `yaml:"planetAttributes" json:"planetAttributes"`
	PositionArray    []float64            `yaml:"position"`
	Position         position             `json:"position"`
	Radius           int                  `yaml:"radius" json:"radius"`
	Statistics       statisticsType       `yaml:"statistics" json:"statistics"`
	TypeID           int                  `yaml:"typeID" json:"typeID"`
}

type planetAttributesType struct {
	HeightMap1   int  `yaml:"heightMap1" json:"heightMap1"`
	HeightMap2   int  `yaml:"heightMap2" json:"heightMap2"`
	Population   bool `yaml:"population" json:"population"`
	ShaderPreset int  `yaml:"shaderPreset" json:"shaderPreset"`
}

type moon struct {
	NPCStations      map[int]npcStation
	PlanetAttributes planetAttributesType `yaml:"planetAttributes" json:"planetAttributes"`
	PositionArray    []float64            `yaml:"position"`
	Position         position             `json:"position"`
	Padius           int                  `yaml:"radius" json:"radius"`
	Statistics       statisticsType       `yaml:"statistics" json:"statistics"`
	TypeID           int                  `yaml:"typeID" json:"typeID"`
}

type npcStation struct {
	GraphicID                int       `yaml:"graphicID" json:"graphicID"`
	IsConquerable            bool      `yaml:"isConquerable" json:"isConquerable"`
	OperationID              int       `yaml:"operationID" json:"operationID"`
	OwnerID                  int       `yaml:"ownerID" json:"ownerID"`
	PositionArray            []float64 `yaml:"position"`
	Position                 position  `json:"position"`
	ReprocessingEfficiency   float32   `yaml:"reprocessingEfficiency" json:"reprocessingEfficiency"`
	ReprocessingHangarFlag   int       `yaml:"reprocessingHangarFlag" json:"reprocessingHangarFlag"`
	ReprocessingStationsTake float32   `yaml:"reprocessingStationsTake" json:"reprocessingStationsTake"`
	TypeID                   int       `yaml:"typeID" json:"typeID"`
	UseOperationName         bool      `yaml:"useOperationName" json:"useOperationName"`
}

type position struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type systemPlanet struct {
	AsteroidBelts []int `json:"asteroid_belts"`
	Moons         []int `json:"moons"`
	PlanetID      int   `json:"planet_id"`
}

type starType struct {
	Id         int                `yaml:"id" json:"id"`
	Radius     int                `yaml:"radius" json:"radius"`
	Statistics starStatisticsType `yaml:"statistics" json:"statistics"`
	TypeID     int                `yaml:"typeID" json:"typeID"`
}

type stargate struct {
	Destination   int       `yaml:"destination" json:"destination"`
	PositionArray []float64 `yaml:"position"`
	Position      position  `json:"position"`
	TypeID        int       `yaml:"typeID" json:"typeID"`
}

func LoadSolarSystem(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeSolarSystem)
	if err != nil {
		fmt.Println(err.Error())
	}

}
