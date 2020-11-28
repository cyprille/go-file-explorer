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
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"go-file-explorer/app/api/filesystem"
	"go-file-explorer/app/common"

	"github.com/joho/godotenv"
)

// Page struct to define the template content
type Page struct {
	AppTitle        string
	Items           map[string][]string
	RootDir         string
	Path            string
	PreviousEnabled bool
	CurrentPage     string
}

// Current path
var path = "./"

// Defines if the user can go previous or not
var previousEnabled = false

// appTitle The name of the app
var appTitle string

// Initializes the parameters from .env file
// @TODO: put this in dedicated package
func initParams() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	appTitle = os.Getenv("APP_TITLE")
}

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
	// Bootstraps the parameters initialization
	// @TODO: put this in dedicated package
	initParams()

	decodedPath, err := url.QueryUnescape(path)

	// Handles the possibility to go previous or not depending on current decodedPath
	if decodedPath == "./" || decodedPath == "." {
		previousEnabled = false
	} else {
		previousEnabled = true
	}

	// Retrieves the content list
	items, err := filesystem.GetPathContent(decodedPath, ShowHiddenFiles(req))

	// If the decodedPath is not found
	if err != nil {
		// Boostraps the template
		common.Templates = template.Must(template.ParseFiles("templates/filesystem/404.html", common.LayoutPath))

		// Renders the 404 error template
		err := common.Templates.ExecuteTemplate(rw, "base", nil)
		common.CheckError(err, 2)

		return
	}

	// Defines the page parameters
	p := Page{
		AppTitle:        appTitle,
		Items:           items,
		RootDir:         filesystem.RootDir,
		Path:            decodedPath,
		PreviousEnabled: previousEnabled,
		CurrentPage:     common.CurrentPage,
	}

	// Boostraps the template
	common.Templates = template.Must(template.ParseFiles("templates/filesystem/list.html", common.LayoutPath))

	// Renders the template
	err = common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
