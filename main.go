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
	handler "go-file-explorer/app/handler"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
)

// main Boostraps the app
func main() {
	flag.Parse()
	defer glog.Flush()

	r := mux.NewRouter()
	r.HandleFunc(`/settings`, handler.SettingsHandler)
	r.HandleFunc(`/api/open/{rest:.+}`, handler.OpenFileHandler)
	r.HandleFunc(`/api/navigation/{rest:.+}`, handler.PathHandler)
	r.HandleFunc(`/api/navigation/`, handler.HomeHandler)
	r.HandleFunc(`/api/`, handler.HomeHandler)
	r.HandleFunc(`/`, handler.HomeHandler)
	http.Handle(`/`, r)

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileServer)

	err := http.ListenAndServe(":"+common.GetParam("SERVER_PORT"), nil)

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
			// We may not always want to Status ok, but for this example we will
			common.LogAccess(w, req, elapsedTime)
		case "POST":
			// We might use http.StatusCreated here
		}
	})
}
