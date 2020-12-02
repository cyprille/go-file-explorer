/*
 * Project: Go File Explorer
 * File: parameters.go
 * ---
 * Created: 2/12/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

package common

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Initialized parameters toggle
var i = false

// InitParams Initializes the parameters from the .env file
func InitParams() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	i = true
}

// GetParam Returns the parameter value from its name
func GetParam(name string) string {
	if i == false {
		InitParams()
	}

	v := os.Getenv(name)

	if v == "" {
		log.Fatalf("Unable to retrieve value for parameter: " + v)
	}

	return v
}
