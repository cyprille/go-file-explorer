/*
 * Project: Go File Explorer
 * File: content.go
 * ---
 * Created: 22/11/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

package filesystem

import (
	"io/ioutil"
	"os"
)

// rootDir is the base directory chroot constant
// @TODO: refacto to handle this in project parameters
// @TODO: Improve this to handle files openning and 404 otherwise
const rootDir = "/Users/cyprillechauvry/workspace/"

// GetPathContent Returns the list of files and directories in the given rootDir/path
func GetPathContent(path string) (map[string][]string, error) {
	// Reads the content of the given path
	content, err := ioutil.ReadDir(rootDir + path)
	if err != nil {
		return nil, err
	}

	// Init the arrays of data
	items := map[string][]string{}
	directories := []string{}
	files := []string{}

	// Loop over content (directories and files)
	for _, c := range content {
		// Retrieves informations on the target path
		file, err := os.Stat(rootDir + path + c.Name())
		if err != nil {
			return nil, err
		}

		// Checks if the target path is a directory or a file
		switch mode := file.Mode(); {
		case mode.IsDir():
			directories = append(directories, c.Name())
		case mode.IsRegular():
			files = append(files, c.Name())
		}

		// Populates the values map
		items["directories"] = directories
		items["files"] = files
	}

	return items, nil
}
