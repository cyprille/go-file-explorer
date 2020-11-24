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

// main Boostraps the app
func main() {
	flag.Parse()
	defer glog.Flush()

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables SITE_TITLE and DB_HOST
	appTitle := os.Getenv("APP_TITLE")
	serverPort := os.Getenv("SERVER_PORT")
	rootDir := os.Getenv("ROOT_DIR")

	fmt.Printf("godotenv : %s = %s \n", "App Title", appTitle)
	fmt.Printf("godotenv : %s = %s \n", "Server Port", serverPort)
	fmt.Printf("godotenv : %s = %s \n", "Root Dir", rootDir)

	r := mux.NewRouter()
	r.HandleFunc(`/api/open/{rest:[A-zÀ-ú0-9=\-\/.% ]+}`, handler.OpenFile)
	r.HandleFunc(`/api/navigation/{rest:[A-zÀ-ú0-9=\-\/.% ]+}`, handler.GoToPath)
	r.HandleFunc(`/api/navigation/`, handler.GoHome)
	r.HandleFunc(`/api/`, handler.GoHome)
	r.HandleFunc(`/`, handler.GoHome)
	http.Handle(`/`, r)

	fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fileServer)

	// @TODO: refacto to handle this in project parameters
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
			// We may not always want to Status ok, but for this example we will
			common.LogAccess(w, req, elapsedTime)
		case "POST":
			// We might use http.StatusCreated here
		}

	})
}
