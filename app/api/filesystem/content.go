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
	"net/url"
	"os"

	"github.com/joho/godotenv"
)

// RootDir is the base directory chroot constant
var RootDir string

// Initializes the parameters from .env file
// @TODO: put this in dedicated package
func initParams() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	RootDir = os.Getenv("ROOT_DIR")
}

// GetPathContent Returns the list of files and directories in the given RootDir/path
func GetPathContent(path string) (map[string][]string, error) {
	// Bootstraps the parameters initialization
	// @TODO: put this in dedicated package
	initParams()

	decodedPath, err := url.QueryUnescape(path)

	// Reads the content of the given path
	content, err := ioutil.ReadDir(RootDir + decodedPath)
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
		file, err := os.Stat(RootDir + decodedPath + c.Name())
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

// RetrieveFilePath Returns the full file path from the RootDir and the given path
func RetrieveFilePath(path string) string {
	// Bootstraps the parameters initialization
	// @TODO: put this in dedicated package
	initParams()

	return RootDir + path
}
