package model

type solarSystem map[int]struct {
	Border                     bool             `yaml:"border" json:"border"`
	Center                     []int            `yaml:"center" json:"center"`
	ConstellationID            int              `json:"constellationID"`
	Corridor                   bool             `yaml:"corridor" json:"corridor"`
	DisallowedAnchorCategories []int            `yaml:"disallowedAnchorCategories" json:"disallowedAnchorCategories"`
	DisallowedAnchorGroups     []int            `yaml:"disallowedAnchorGroups" json:"disallowedAnchorGroups"`
	Fringe                     bool             `yaml:"fringe" json:"fringe"`
	Hub                        bool             `yaml:"hub" json:"hub"`
	International              bool             `yaml:"international" json:"international"`
	Luminosity                 float32          `yaml:"luminosity" json:"luminosity"`
	Max                        []float64        `yaml:"max" json:"max"`
	Min                        []float64        `yaml:"min" json:"min"`
	Planets                    map[int]planet   `yaml:"planets" json:"planets"`
	Radius                     float64          `yaml:"radius" json:"radius"`
	Regional                   bool             `yaml:"regional" json:"regional"`
	Security                   float64          `yaml:"security" json:"security"`
	SecurityClass              string           `yaml:"securityClass" json:"securityClass"`
	SolarSystemID              int              `yaml:"solarSystemID" json:"solarSystemID"`
	SolarSystemNameID          int              `yaml:"solarSystemNameID" json:"solarSystemNameID"`
	Star                       starType         `yaml:"star" json:"star"`
	Stargates                  map[int]stargate `yaml:"stargates" json:"stargates"`
	SunTypeID                  int              `yaml:"sunTypeID" json:"sunTypeID"`
}

type asteroidBelt struct {
	Position   []float64      `yaml:"position" json:"position"`
	Statistics statisticsType `yaml:"statistics" json:"statistics"`
	TypeID     int            `yaml:"typeID" json:"typeID"`
}

type statisticsType struct {
	Density        float32 `yaml:"density" json:"density"`
	Eccentricity   string  `yaml:"eccentricity" json:"eccentricity"`
	EscapeVelocity float32 `yaml:"escapeVelocity" json:"escapeVelocity"`
	Fragmented     bool    `yaml:"fragmented" json:"fragmented"`
	Life           float32 `yaml:"life" json:"life"`
	Locked         bool    `yaml:"locked" json:"locked"`
	MassDust       string  `yaml:"massDust" json:"massDust"`
	MassGas        string  `yaml:"massGas" json:"massGas"`
	OrbitPeriod    float32 `yaml:"orbitPeriod" json:"orbitPeriod"`
	OrbitRadius    float32 `yaml:"orbitRadius" json:"orbitRadius"`
	Pressure       string  `yaml:"pressure" json:"pressure"`
	Radius         float32 `yaml:"radius" json:"radius"`
	RotationRate   float32 `yaml:"rotationRate" json:"rotationRate"`
	SpectralClass  string  `yaml:"spectralClass" json:"spectralClass"`
	SurfaceGravity float32 `yaml:"surfaceGravity" json:"surfaceGravity"`
	Temperature    float32 `yaml:"temperature" json:"temperature"`
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
	Position         []float64            `yaml:"position" json:"position"`
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
	npcStations      map[int]npcStation
	PlanetAttributes planetAttributesType `yaml:"planetAttributes" json:"planetAttributes"`
	Position         []float64            `yaml:"position" json:"position"`
	Padius           int                  `yaml:"radius" json:"radius"`
	Statistics       statisticsType       `yaml:"statistics" json:"statistics"`
	TypeID           int                  `yaml:"typeID" json:"typeID"`
}

type npcStation struct {
	graphicID                int       `yaml:"graphicID" json:"graphicID"`
	isConquerable            bool      `yaml:"isConquerable" json:"isConquerable"`
	operationID              int       `yaml:"operationID" json:"operationID"`
	ownerID                  int       `yaml:"ownerID" json:"ownerID"`
	position                 []float64 `yaml:"position" json:"position"`
	reprocessingEfficiency   float32   `yaml:"reprocessingEfficiency" json:"reprocessingEfficiency"`
	reprocessingHangarFlag   int       `yaml:"reprocessingHangarFlag" json:"reprocessingHangarFlag"`
	reprocessingStationsTake float32   `yaml:"reprocessingStationsTake" json:"reprocessingStationsTake"`
	typeID                   int       `yaml:"typeID" json:"typeID"`
	useOperationName         bool      `yaml:"useOperationName" json:"useOperationName"`
}

type starType struct {
	Id         int                `yaml:"id" json:"id"`
	Radius     int                `yaml:"radius" json:"radius"`
	Statistics starStatisticsType `yaml:"statistics" json:"statistics"`
	TypeID     int                `yaml:"typeID" json:"typeID"`
}

type stargate struct {
	Destination int       `yaml:"destination" json:"destination"`
	Position    []float64 `yaml:"position" json:"position"`
	TypeID      int       `yaml:"typeID" json:"typeID"`
}
