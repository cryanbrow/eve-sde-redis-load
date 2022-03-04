package model

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var sdeStaStations []staStation
var staStations = make(map[int]staStation)

type staStation struct {
	ConstellationID          int     `yaml:"constellationID" json:"constellation_id"`
	CorporationID            int     `yaml:"corporationID" json:"corporation_id"`
	DockingCostPerVolume     int     `yaml:"dockingCostPerVolume" json:"docking_cost_per_volume"`
	MaxShipVolumeDockable    int     `yaml:"maxShipVolumeDockable" json:"max_ship_volume_dockable"`
	OfficeRentalCost         int     `yaml:"officeRentalCost" json:"office_rental_cost"`
	OperationID              int     `yaml:"operationID" json:"operation_id"`
	RegionID                 int     `yaml:"regionID" json:"region_id"`
	ReprocessingEfficiency   float32 `yaml:"reprocessingEfficiency" json:"reprocessing_efficiency"`
	ReprocessingHangarFlag   int     `yaml:"reprocessingHangarFlag" json:"reprocessing_hangar_flag"`
	ReprocessingStationsTake float32 `yaml:"reprocessingStationsTake" json:"reprocessing_stations_take"`
	Security                 float32 `yaml:"security" json:"security"`
	SolarSystemID            int     `yaml:"solarSystemID" json:"solar_system_id"`
	StationID                int     `yaml:"stationID" json:"station_id"`
	StationName              string  `yaml:"stationName" json:"station_name"`
	StationTypeID            int     `yaml:"stationTypeID" json:"station_type_id"`
	X                        float64 `yaml:"x" json:"x"`
	Y                        float64 `yaml:"y" json:"y"`
	Z                        float64 `yaml:"z" json:"z"`
}

func LoadStaStations(path string) {
	var file *os.File
	var err error
	file, _ = os.Open(filepath.Clean(path))
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&sdeStaStations)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, element := range sdeStaStations {
		staStations[element.StationID] = element
		if element.StationID == 60015163 {
			singleStaStationJSON, _ := json.MarshalIndent(element, "", "  ")
			singleStaStationJSONString := string(singleStaStationJSON[:])

			fmt.Println(singleStaStationJSONString)
		}

		/*nameJSON, _ := json.Marshal(localname)
		redisKey := "name:" + strconv.Itoa(element.ItemID)
		status := data.Rdb.Set(context.Background(), redisKey, nameJSON, 0)
		statusText, _ := status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(nameJSON))

		idJSON, _ := json.Marshal(localID)
		redisKey = "id:" + strconv.Itoa(element.ItemID)
		status = data.Rdb.Set(context.Background(), redisKey, idJSON, 0)
		statusText, _ = status.Result()
		fmt.Printf("status text: %s \n", statusText)
		fmt.Println(string(idJSON))
		*/
	}

}
