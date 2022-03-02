package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type solarSystem struct {
	Border                     bool             `yaml:"border" json:"border"`
	Center                     []float64        `yaml:"center" json:"-"`
	ConstellationID            int              `yaml:"-" json:"constellation_id"`
	Corridor                   bool             `yaml:"corridor" json:"corridor"`
	DisallowedAnchorCategories []int            `yaml:"disallowedAnchorCategories" json:"disallowed_anchor_categories"`
	DisallowedAnchorGroups     []int            `yaml:"disallowedAnchorGroups" json:"disallowed_anchor_groups"`
	Fringe                     bool             `yaml:"fringe" json:"fringe"`
	Hub                        bool             `yaml:"hub" json:"hub"`
	International              bool             `yaml:"international" json:"international"`
	Luminosity                 float64          `yaml:"luminosity" json:"luminosity"`
	MaxArray                   []float64        `yaml:"max" json:"-"`
	Max                        position         `yaml:"-" json:"max"`
	MinArray                   []float64        `yaml:"min"`
	Min                        position         `yaml:"-" json:"min"`
	Planets                    map[int]planet   `yaml:"planets"`
	Position                   position         `yaml:"-" json:"position"`
	Radius                     float64          `yaml:"radius" json:"radius"`
	Regional                   bool             `yaml:"regional" json:"regional"`
	RegionID                   int              `yaml:"-" json:"region_id"`
	Security                   float64          `yaml:"security" json:"security"`
	SecurityClass              string           `yaml:"securityClass" json:"security_class"`
	SolarSystemID              int              `yaml:"solarSystemID" json:"solar_system_id"`
	SolarSystemNameID          int              `yaml:"solarSystemNameID" json:"solar_system_name_id"`
	Star                       starType         `yaml:"star" json:"-"`
	StarID                     int              `yaml:"-" json:"star_id"`
	Stargates                  map[int]stargate `yaml:"stargates" json:"-"`
	StargateIDs                []int            `yaml:"-" json:"stargates"`
	SystemPlanets              []systemPlanet   `yaml:"-" json:"planets"`
	SunTypeID                  int              `yaml:"sunTypeID" json:"sun_type_id"`
}

type asteroidBelt struct {
	Name          string         `yaml:"-" json:"name"`
	PositionArray []float64      `yaml:"position" json:"-"`
	Position      position       `yaml:"-" json:"position"`
	Statistics    statisticsType `yaml:"statistics" json:"statistics"`
	SystemID      int            `yaml:"-" json:"system_id"`
	TypeID        int            `yaml:"typeID" json:"type_id"`
}

type statisticsType struct {
	Density        float32 `yaml:"density" json:"density"`
	Eccentricity   float64 `yaml:"eccentricity" json:"eccentricity"`
	EscapeVelocity float64 `yaml:"escapeVelocity" json:"escape_velocity"`
	Fragmented     bool    `yaml:"fragmented" json:"fragmented"`
	Life           float64 `yaml:"life" json:"life"`
	Locked         bool    `yaml:"locked" json:"locked"`
	MassDust       float64 `yaml:"massDust" json:"mass_dust"`
	MassGas        float64 `yaml:"massGas" json:"mass_gas"`
	OrbitPeriod    float64 `yaml:"orbitPeriod" json:"orbit_period"`
	OrbitRadius    float64 `yaml:"orbitRadius" json:"orbit_radius"`
	Pressure       float64 `yaml:"pressure" json:"pressure"`
	Radius         float64 `yaml:"radius" json:"radius"`
	RotationRate   float64 `yaml:"rotationRate" json:"rotation_rate"`
	SpectralClass  string  `yaml:"spectralClass" json:"spectral_class"`
	SurfaceGravity float64 `yaml:"surfaceGravity" json:"surface_gravity"`
	Temperature    float64 `yaml:"temperature" json:"temperature"`
}

