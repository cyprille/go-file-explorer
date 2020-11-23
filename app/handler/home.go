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

// Base directory chroot constant
// @TODO: refacto to handle this in project parameters
const rootDir = "/Users/cyprillechauvry/workspace"

// Page struct to define the template content
type Page struct {
	Title            string
	Items            []string
	WorkingDirectory string
	PreviousEnabled  bool
}

// Current working directory
var workDir = rootDir

// Defines if the user can go previous or not
var previousEnabled = false

// GoHome handles the response from the home path call
func GoHome(rw http.ResponseWriter, req *http.Request) {
	workDir = rootDir

	// Calls the navigation
	Navigate(rw, workDir)
}

// GoBack handles the response from the back path call
func GoBack(rw http.ResponseWriter, req *http.Request) {
	// Appends the return signal to the current working directory
	workDir = workDir + "/../"

	// Cleans the path to interpret the return signal "../"
	workDir = filepath.Clean(workDir)

	// Calls the navigation
	Navigate(rw, workDir)
}

// GoNext handles the response from a path call
func GoNext(rw http.ResponseWriter, req *http.Request) {
	// Retrieves the link by removing the "/api/navigation/" prefix
	link := strings.TrimPrefix(req.RequestURI, "/api/navigation/")

	// Appends the link to the working directory
	workDir = workDir + "/" + link

	// Calls the navigation
	Navigate(rw, workDir)
}

// Navigate displays the content of the given workDir parameter
func Navigate(rw http.ResponseWriter, workDir string) {
	// Security protection to check if the path is a children of the rootDir, otherwise, throws an error
	if strings.HasPrefix(workDir, rootDir) == false {
		// Renders the 403 error template
		common.Templates = template.Must(template.ParseFiles("templates/filesystem/403.html", common.LayoutPath))

		// Handles the errors
		err := common.Templates.ExecuteTemplate(rw, "base", nil)
		common.CheckError(err, 2)

		return
	}

	// Handles the possibility to go previous or not
	if rootDir == workDir {
		previousEnabled = false
	} else {
		previousEnabled = true
	}

	// Retrieves the content list
	items := filesystem.GetPathContent(workDir)

	// Defines the page parameters
	p := Page{
		Title:            "Home",
		Items:            items,
		WorkingDirectory: workDir,
		PreviousEnabled:  previousEnabled,
	}

	// Renders the template
	common.Templates = template.Must(template.ParseFiles("templates/filesystem/list.html", common.LayoutPath))

	// Handles the errors
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
