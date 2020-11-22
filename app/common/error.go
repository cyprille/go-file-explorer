/*
 * Project: Go File Explorer
 * File: error.go
 * ---
 * Created: 3/11/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

package common

import (
	"log"
	"runtime"

	"github.com/golang/glog"
)

// CheckError Handles the type of error
/*
 * 0 = Info
 * 1 = Warning
 * 2 = Error - should be most common
 * 3 = Fatal
 */
func CheckError(err error, level int) {
	if err != nil {
		var stack [4096]byte

		runtime.Stack(stack[:], false)
		log.Printf("%q\n%s\n", err, stack[:])

		switch level {
		case 0:
			glog.Infoln(err)
		case 1:
			glog.Warningln(err)
		case 2:
			glog.Errorln(err)
		case 3:
			glog.Fatalln(err)
		}

		glog.Flush()
	}
}
