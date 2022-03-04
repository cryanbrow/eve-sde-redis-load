package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	MinArray                   []float64        `yaml:"min" json:"-"`
	Min                        position         `yaml:"-" json:"min"`
	Name                       string           `yaml:"-" json:"name"`
	Planets                    map[int]planet   `yaml:"planets" json:"-"`
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
	Stations                   []int            `yaml:"-" json:"stations"`
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
	AsteroidBelts     map[int]asteroidBelt `yaml:"asteroidBelts" json:"-"`
	AsteroidBeltArray []int                `yaml:"-" json:"asteroid_belts"`
	CelestialIndex    int                  `yaml:"celestialIndex" json:"celestial_index"`
	Moons             map[int]moon         `yaml:"moons" json:"-"`
	MoonArray         []int                `yaml:"-" json:"moons"`
	Name              string               `yaml:"-" json:"name"`
	PlanetAttributes  planetAttributesType `yaml:"planetAttributes" json:"planet_attributes"`
	PlanetID          int                  `yaml:"-" json:"planet_id"`
	PositionArray     []float64            `yaml:"position" json:"-"`
	Position          position             `yaml:"-" json:"position"`
	Radius            int                  `yaml:"radius" json:"radius"`
	Statistics        statisticsType       `yaml:"statistics" json:"statistics"`
	SystemID          int                  `yaml:"-" json:"system_id"`
	TypeID            int                  `yaml:"typeID" json:"type_id"`
}

type planetAttributesType struct {
	HeightMap1   int  `yaml:"heightMap1" json:"height_map1"`
	HeightMap2   int  `yaml:"heightMap2" json:"height_map2"`
	Population   bool `yaml:"population" json:"population"`
	ShaderPreset int  `yaml:"shaderPreset" json:"shader_preset"`
}

type moon struct {
	ID               int                  `yaml:"-" json:"moon_id"`
	Name             string               `yaml:"-" json:"name"`
	NPCStations      map[int]npcStation   `yaml:"npcStations" json:"-"`
	PlanetAttributes planetAttributesType `yaml:"planetAttributes" json:"planet_attributes"`
	PositionArray    []float64            `yaml:"position" json:"-"`
	Position         position             `yaml:"-" json:"position"`
	Radius           int                  `yaml:"radius" json:"radius"`
	Stations         []int                `yaml:"-" json:"stations"`
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

	sdeSolarSystem.Max.X = sdeSolarSystem.MaxArray[0]
	sdeSolarSystem.Max.Y = sdeSolarSystem.MaxArray[1]
	sdeSolarSystem.Max.Z = sdeSolarSystem.MaxArray[2]

	sdeSolarSystem.Min.X = sdeSolarSystem.MinArray[0]
	sdeSolarSystem.Min.Y = sdeSolarSystem.MinArray[1]
	sdeSolarSystem.Min.Z = sdeSolarSystem.MinArray[2]

	sdeSolarSystem.Position.X = sdeSolarSystem.Center[0]
	sdeSolarSystem.Position.Y = sdeSolarSystem.Center[1]
	sdeSolarSystem.Position.Z = sdeSolarSystem.Center[2]

	sdeSolarSystem.StarID = sdeSolarSystem.Star.Id

	sdeSolarSystem.StargateIDs = getStargateMapKeys(sdeSolarSystem.Stargates)

	for planetKey, planet := range sdeSolarSystem.Planets {
		planet.AsteroidBeltArray = getAsteroidBeltMapKeys(planet.AsteroidBelts)
		planet.MoonArray = getMoonMapKeys(planet.Moons)
		planet.Name = names[planetKey].ItemName
		planet.Position.X = planet.PositionArray[0]
		planet.Position.Y = planet.PositionArray[1]
		planet.Position.Z = planet.PositionArray[2]
		planet.SystemID = sdeSolarSystem.SolarSystemID
		planet.PlanetID = planetKey

		singlePlanetJSON, _ := json.MarshalIndent(planet, "", "  ")
		singlePlanetJSONString := string(singlePlanetJSON[:])

		fmt.Println(singlePlanetJSONString)

		var localSystemPlanet systemPlanet
		localSystemPlanet.PlanetID = planetKey
		localSystemPlanet.Moons = getMoonMapKeys(planet.Moons)
		localSystemPlanet.AsteroidBelts = getAsteroidBeltMapKeys(planet.AsteroidBelts)
		sdeSolarSystem.SystemPlanets = append(sdeSolarSystem.SystemPlanets, localSystemPlanet)

		for asteroidBeltKey, asteroidBelt := range planet.AsteroidBelts {
			asteroidBelt.Name = names[asteroidBeltKey].ItemName
			asteroidBelt.Position.X = asteroidBelt.PositionArray[0]
			asteroidBelt.Position.Y = asteroidBelt.PositionArray[1]
			asteroidBelt.Position.Z = asteroidBelt.PositionArray[2]
			asteroidBelt.SystemID = sdeSolarSystem.SolarSystemID

			singleAsteroidBeltJSON, _ := json.MarshalIndent(asteroidBelt, "", "  ")
			singleAsteroidBeltJSONString := string(singleAsteroidBeltJSON[:])

			fmt.Println(singleAsteroidBeltJSONString)
		}

		for moonKey, moon := range planet.Moons {
			moon.ID = moonKey
			moon.Name = names[moonKey].ItemName
			moon.Position.X = moon.PositionArray[0]
			moon.Position.Y = moon.PositionArray[1]
			moon.Position.Z = moon.PositionArray[2]
			moon.Stations = getStationMapKeys(moon.NPCStations)

			singleMoonJSON, _ := json.MarshalIndent(moon, "", "  ")
			singleMoonJSONString := string(singleMoonJSON[:])

			fmt.Println(singleMoonJSONString)

			for stationKey, station := range moon.NPCStations {
				fmt.Println(station)

				sdeSolarSystem.Stations = append(sdeSolarSystem.Stations, stationKey)
			}
		}
	}

	sdeSolarSystem.Name = determineSystemNameFromFilePath(path)

	sdeSolarSystem.ConstellationID = ids[determineConstallationNameFromFilePath(path)].ItemID

	sdeSolarSystem.RegionID = ids[determineRegionNameFromFilePath(path)].ItemID

	//singleSolarSystemJSON, _ := json.MarshalIndent(sdeSolarSystem, "", "  ")
	//str1 := string(singleSolarSystemJSON[:])

	//fmt.Println(str1)

}

func getStargateMapKeys(m map[int]stargate) []int {
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func getMoonMapKeys(m map[int]moon) []int {
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func getAsteroidBeltMapKeys(m map[int]asteroidBelt) []int {
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func getStationMapKeys(m map[int]npcStation) []int {
	j := 0
	keys := make([]int, len(m))
	for k := range m {
		keys[j] = k
		j++
	}
	return keys
}

func determineSystemNameFromFilePath(path string) string {
	values := strings.Split(path, string(os.PathSeparator))
	if len(values) == 9 {
		return values[7]
	}
	return ""
}

func determineConstallationNameFromFilePath(path string) string {
	values := strings.Split(path, string(os.PathSeparator))
	if len(values) == 9 {
		return values[6]
	}
	return ""
}

func determineRegionNameFromFilePath(path string) string {
	values := strings.Split(path, string(os.PathSeparator))
	if len(values) == 9 {
		return values[5]
	}
	return ""
}
