package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	// https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/sde.zip
	// https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/checksum

	//DownloadFile("sde.zip", "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/sde.zip")
	//DownloadFile("checksum", "https://eve-static-data-export.s3-eu-west-1.amazonaws.com/tranquility/checksum")

	UnzipFile()

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
		fmt.Println(DetermineModelType(dstFile.Name()))

		dstFile.Close()
		fileInArchive.Close()
	}
}

func DetermineModelType(fileName string) string {
	if strings.HasPrefix(fileName, "sde"+string(os.PathSeparator)+"sde"+string(os.PathSeparator)+"bsd") {
		return "bsd file"
	}
	switch fileName {
	case local_model.Agents:
		return *ids.Agents[0].ID, nil
	case local_model.Alliances:
		return *ids.Alliances[0].ID, nil
	case local_model.Characters:
		return *ids.Characters[0].ID, nil
	case local_model.Constellations:
		return *ids.Constellations[0].ID, nil
	case local_model.Corporations:
		return *ids.Corporations[0].ID, nil
	case local_model.Factions:
		return *ids.Factions[0].ID, nil
	case local_model.InventoryTypes:
		return *ids.InventoryTypes[0].ID, nil
	case local_model.Regions:
		return *ids.Regions[0].ID, nil
	case local_model.Systems:
		return *ids.Systems[0].ID, nil
	default:
		return 0, errors.New("all fields nil")
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
