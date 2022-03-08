package main

import (
	"archive/zip"
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

	data.ConfigureCaching()
	// https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/sde.zip
	// https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/checksum

	//DownloadFile("sde.zip", "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/sde.zip")
	//DownloadFile("checksum", "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/checksum")

	//model.LoadStaStations("sde" + string(os.PathSeparator) + "sde" + string(os.PathSeparator) + "bsd" + string(os.PathSeparator) + "staStations.yaml")
	//model.LoadInvNames("sde" + string(os.PathSeparator) + "sde" + string(os.PathSeparator) + "bsd" + string(os.PathSeparator) + "invNames.yaml")
	//helpers.ReturnDirNames("sde" + string(os.PathSeparator) + "sde" + string(os.PathSeparator) + "fsd" + string(os.PathSeparator) + "universe" + string(os.PathSeparator) + "eve" + string(os.PathSeparator) + "Metropolis" + string(os.PathSeparator))
	UnzipFile()
	//model.LoadRegion("sde" + string(os.PathSeparator) + "sde" + string(os.PathSeparator) + "fsd" + string(os.PathSeparator) + "universe" + string(os.PathSeparator) + "eve" + string(os.PathSeparator) + "Metropolis" + string(os.PathSeparator) + "region.staticdata")
	//model.LoadConstellation("sde" + string(os.PathSeparator) + "sde" + string(os.PathSeparator) + "fsd" + string(os.PathSeparator) + "universe" + string(os.PathSeparator) + "eve" + string(os.PathSeparator) + "Metropolis" + string(os.PathSeparator) + "Eugidi" + string(os.PathSeparator) + "constellation.staticdata")
	//model.LoadSolarSystem("sde" + string(os.PathSeparator) + "sde" + string(os.PathSeparator) + "fsd" + string(os.PathSeparator) + "universe" + string(os.PathSeparator) + "eve" + string(os.PathSeparator) + "Metropolis" + string(os.PathSeparator) + "Aptetter" + string(os.PathSeparator) + "Erstur" + string(os.PathSeparator) + "solarsystem.staticdata")

}

func UnzipFile() {
	dst := "sde"
	archive, err := zip.OpenReader("sde.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()

	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		fmt.Println("unzipping file ", filePath)

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
	if strings.HasPrefix(fileName, "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"bsd"+string(os.PathSeparator)+"invNames.yaml") {
		model.LoadInvNames(fileName)
	} else if strings.HasPrefix(fileName, "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"bsd"+string(os.PathSeparator)+"invUniqueNames.yaml") {
		model.LoadUniqueNames(fileName)
	} else if strings.HasPrefix(fileName, "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"bsd"+string(os.PathSeparator)+"staStations.yaml") {
		model.LoadStaStations(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Agents+model.Yaml {
		model.LoadRedisAgents(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.AgentsInSpace+model.Yaml {
		model.LoadRedisAgentsInSpace(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Ancestries+model.Yaml {
		model.LoadRedisAncestries(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Bloodlines+model.Yaml {
		model.LoadRedisBloodlines(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Blueprints+model.Yaml {
		model.LoadRedisBlueprints(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.CategoryIDs+model.Yaml {
		model.LoadRedisCategoryIDs(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Certificates+model.Yaml {
		model.LoadRedisCertificates(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.CharacterAttributes+model.Yaml {
		model.LoadRedisCharacterAttributes(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.ContrabandTypes+model.Yaml {
		model.LoadRediscontrabandTypes(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.ControlTowerResources+model.Yaml {
		model.LoadRedisControlTowerAttributes(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.CorporationActivities+model.Yaml {
		model.LoadRedisCorporationActivities(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.DogmaAttributeCategories+model.Yaml {
		model.LoadRedisDogmaAttributeCategories(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.DogmaAttributes+model.Yaml {
		model.LoadRedisDogmaAttributes(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.DogmaEffects+model.Yaml {
		model.LoadRedissdeDogmaEffects(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Factions+model.Yaml {
		model.LoadRedisFactions(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.GraphicIDs+model.Yaml {
		model.LoadRedisGraphicIDs(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.GroupIDs+model.Yaml {
		model.LoadRedisGroupIDs(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.IconIDs+model.Yaml {
		model.LoadRedisIconIDS(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Landmarks+model.Yaml {
		model.LoadLandmarks(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.MarketGroups+model.Yaml {
		model.LoadMarketGroups(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.MetaGroups+model.Yaml {
		model.LoadMetaGroups(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.NpcCorporationDivisions+model.Yaml {
		model.LoadNPCCorporationDivisions(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.NpcCorporations+model.Yaml {
		model.LoadNPCCorporations(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.PlanetSchematics+model.Yaml {
		model.LoadPlanetSchematics(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Races+model.Yaml {
		model.LoadRaces(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.ResearchAgents+model.Yaml {
		model.LoadResearchAgents(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.SkinLicenses+model.Yaml {
		model.LoadSkinLicenses(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.SkinMaterials+model.Yaml {
		model.LoadSkinMaterials(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Skins+model.Yaml {
		model.LoadSkins(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.StationOperations+model.Yaml {
		model.LoadStationOperations(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.StationServices+model.Yaml {
		model.LoadStationServices(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TournamentRuleSets+model.Yaml {
		//TODO not implemented
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TranslationLanguages+model.Yaml {
		//TODO not implemented
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TypeDogma+model.Yaml {
		model.LoadTypeDogmas(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TypeIDs+model.Yaml {
		model.LoadTypeIDs(fileName)
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.TypeMaterials+model.Yaml {
		model.LoadTypeMaterials(fileName)
	} else if strings.Contains(fileName, "constellation.staticdata") {
		fmt.Println(fileName)
	} else if strings.Contains(fileName, "solarsystem.staticdata") {
		model.LoadSolarSystem(fileName)
	} else if strings.Contains(fileName, "solarsystem.staticdata") {
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
