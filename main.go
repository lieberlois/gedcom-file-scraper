package main

import (
	"flag"
	"fmt"
	"gedcomFiles/util"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// Read CLI flags
	targetPath := flag.String("targetFolder", "data", "Path to target folder")
	pathToGedcom := flag.String("gedcomPath", "FamilieNeu.ged", "Path to .gedcom-File")
	flag.Parse()

	if *pathToGedcom == "" {
		log.Fatal("Missing path to .gedcom-File")
	}

	// Get paths from .gedcom-File
	abspathToGedcomm, err := filepath.Abs(*pathToGedcom)
	if err != nil {
		log.Fatal("Error getting absolute path to .gedcom-File!")
	}
	paths, err := util.ExtractGedcomPaths(*pathToGedcom, filepath.Dir(abspathToGedcomm))


	if err != nil {
		log.Fatal("Error loading paths.")
	}

	fmt.Printf("Found %d paths in the .gedcom-File.\n", len(paths))

	// Create target folder if it doesnt exist.
	if _, err := os.Stat(*targetPath); os.IsNotExist(err) {
		err = os.Mkdir(*targetPath, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating target folder!")
		}
	}

	amount := 0
	for _, path := range paths {
		if fileInfo, err := os.Stat(path); err == nil {
			err := util.CopyFile(path, fmt.Sprintf("%s/%s", *targetPath, fileInfo.Name()))
			if err != nil {
				fmt.Printf("Copy of file '%s' failed", path)
				continue
			}
			amount++
		}
	}
	fmt.Printf("Successfully copied %d of %d images.", amount, len(paths))
}