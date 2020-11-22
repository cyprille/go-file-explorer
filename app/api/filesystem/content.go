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

// GetPathContent Returns the list of files and directories in the path given in currentDirectory
func GetPathContent(rootDir string, path string) ([]string, string) {
	var dir = rootDir + "/" + path

	files, error := ioutil.ReadDir(dir)
	if error != nil {
		log.Fatal(error)
	}

	var items = []string{}
	for _, f := range files {
		items = append(items, f.Name())
	}

	return items, dir
}
