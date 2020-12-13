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
	pathToGedcom := flag.String("gedcomPath", "FamilieNeu.ged", "Path to .gedcom-file")
	flag.Parse()

	if *pathToGedcom == "" {
		log.Fatal("Missing path to .gedcom-file")
	}

	// Extract paths from .gedcom-File
	absoluteGedcomPath, err := filepath.Abs(*pathToGedcom)
	if err != nil {
		log.Fatal("Error getting absolute path to .gedcom-file!")
	}
	paths, err := util.ExtractGedcomPaths(*pathToGedcom, filepath.Dir(absoluteGedcomPath))

	if err != nil {
		log.Fatal("Error loading paths.")
	}

	log.Printf("Found %d paths in the .gedcom-file.\n", len(paths))

	// Create target folder if it doesnt exist.
	if _, err := os.Stat(*targetPath); os.IsNotExist(err) {
		err = os.Mkdir(*targetPath, os.ModePerm)
		if err != nil {
			log.Fatal("Error creating target folder!")
		}
	}

	// Attempt to copy files
	amount := 0
	for _, path := range paths {
		if fileInfo, err := os.Stat(path); err == nil {
			err := util.CopyFile(path, fmt.Sprintf("%s/%s", *targetPath, fileInfo.Name()))
			if err != nil {
				log.Printf("Copy of file '%s' failed", path)
				continue
			}
			amount++
		}
	}
	log.Printf("Successfully copied %d of %d files.", amount, len(paths))
}