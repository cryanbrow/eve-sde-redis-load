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
		group_ids goes before invUniqueNames
		invUniqueNames goes before any solar system/region/constellation
	*/

	data.ConfigureCaching()
	// https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/sde.zip
	// https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/checksum

	//DownloadFile("sde.zip", "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/sde.zip")
	//DownloadFile("checksum", "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/checksum")

	UnzipFile()

	//model.LoadUniqueNames("sde" + string(os.PathSeparator) + "sde" + string(os.PathSeparator) + "bsd" + string(os.PathSeparator) + "invUniqueNames.yaml")

	model.LoadSolarSystem("sde" + string(os.PathSeparator) + "sde" + string(os.PathSeparator) + "fsd" + string(os.PathSeparator) + "universe" + string(os.PathSeparator) + "eve" + string(os.PathSeparator) + "Metropolis" + string(os.PathSeparator) + "Aptetter" + string(os.PathSeparator) + "Erstur" + string(os.PathSeparator) + "solarsystem.staticdata")

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
	if strings.HasPrefix(fileName, "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"bsd"+string(os.PathSeparator)+"invUniqueNames.yaml") {
		model.LoadUniqueNames("sde" + string(os.PathSeparator) + "sde" + string(os.PathSeparator) + "bsd" + string(os.PathSeparator) + "invUniqueNames.yaml")
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Agents+model.Yaml {
		fmt.Println("Agents")
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.AgentsInSpace+model.Yaml {
		fmt.Println("AgentsInSpace")
	} else if fileName == "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"fsd"+string(os.PathSeparator)+model.Ancestries+model.Yaml {
		fmt.Println("Ancestries")
	} else {
		fmt.Println("fsd file")
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
