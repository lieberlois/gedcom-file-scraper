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
		if strings.Contains(line, ".ged"){
			continue
		}

		if strings.Contains(line, "FILE "){
			splitPath := strings.Split(line, "FILE")
			path := strings.TrimSpace(splitPath[len(splitPath) - 1])

			res = append(res, path)
			splitPath = strings.Split(path, "\\")

			nextLine := lines[idx+1]
			if strings.Contains(nextLine, "ALTPATH "){
				filename := strings.TrimSpace(splitPath[len(splitPath) - 1])

				if filename[:2] == ".\\" {
					filename = filename[:2]
				}

				splitPath = strings.Split(nextLine, "ALTPATH")
				altpath := strings.TrimSpace(splitPath[len(splitPath) - 1])

				res = append(res, filepath.Join(basepath, altpath, filename))
			}
		}
	}

	return res, nil
}
