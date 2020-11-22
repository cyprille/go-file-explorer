/*
 * Project: Go File Explorer
 * File: main.go
 * ---
 * Created: 4/11/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	common "go-file-explorer/app/common"
	directory "go-file-explorer/app/directory"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

// main Boostraps the app
func main() {
	flag.Parse()
	defer glog.Flush()

	router := mux.NewRouter()
	http.Handle("/", httpInterceptor(router))

	router.HandleFunc("/", directory.GetDirectories).Methods("GET")

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileServer)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println(err)
	}
}

// httpInterceptor Handles the application routing
func httpInterceptor(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		router.ServeHTTP(w, req)

		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)

		switch req.Method {
		case "GET":
			// We may not always want to StatusOK, but for the sake of
			// this example we will
			common.LogAccess(w, req, elapsedTime)
		case "POST":
			// here we might use http.StatusCreated
		}

	})
}
