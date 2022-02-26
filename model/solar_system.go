package model

type solarSystem map[int]struct {
	border                     bool
	center                     []int
	corridor                   bool
	disallowedAnchorCategories []int
	disallowedAnchorGroups     []int
	fringe                     bool
	hub                        bool
	international              bool
	luminosity                 float32
	max                        []float64
	min                        []float64
	planets                    map[int]planet
	radius                     float64
	regional                   bool
	security                   float64
	securityClass              string
	solarSystemID              int
	solarSystemNameID          int
	star                       starType
	stargates                  map[int]stargate
	sunTypeID                  int
}

type asteroidBelt struct {
	position   []float64
	statistics struct {
		density        float32
		eccentricity   string
		escapeVelocity float32
		fragmented     bool
		life           float32
		locked         bool
		massDust       string
		massGas        string
		orbitPeriod    float32
		orbitRadius    float32
		pressure       string
		radius         float32
		rotationRate   float32
		spectralClass  string
		surfaceGravity float32
		temperature    float32
	}
	typeID int
}

type planet struct {
	asteroidBelts
	celestialIndex
	moons
	planetAttributes
	position
	radius
	statistics
	typeID
}

type moon struct {
	planetAttributes
	position
	radius
	statistics
	typeID
}

type starType struct {
	id
	radius
	statistics
	typeID
}

type stargate struct {
	destination
	position
	typeID
}
