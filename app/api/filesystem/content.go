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
	"log"
)

// GetPathContent Returns the list of files and directories in the given rootDir/path
func GetPathContent(path string) []string {
	// Reads the content of the given path
	files, error := ioutil.ReadDir(path)
	if error != nil {
		log.Fatal(error)
	}

	// Stores the data in an array
	var items = []string{}
	for _, f := range files {
		items = append(items, f.Name())
	}

	return items
}
