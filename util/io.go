package util

import (
	"bufio"
	"io/ioutil"
	"os"
)

// ReadLinesFromFile scans through a file and returns a slice of strings for each line
func ReadLinesFromFile(path string) ([]string, error) {
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

// CopyFile attempts to copy a file from path A to path B returning occuring errors.
func CopyFile(sourceFile string, destinationFile string) error {
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