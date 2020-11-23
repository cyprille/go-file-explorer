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
	"fmt"
	"go-file-explorer/app/api/filesystem"
	"go-file-explorer/app/common"
	"html/template"
	"net/http"
)

// Page struct to define templates content
type Page struct {
	Title            string
	Items            []string
	CurrentDirectory string
}

// Base directory chroot
const rootDir = "/Users/cyprillechauvry/workspace"

// Current work directory
var workDir = rootDir

// GetHome Handles the response from the home path call
func GetHome(rw http.ResponseWriter, req *http.Request) {
	// Retrieves the "child" URI parameter
	var child = req.URL.Query()["child"]

	// Retrieves the path if we passed a child param to navigate to
	if len(child) > 0 {
		workDir = workDir + "/" + child[0]
	}

	fmt.Print("\nDump: ", workDir, "\n")

	// Retrieves the content list
	var items = filesystem.GetPathContent(workDir)

	// Defines the page
	p := Page{
		Title:            "Home",
		Items:            items,
		CurrentDirectory: workDir,
	}

	// Renders the template
	common.Templates = template.Must(template.ParseFiles("templates/filesystem/home.html", common.LayoutPath))

	// Handles errors
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
