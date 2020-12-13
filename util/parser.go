package util

import (
	"log"
	"path/filepath"
	"strings"
)

func ExtractGedcomPaths(path string, basepath string) ([]string, error) {
	lines, err := ReadLinesFromFile(path)
	if err != nil {
		log.Fatalf("Error parsing Gedcom File: %s", err)
	}
	res := make([]string, 0)

	for idx, line := range lines {
		// Ignore the .gedcom-file itself
		if strings.Contains(line, ".ged"){
			continue
		}

		// Standard paths are prefixed by 2 FILE
		if strings.Contains(line, "FILE "){
			splitPath := strings.Split(line, "FILE")
			path := strings.TrimSpace(splitPath[len(splitPath) - 1])

			res = append(res, path)

			nextLine := lines[idx+1]
			// Alternative paths are prefixed by 2 ALTPATH
			if strings.Contains(nextLine, "ALTPATH "){
				_, filename := filepath.Split(path)

				splitPath = strings.Split(nextLine, "ALTPATH")
				altpath := strings.TrimSpace(splitPath[len(splitPath) - 1])

				res = append(res, filepath.Join(basepath, altpath, filename))
			}
		}
	}

	return res, nil
}
