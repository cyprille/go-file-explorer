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
)

// ShowHiddenFiles Handles the possibility to display hidden files or not
var ShowHiddenFiles = false

// SettingsHandler handles the response for the settings page
func SettingsHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// Defines CurrentPage parameter
		common.CurrentPage = "settings"

		v := map[string]interface{}{
			"ShowHiddenFiles": ShowHiddenFiles,
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

		// Sets the value of the parameter
		if showHiddenFilesValue == "on" {
			ShowHiddenFiles = true
		} else {
			ShowHiddenFiles = false
		}

		// Redirects to the previous page
		http.Redirect(rw, req, req.Header.Get("Referer"), 302)
	}
}
