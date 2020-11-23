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
)

// rootDir is the base directory chroot constant
// @TODO: refacto to handle this in project parameters
const rootDir = "/Users/cyprillechauvry/workspace/"

// GetPathContent Returns the list of files and directories in the given rootDir/path
func GetPathContent(path string) ([]string, error) {
	// Reads the content of the given path
	files, err := ioutil.ReadDir(rootDir + path)
	if err != nil {
		return nil, err
	}

	// Stores the data in an array
	items := []string{}
	for _, f := range files {
		items = append(items, f.Name())
	}

	return items, nil
}
