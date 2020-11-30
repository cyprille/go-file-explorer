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
	"go-file-explorer/app/common"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

// SettingsHandler handles the response for the settings page
func SettingsHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// Defines CurrentPage parameter
		common.CurrentPage = "settings"

		v := map[string]interface{}{
			"ShowHiddenFiles": ShowHiddenFiles(req),
			"CurrentPage":     common.CurrentPage,
		}

		// Boostraps the template
		common.Templates = template.Must(template.ParseFiles("templates/settings.html", common.LayoutPath))

		// Renders the template
		err := common.Templates.ExecuteTemplate(rw, "base", v)
		common.CheckError(err, 2)

	case "POST":
		// Parses the form and checks errors
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(rw, "ParseForm() err: %v", err)
			return
		}

		// Retrieve the show-hidden-files parameter value
		showHiddenFilesValue := req.FormValue("show-hidden-files")

		var showHiddenFiles bool

		// Sets the value of the parameter
		if showHiddenFilesValue == "on" {
			showHiddenFiles = true
		} else {
			showHiddenFiles = false
		}

		// Stores a cookie for ShowHiddenFiles parameter
		expiration := time.Now().Add(5 * 365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "show-hidden-files", Value: strconv.FormatBool(showHiddenFiles), Expires: expiration}
		http.SetCookie(rw, &cookie)

		// Redirects to the parent page
		http.Redirect(rw, req, req.Header.Get("Referer"), 302)
	}
}

// ShowHiddenFiles Returns if the hidden files must be shown
func ShowHiddenFiles(req *http.Request) bool {
	// Retrieves the cookie for the parameter ShowHiddenFiles
	showHiddenFilesCookie, _ := req.Cookie("show-hidden-files")

	// Returns the default value for showHiddenFiles parameter if there is no defined cookie
	if showHiddenFilesCookie == nil {
		return false
	}

	// Stores the cookie value to a bool type
	var showHiddenFiles, _ = strconv.ParseBool(showHiddenFilesCookie.Value)

	return showHiddenFiles
}
