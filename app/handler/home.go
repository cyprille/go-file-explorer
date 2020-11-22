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
)

// Home Responds to the home path call
func Home(rw http.ResponseWriter, req *http.Request) {
	var rootDir = "/Users/cyprillechauvry/workspace"
	var path = ""

	if len(req.URL.Query()["child"]) > 0 {
		path = req.URL.Query()["child"][0]
	}

	var items, currentDir = filesystem.GetPathContent(rootDir, path)

	type Page struct {
		Title            string
		Items            []string
		CurrentDirectory string
	}

	p := Page{
		Title:            "Home",
		Items:            items,
		CurrentDirectory: currentDir,
	}

	common.Templates = template.Must(template.ParseFiles("templates/filesystem/home.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
