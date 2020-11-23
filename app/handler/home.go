/*
 * Project: Go File Explorer
 * File: home.go
 * ---
 * Created: 22/11/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

package handler

import (
	"go-file-explorer/app/api/filesystem"
	"go-file-explorer/app/common"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

// Page struct to define templates content
type Page struct {
	Title            string
	Items            []string
	CurrentDirectory string
	PreviousEnabled  bool
}

// Base directory chroot
const rootDir = "/Users/cyprillechauvry/workspace"

// Current work directory
var workDir = rootDir

// Defines if the user can go previous or not
var previousEnabled = false

// GetHome Handles the response from the home path call
func GetHome(rw http.ResponseWriter, req *http.Request) {
	// Retrieves the "child" URI parameter
	var child = req.URL.Query()["child"]

	// Retrieves the path if we passed a child param to navigate to
	if len(child) > 0 {
		workDir = workDir + "/" + child[0]
	}

	// Cleans the path to interpret the return signal "../"
	workDir = filepath.Clean(workDir)

	// Security protection to check if the path is a children of the rootDir, otherwise, throws an error
	if strings.HasPrefix(workDir, rootDir) == false {
		// Renders the 403 error template
		common.Templates = template.Must(template.ParseFiles("templates/filesystem/403.html", common.LayoutPath))

		// Handles the errors
		err := common.Templates.ExecuteTemplate(rw, "base", nil)
		common.CheckError(err, 2)

		return
	}

	// Retrieves the content list
	var items = filesystem.GetPathContent(workDir)

	// Handles the possibility to go previous or not
	if rootDir == workDir {
		previousEnabled = false
	} else {
		previousEnabled = true
	}

	// Defines the page parameters
	p := Page{
		Title:            "Home",
		Items:            items,
		CurrentDirectory: workDir,
		PreviousEnabled:  previousEnabled,
	}

	// Renders the template
	common.Templates = template.Must(template.ParseFiles("templates/filesystem/home.html", common.LayoutPath))

	// Handles the errors
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
