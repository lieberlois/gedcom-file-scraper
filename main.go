package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"flag"
)

func main() {
	pathFile := flag.String("pathFile", "paths.txt", "Path to .txt-File containing filepaths")
	targetPath := flag.String("targetPath", "data", "Path to target folder")
	flag.Parse()


	paths, err := getPaths(*pathFile)

	if err != nil {
		log.Fatal("Error loading paths.")
	}

	if _, err := os.Stat(*targetPath); os.IsNotExist(err) {
		err = os.Mkdir(*targetPath, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating target folder")
		}
	}

	amount := 0
	for _, path := range paths {
		if fileInfo, err := os.Stat(path); err == nil {
			err := copyFile(path, fmt.Sprintf("%s/%s", *targetPath, fileInfo.Name()))
			if err != nil {
				fmt.Printf("Copy of file '%s' failed", path)
				continue
			}
			amount++
		}
	}
	fmt.Printf("Successfully copied %d of %d images.", amount, len(paths))
}

func getPaths(path string) ([]string, error) {
	result := make([]string, 0)
	file, err := os.Open(path)

	if err != nil {
		return result, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)


	for sc.Scan() {
		result = append(result, sc.Text())
	}
	return result, nil
}

func copyFile(sourceFile string, destinationFile string) error {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {
		return err
	}
	return nil
}