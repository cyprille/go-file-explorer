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
	"strings"
)

// Page struct to define the template content
type Page struct {
	Title            string
	Items            []string
	WorkingDirectory string
	Path             string
	PreviousEnabled  bool
}

// Current path
var path = "./"

// Defines if the user can go previous or not
var previousEnabled = false

// GoHome handles the response from the home path call
func GoHome(rw http.ResponseWriter, req *http.Request) {
	// Calls the navigation
	navigate(rw, "./")
}

// GoToPath handles the response from a path call
func GoToPath(rw http.ResponseWriter, req *http.Request) {
	// Retrieves the link by removing the "/api/navigation/" prefix
	path = strings.TrimPrefix(req.RequestURI, "/api/navigation/")

	// Calls the navigation
	navigate(rw, path)
}

// navigate displays the content of the given path parameter
func navigate(rw http.ResponseWriter, path string) {
	// Handles the possibility to go previous or not depending on current path
	if path == "./" || path == "." {
		previousEnabled = false
	} else {
		previousEnabled = true
	}

	// Retrieves the content list
	items, err := filesystem.GetPathContent(path)

	// If the path is not found
	if err != nil {
		// Renders the 404 error template
		common.Templates = template.Must(template.ParseFiles("templates/filesystem/404.html", common.LayoutPath))

		// Handles the errors
		err := common.Templates.ExecuteTemplate(rw, "base", nil)
		common.CheckError(err, 2)

		return
	}

	// Defines the page parameters
	p := Page{
		Title:           "Home",
		Items:           items,
		Path:            path,
		PreviousEnabled: previousEnabled,
	}

	// Renders the template
	common.Templates = template.Must(template.ParseFiles("templates/filesystem/list.html", common.LayoutPath))

	// Handles the errors
	err = common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
