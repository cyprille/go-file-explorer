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
	"os"

	"go-file-explorer/app/common"
)

// Init vars
var symlinks map[string][]string
var directories map[string][]string
var files map[string][]string

// GetPathContent Returns the list of files and directories in the given RootDir/path
func GetPathContent(path string, showHiddenFiles bool) (map[string]map[string][]string, error) {
	// Declares data types
	symlinks = map[string][]string{}
	directories = map[string][]string{}
	files = map[string][]string{}

	rd := common.GetParam("ROOT_DIR")

	// Reads the content of the given path
	content, err := ioutil.ReadDir(rd + path)
	if err != nil {
		return nil, err
	}

	// Loop over content (directories files)
	for _, c := range content {
		// Jumps the iteration if we won't show hidden files and if the file is an hidden one
		if showHiddenFiles == false && c.Name()[0:1] == "." {
			continue
		}

		sortContent(rd, path, c)
	}

	// Init and populates the values map
	items := map[string]map[string][]string{
		"1_directories": directories,
		"2_symlinks":    symlinks,
		"3_files":       files,
	}

	return items, nil
}

// sortContent Sorts content depending on items types
func sortContent(rd string, path string, c os.FileInfo) (bool, error) {
	// Retrieves informations on the target path
	file, err := os.Stat(rd + path + c.Name())
	if err != nil {
		return false, err
	}

	// Retrieves the file information
	fileInfo, err := os.Lstat(rd + path + c.Name())

	// Checks if the parsed item is a symlink
	if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		storeContent("symlinks", c.Name())
	} else {
		// Checks if the parsed item is a directory or a file
		switch mode := file.Mode(); {
		case mode.IsDir():
			storeContent("directories", c.Name())
		case mode.IsRegular():
			storeContent("files", c.Name())
		}
	}

	return true, nil
}

// storeContent Stores the content value depending on the item mode
func storeContent(m string, v string) {
	// Array container
	var ac map[string][]string

	// Handles the mode of the item
	switch m {
	case "symlinks":
		ac = symlinks
	case "directories":
		ac = directories
	case "files":
		ac = files
	}

	// Dispatches the item depending on the item's visibility
	if v[0:1] == "." {
		ac["hidden"] = append(ac["hidden"], v)
	} else {
		ac["regular"] = append(ac["regular"], v)
	}
}

// RetrieveFilePath Returns the full file path from the RootDir and the given path
func RetrieveFilePath(path string) string {
	return common.GetParam("ROOT_DIR") + path
}
