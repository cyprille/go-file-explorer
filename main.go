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
	"log"
	"net/http"
	"os"
	"time"

	common "go-file-explorer/app/common"
	handler "go-file-explorer/app/handler"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// serverPort is the listening port of the server
var serverPort string

// Initializes the parameters from .env file
// @TODO: put this in dedicated package
func initParams() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	serverPort = os.Getenv("SERVER_PORT")
}

// main Boostraps the app
func main() {
	// Bootstraps the parameters initialization
	// @TODO: put this in dedicated package
	initParams()

	flag.Parse()
	defer glog.Flush()

	r := mux.NewRouter()
	r.HandleFunc(`/settings`, handler.SettingsHandler)
	r.HandleFunc(`/api/open/{rest:[A-zÀ-ú0-9=\-\/.% ]+}`, handler.OpenFileHandler)
	r.HandleFunc(`/api/navigation/{rest:[A-zÀ-ú0-9=\-\/.% ]+}`, handler.PathHandler)
	r.HandleFunc(`/api/navigation/`, handler.HomeHandler)
	r.HandleFunc(`/api/`, handler.HomeHandler)
	r.HandleFunc(`/`, handler.HomeHandler)
	http.Handle(`/`, r)

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileServer)

	err := http.ListenAndServe(":"+serverPort, nil)

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
