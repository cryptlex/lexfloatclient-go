package main

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	baseURL := "https://dl.cryptlex.com/downloads/"
      libVersion :=  "v4.9.0";
	basePath := "./libs/"
	fmt.Println("Downloading LexFloatClient libs " + libVersion + " ...")
	url := baseURL + libVersion + "/LexFloatClient-Static-Mac.zip"
	err := downloadFile(url, "libs/clang/universal/libLexFloatClient.a", basePath+"darwin_universal/libLexFloatClient.a")
	if err != nil {
		panic(err)
	}
	url = baseURL + libVersion + "/LexFloatClient-Win.zip"
	err = downloadFile(url, "libs/vc14/x64/LexFloatClient.dll", basePath+"windows_amd64/LexFloatClient.dll")
	if err != nil {
		panic(err)
	}
	url = baseURL + libVersion + "/LexFloatClient-Static-Linux.zip"
	err = downloadFile(url, "libs/gcc/amd64/libLexFloatClient.a", basePath+"linux_amd64/libLexFloatClient.a")
	if err != nil {
		panic(err)
	}
	err = downloadFile(url, "libs/gcc/arm64/libLexFloatClient.a", basePath+"linux_arm64/libLexFloatClient.a")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println("LexFloatClient libs downloaded successfully!")
}

func downloadFile(url string, packagePath string, targetpath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = unzip(body, packagePath, targetpath)

	if err != nil {
		return err
	}
	return nil
}

func unzip(body []byte, packagePath string, targetpath string) error {
	reader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return err
	}

	for _, file := range reader.File {
		if file.Name == packagePath {
			fileReader, err := file.Open()
			if err != nil {
				return err
			}
			defer fileReader.Close()
			targetFile, err := os.Create(targetpath) // OpenFile(targetpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC,)
			if err != nil {
				return err
			}
			defer targetFile.Close()

			if _, err := io.Copy(targetFile, fileReader); err != nil {
				return err
			}
		}
	}

	return nil
}
