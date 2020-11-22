/*
 * Project: Go File Explorer
 * File: access.go
 * ---
 * Created: 3/11/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

// Note: inspiration for this from https://gist.github.com/cespare/3985516

package common

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang/glog"
)

const (
	logFile = "access_log.txt"
)

// AccessLog Defines the structure of log access request objects
type AccessLog struct {
	ip, method, uri, protocol, host string
	elapsedTime                     time.Duration
}

// LogAccess Creates the log of access request
func LogAccess(w http.ResponseWriter, req *http.Request, duration time.Duration) {
	clientIP := req.RemoteAddr

	if colon := strings.LastIndex(clientIP, ":"); colon != -1 {
		clientIP = clientIP[:colon]
	}

	record := &AccessLog{
		ip:          clientIP,
		method:      req.Method,
		uri:         req.RequestURI,
		protocol:    req.Proto,
		host:        req.Host,
		elapsedTime: duration,
	}

	writeAccessLog(record)
}

// writeAccessLog Writes the log of access request
func writeAccessLog(record *AccessLog) {
	logRecord := "" + record.ip + " " + record.protocol + " " + record.method + ": " + record.uri + ", host: " + record.host + " (load time: " + strconv.FormatFloat(record.elapsedTime.Seconds(), 'f', 5, 64) + " seconds)"
	glog.Infoln(logRecord)
	glog.Flush()
}
