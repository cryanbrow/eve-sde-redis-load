package main

import (
	"archive/zip"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/cryanbrow/eve-sde-redis-load/data"
	"github.com/cryanbrow/eve-sde-redis-load/model"
)

func main() {

	//Order concerns
	/*
		invUniqueNames goes before any solar system/region/constellation
	*/

	data.ConfigureInMemoryCaching()
	//DownloadFile("sde.zip", "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/sde.zip")
	//DownloadFile("checksum", "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/checksum")
	UnzipFile()
	println("Number of keys in cache: ", data.NonExpiringCache.ItemCount())
	value, found := data.NonExpiringCache.Get("staStation:60003760")
	println("Value found: ", found)
	if found {
		prettyPrintJSON(value)
	} else {
		fmt.Println("Value not found")
	}
}

func UnzipFile() {
	dst := "sde"
	archive, err := zip.OpenReader("sde.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	fmt.Println("File count: ", len(archive.File))
	counter := 0
	for _, f := range archive.File {
		if counter%1000 == 1 {
			fmt.Println("Current File count: ", counter)
		}
		counter++
		filePath := filepath.Join(dst, f.Name)
		//fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			return
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			os.MkdirAll(filePath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		_, err = os.ReadFile(dstFile.Name())
		check(err)
		DetermineModelType(dstFile.Name())

		dstFile.Close()
		fileInArchive.Close()
	}
}

func DetermineModelType(fileName string) {
	if strings.HasPrefix(fileName, "sde"+string(os.PathSeparator)+"bsd"+string(os.PathSeparator)+"invNames.yaml") {
		model.LoadInvNames(fileName)
	} else if strings.HasPrefix(fileName, "sde"+string(os.PathSeparator)+"bsd"+string(os.PathSeparator)+"invUniqueNames.yaml") {
		model.LoadUniqueNames(fileName)
	} else if strings.HasPrefix(fileName, "sde"+string(os.PathSeparator)+"bsd"+string(os.PathSeparator)+"staStations.yaml") {
		model.LoadStaStations(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Agents+model.Yaml {
		model.LoadRedisAgents(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.AgentsInSpace+model.Yaml {
		model.LoadRedisAgentsInSpace(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Ancestries+model.Yaml {
		model.LoadRedisAncestries(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Bloodlines+model.Yaml {
		model.LoadRedisBloodlines(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Blueprints+model.Yaml {
		model.LoadRedisBlueprints(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.CategoryIDs+model.Yaml {
		model.LoadRedisCategoryIDs(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Certificates+model.Yaml {
		model.LoadRedisCertificates(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.CharacterAttributes+model.Yaml {
		model.LoadRedisCharacterAttributes(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.ContrabandTypes+model.Yaml {
		model.LoadRediscontrabandTypes(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.ControlTowerResources+model.Yaml {
		model.LoadRedisControlTowerAttributes(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.CorporationActivities+model.Yaml {
		model.LoadRedisCorporationActivities(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.DogmaAttributeCategories+model.Yaml {
		model.LoadRedisDogmaAttributeCategories(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.DogmaAttributes+model.Yaml {
		model.LoadRedisDogmaAttributes(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.DogmaEffects+model.Yaml {
		model.LoadRedissdeDogmaEffects(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Factions+model.Yaml {
		model.LoadRedisFactions(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.GraphicIDs+model.Yaml {
		model.LoadRedisGraphicIDs(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.GroupIDs+model.Yaml {
		model.LoadRedisGroupIDs(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.IconIDs+model.Yaml {
		model.LoadRedisIconIDS(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Landmarks+model.Yaml {
		model.LoadLandmarks(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.MarketGroups+model.Yaml {
		model.LoadMarketGroups(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.MetaGroups+model.Yaml {
		model.LoadMetaGroups(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.NpcCorporationDivisions+model.Yaml {
		model.LoadNPCCorporationDivisions(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.NpcCorporations+model.Yaml {
		model.LoadNPCCorporations(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.PlanetSchematics+model.Yaml {
		model.LoadPlanetSchematics(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Races+model.Yaml {
		model.LoadRaces(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.ResearchAgents+model.Yaml {
		model.LoadResearchAgents(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.SkinLicenses+model.Yaml {
		model.LoadSkinLicenses(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.SkinMaterials+model.Yaml {
		model.LoadSkinMaterials(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Skins+model.Yaml {
		model.LoadSkins(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.StationOperations+model.Yaml {
		model.LoadStationOperations(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.StationServices+model.Yaml {
		model.LoadStationServices(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TournamentRuleSets+model.Yaml {
		//TODO not implemented
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TranslationLanguages+model.Yaml {
		//TODO not implemented
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TypeDogma+model.Yaml {
		model.LoadTypeDogmas(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TypeIDs+model.Yaml {
		model.LoadTypeIDs(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TypeMaterials+model.Yaml {
		model.LoadTypeMaterials(fileName)
	} else if strings.Contains(fileName, "constellation.staticdata") {
		//TODO not implemented
	} else if strings.Contains(fileName, "solarsystem.yaml") {
		model.LoadSolarSystem(fileName)
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func CheckChecksum(zipFilePath, checksumFilePath string) bool {
	// Read the checksum from the checksum file
	checksumFile, err := os.Open(checksumFilePath)
	if err != nil {
		fmt.Println("Error opening checksum file:", err)
		return false
	}
	defer checksumFile.Close()

	var expectedChecksum string
	_, err = fmt.Fscanf(checksumFile, "%s", &expectedChecksum)
	if err != nil {
		fmt.Println("Error reading checksum file:", err)
		return false
	}

	// Calculate the checksum of the zip file
	zipFile, err := os.Open(zipFilePath)
	if err != nil {
		fmt.Println("Error opening zip file:", err)
		return false
	}
	defer zipFile.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, zipFile); err != nil {
		fmt.Println("Error calculating checksum:", err)
		return false
	}

	calculatedChecksum := fmt.Sprintf("%x", hash.Sum(nil))

	// Compare the checksums
	return calculatedChecksum == expectedChecksum
}

func prettyPrintJSON(value interface{}) {
	jsonString, ok := value.(string)
	if !ok {
		fmt.Println("Value is not a valid JSON string")
		return
	}

	var prettyJSON map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &prettyJSON)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	prettyBytes, err := json.MarshalIndent(prettyJSON, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	fmt.Println(string(prettyBytes))
}
