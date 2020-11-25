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

var showHiddenFiles = false

// SettingsHandler handles the response for the settings page
func SettingsHandler(rw http.ResponseWriter, req *http.Request) {
	// Defines CurrentPage parameter
	common.CurrentPage = "settings"

	v := map[string]interface{}{
		"ShowHiddenFiles": showHiddenFiles,
		"CurrentPage":     common.CurrentPage,
	}

	// Boostraps the template
	common.Templates = template.Must(template.ParseFiles("templates/settings.html", common.LayoutPath))

	// Renders the template
	err := common.Templates.ExecuteTemplate(rw, "base", v)
	common.CheckError(err, 2)
}

// SetShowHiddenFiles Sets the value for the parameter showHiddenFiles
func SetShowHiddenFiles(value bool) {
	showHiddenFiles = value

	fmt.Print(showHiddenFiles)
}