type starStatisticsType struct {
	Age           float64 `yaml:"age" json:"age"`
	Life          float64 `yaml:"life" json:"life"`
	Locked        bool    `yaml:"locked" json:"locked"`
	Luminosity    float64 `yaml:"luminosity" json:"luminosity"`
	Radius        float64 `yaml:"radius" json:"radius"`
	SpectralClass string  `yaml:"spectralClass" json:"spectral_class"`
	Temperature   float64 `yaml:"temperature" json:"temperature"`
}

type planet struct {
	AsteroidBelts    map[int]asteroidBelt `yaml:"asteroidBelts" json:"asteroid_belts"`
	CelestialIndex   int                  `yaml:"celestialIndex" json:"celestial_index"`
	Moons            map[int]moon         `yaml:"moons" json:"moons"`
	PlanetAttributes planetAttributesType `yaml:"planetAttributes" json:"planet_attributes"`
	PositionArray    []float64            `yaml:"position" json:"-"`
	Position         position             `yaml:"-" json:"position"`
	Radius           int                  `yaml:"radius" json:"radius"`
	Statistics       statisticsType       `yaml:"statistics" json:"statistics"`
	TypeID           int                  `yaml:"typeID" json:"type_id"`
}

type planetAttributesType struct {
	HeightMap1   int  `yaml:"heightMap1" json:"height_map1"`
	HeightMap2   int  `yaml:"heightMap2" json:"height_map2"`
	Population   bool `yaml:"population" json:"population"`
	ShaderPreset int  `yaml:"shaderPreset" json:"shader_preset"`
}

type moon struct {
	NPCStations      map[int]npcStation
	PlanetAttributes planetAttributesType `yaml:"planetAttributes" json:"planet_attributes"`
	PositionArray    []float64            `yaml:"position" json:"-"`
	Position         position             `yaml:"-" json:"position"`
	Padius           int                  `yaml:"radius" json:"radius"`
	Statistics       statisticsType       `yaml:"statistics" json:"statistics"`
	TypeID           int                  `yaml:"typeID" json:"type_id"`
}

type npcStation struct {
	GraphicID                int       `yaml:"graphicID" json:"graphic_id"`
	IsConquerable            bool      `yaml:"isConquerable" json:"is_conquerable"`
	OperationID              int       `yaml:"operationID" json:"operation_id"`
	OwnerID                  int       `yaml:"ownerID" json:"ownerID"`
	PositionArray            []float64 `yaml:"position" json:"-"`
	Position                 position  `yaml:"-" json:"position"`
	ReprocessingEfficiency   float32   `yaml:"reprocessingEfficiency" json:"reprocessing_efficiency"`
	ReprocessingHangarFlag   int       `yaml:"reprocessingHangarFlag" json:"reprocessing_hangar_flag"`
	ReprocessingStationsTake float32   `yaml:"reprocessingStationsTake" json:"reprocessing_stations_take"`
	TypeID                   int       `yaml:"typeID" json:"type_id"`
	UseOperationName         bool      `yaml:"useOperationName" json:"use_operation_name"`
}

type position struct {
	X float64 `yaml:"-" json:"x"`
	Y float64 `yaml:"-" json:"y"`
	Z float64 `yaml:"-" json:"z"`
}

type systemPlanet struct {
	AsteroidBelts []int `yaml:"-" json:"asteroid_belts"`
	Moons         []int `yaml:"-" json:"moons"`
	PlanetID      int   `yaml:"-" json:"planet_id"`
}

type starType struct {
	Id         int                `yaml:"id" json:"id"`
	Radius     int                `yaml:"radius" json:"radius"`
	Statistics starStatisticsType `yaml:"statistics" json:"statistics"`
	TypeID     int                `yaml:"typeID" json:"type_id"`
}

type stargate struct {
	Destination   int       `yaml:"destination" json:"destination"`
	PositionArray []float64 `yaml:"position"`
	Position      position  `yaml:"-" json:"position"`
	TypeID        int       `yaml:"typeID" json:"type_id"`
}

func LoadSolarSystem(path string) {
	var sdeSolarSystem solarSystem
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

	singleSolarSystemJSON, _ := json.Marshal(sdeSolarSystem)
	str1 := string(singleSolarSystemJSON[:])

	fmt.Println(str1)

}
