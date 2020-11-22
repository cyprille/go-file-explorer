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
	type Page struct {
		Title            string
		Directories      []string
		CurrentDirectory string
	}

	files, error := ioutil.ReadDir(".")
	if error != nil {
		log.Fatal(error)
	}

	var directories = []string{}
	for _, f := range files {
		directories = append(directories, f.Name())
	}

	// var directories = map[int]string{}
	// for i, f := range files {
	// 	directories = append(directories, i, f.Name())
	// }

	// fmt.Print(directories)

	var currentDirectory = "/Users/cyprillechauvry/workspace/"

	p := Page{
		Title:            "directory_root",
		Directories:      directories,
		CurrentDirectory: currentDirectory,
	}

	common.Templates = template.Must(template.ParseFiles("templates/directory/root.html", common.LayoutPath))
	err := common.Templates.ExecuteTemplate(rw, "base", p)
	common.CheckError(err, 2)
}
