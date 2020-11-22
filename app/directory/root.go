/*
 * Project: Go File Explorer
 * File: root.go
 * ---
 * Created: 4/11/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

package directory

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"go-file-explorer/app/common"
)

// GetDirectories returns the root directories
func GetDirectories(rw http.ResponseWriter, req *http.Request) {
	var currentDirectory = "/Users/cyprillechauvry/workspace/"

	type Page struct {
		Title            string
		Items            []string
		CurrentDirectory string
	}

	files, error := ioutil.ReadDir(currentDirectory)
	if error != nil {
		log.Fatal(error)
	}

	var items = []string{}
	for _, f := range files {
		items = append(items, f.Name())
	}

	p := Page{
		Title:            "directory_root",
		Items:            items,
		CurrentDirectory: currentDirectory,
	}

	common.Templates = template.Must(template.ParseFiles("templates/directory/root.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
