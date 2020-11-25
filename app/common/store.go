/*
 * Project: Go File Explorer
 * File: store.go
 * ---
 * Created: 25/11/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

package common

import (
	"github.com/rapidloop/svk"
)

var storeFile = "/sessions.db"

// Read returns the value from the store
func Read(key string) {
	// Opens the store
	store, err := svk.Open(storeFile)

	// Fetches from boltdb and does gob decode
	err := svk.Get(sessionId, &info)

	// Closes the store
	store.Close()
}

// Write saves the value in the store
func Write(key string, value string) {
	// Opens the store
	store, err := svk.Open(storeFile)

	// Encodes the value with gob and updates the boltdb
	err := svk.Put(sessionId, info)

	// Closes the store
	store.Close()
}

// Delete drops the value from the store
func Delete(key string) {
	// Opens the store
	store, err := svk.Open(storeFile)

	// Deletes seeks in boltdb and deletes the record
	err := svk.Delete(sessionId)

	// Closes the store
	store.Close()
}
