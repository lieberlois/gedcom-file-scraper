This is an open-source project which can be used to gather referenced data like
images or other files in .GEDCOM-files. Since these files only store references to 
absolute/relative paths, the referenced files can be scattered across the filesystem.

This tool attempts to read the paths from a .GEDCOM-file and attempts to find all referenced files.
The referenced files are then copied to a new folder which then contains all (found) files.

Build: go build -o output_path

Usage: [output_path] -gedcomPath path -targetFolder name

-gedcomPath string
        Path to .gedcom-file (default "FamilieNeu.ged")
-targetFolder string
        Path to target folder (default "data")