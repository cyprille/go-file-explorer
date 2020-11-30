/*
 * Project: Go File Explorer
 * File: filesystem.go
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
	"os"

	"github.com/joho/godotenv"
)

// RootDir is the base directory chroot constant
var RootDir string

// Initializes the parameters from .env file
// @TODO: put this in dedicated package
func initParams() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	RootDir = os.Getenv("ROOT_DIR")
}

// GetPathContent Returns the list of files and directories in the given RootDir/path
func GetPathContent(path string, showHiddenFiles bool) (map[string]map[string][]string, error) {
	// Bootstraps the parameters initialization
	// @TODO: put this in dedicated package
	initParams()

	// Reads the content of the given path
	content, err := ioutil.ReadDir(RootDir + path)
	if err != nil {
		return nil, err
	}

	// Init the arrays of data
	items := map[string]map[string][]string{
		"1_directories": {},
		"2_symlinks":    {},
		"3_files":       {},
	}

	// Declares data types
	hiddenDirectories := []string{}
	directories := []string{}
	hiddenSymlinks := []string{}
	symlinks := []string{}
	hiddenFiles := []string{}
	files := []string{}

	// Loop over content (directories files)
	for _, c := range content {
		// Jumps the iteration if we won't show hidden files and if the file is an hidden one
		if showHiddenFiles == false && c.Name()[0:1] == "." {
			continue
		}

		// Retrieves informations on the target path
		file, err := os.Stat(RootDir + path + c.Name())
		if err != nil {
			return nil, err
		}

		fileInfo, err := os.Lstat(RootDir + path + c.Name())

		// If the parsed file is a symlink
		if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
			// Handles the hidden mode of the item
			if c.Name()[0:1] == "." {
				hiddenSymlinks = append(hiddenSymlinks, c.Name())
			} else {
				symlinks = append(symlinks, c.Name())
			}
		} else {
			// Checks if the target path is a directory or a file
			switch mode := file.Mode(); {
			case mode.IsDir():
				// Handles the hidden mode of the item
				if c.Name()[0:1] == "." {
					hiddenDirectories = append(hiddenDirectories, c.Name())
				} else {
					directories = append(directories, c.Name())
				}
			case mode.IsRegular():
				// Handles the hidden mode of the item
				if c.Name()[0:1] == "." {
					hiddenFiles = append(hiddenFiles, c.Name())
				} else {
					files = append(files, c.Name())
				}
			}
		}

		// Populates the values map
		items["1_directories"]["hidden"] = hiddenDirectories
		items["1_directories"]["regular"] = directories
		items["2_symlinks"]["hidden"] = hiddenSymlinks
		items["2_symlinks"]["regular"] = symlinks
		items["3_files"]["hidden"] = hiddenFiles
		items["3_files"]["regular"] = files
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
