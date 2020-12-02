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
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"os"
	"strings"

	"go-file-explorer/app/api/filesystem"
	"go-file-explorer/app/common"
)

// Page struct to define the template content
type Page struct {
	AppTitle      string
	Items         map[string]map[string][]string
	RootDir       string
	Path          string
	Breadcrumbs   []string
	Depth         int
	BackLinks     []string
	RealDepth     int
	ParentEnabled bool
	CurrentPage   string
	DarkMode      bool
}

// Current path
var path = "./"

// Defines if the user can go to the parent or not
var parentEnabled = false

// HomeHandler handles the response from the home path call
func HomeHandler(rw http.ResponseWriter, req *http.Request) {
	// Defines CurrentPage parameter
	common.CurrentPage = "home"

	// Calls the navigation
	navigate(rw, req, "./")
}

// PathHandler handles the response from a path call
func PathHandler(rw http.ResponseWriter, req *http.Request) {
	// Defines CurrentPage parameter
	common.CurrentPage = "home"

	// Retrieves the link by removing the "/api/navigation/" prefix
	path = strings.TrimPrefix(req.RequestURI, "/api/navigation/")

	// Calls the navigation
	navigate(rw, req, path)
}

// OpenFileHandler Opens the file from the rootDir and the given path
// @TODO: Refacto this to handle the content's display with a stream
func OpenFileHandler(rw http.ResponseWriter, req *http.Request) {
	// Defines CurrentPage parameter
	common.CurrentPage = "file"

	// Retrieves the link by removing the "/api/navigation/" prefix
	path = strings.TrimPrefix(req.RequestURI, "/api/open/")

	// Retrieves the full file path from the filestystem
	filePath := filesystem.RetrieveFilePath(path)

	f, _ := os.Open(filePath)
	scanner := bufio.NewScanner(f)

	// Loop over all lines in the file and print them.
	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)
	}
}

// navigate displays the content of the given path parameter
func navigate(rw http.ResponseWriter, req *http.Request, path string) {
	// Decode special cars in path
	decodedPath, err := url.QueryUnescape(path)

	// Handles the possibility to go to the parent or not depending on current decodedPath
	if decodedPath == "./" || decodedPath == "." {
		parentEnabled = false
	} else {
		parentEnabled = true
	}

	// Retrieves the content list
	items, err := filesystem.GetPathContent(decodedPath, GetCookie(req, "show-hidden-files"))

	// If the decodedPath is not found
	if err != nil {
		// Boostraps the template
		common.Templates = template.Must(template.ParseFiles("templates/filesystem/404.html", common.LayoutPath))

		// Renders the 404 error template
		err := common.Templates.ExecuteTemplate(rw, "base", nil)
		common.CheckError(err, 2)

		return
	}

	// Generates values for breadcrumbs display
	breadcrumbs := strings.Split(strings.TrimSuffix(decodedPath, "/"), "/")
	realDepth := len(breadcrumbs)

	// Calculate the highest index to represent the depth of the breadcrumbs
	// It's just because golang templates cannot do arythmetics and range starts at 0 :/
	depth := realDepth - 1
	var backLinks []string

	for i := range breadcrumbs {
		// Generates the number of back links depending on the depth
		nb := realDepth - (i + 1)

		// Append the back link in the same order as the breadcrumbs
		backLinks = append(backLinks, strings.Repeat("../", nb))
	}

	// Defines the page parameters
	p := Page{
		AppTitle:      common.GetParam("APP_TITLE"),
		Items:         items,
		RootDir:       common.GetParam("ROOT_DIR"),
		Path:          decodedPath,
		Breadcrumbs:   breadcrumbs,
		Depth:         depth,
		BackLinks:     backLinks,
		RealDepth:     realDepth,
		ParentEnabled: parentEnabled,
		CurrentPage:   common.CurrentPage,
		DarkMode:      GetCookie(req, "dark-mode"),
	}

	// Boostraps the template
	common.Templates = template.Must(template.ParseFiles("templates/filesystem/list.html", common.LayoutPath))

	// Renders the template
	err = common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
