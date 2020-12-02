/*
 * Project: Go File Explorer
 * File: home.go
 * ---
 * Created: 22/11/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

package handler

import (
	"fmt"
	"go-file-explorer/app/common"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

// SettingsHandler handles the response for the settings page
func SettingsHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		// Defines CurrentPage parameter
		common.CurrentPage = "settings"

		v := map[string]interface{}{
			"AppTitle":        common.GetParam("APP_TITLE"),
			"CurrentPage":     common.CurrentPage,
			"ShowHiddenFiles": GetCookie(req, "show-hidden-files"),
			"DarkMode":        GetCookie(req, "dark-mode"),
		}

		// Boostraps the template
		common.Templates = template.Must(template.ParseFiles("templates/settings.html", common.LayoutPath))

		// Renders the template
		err := common.Templates.ExecuteTemplate(rw, "base", v)
		common.CheckError(err, 2)

	case "POST":
		// Parses the form and checks errors
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(rw, "ParseForm() err: %v", err)
			return
		}

		// Retrieves and stores the show-hidden-files and dark-mode cookies
		setCookie(rw, req, "show-hidden-files")
		setCookie(rw, req, "dark-mode")

		// Redirects to the parent page
		http.Redirect(rw, req, req.Header.Get("Referer"), 302)
	}
}

// Sets a cookie value in the response
func setCookie(rw http.ResponseWriter, req *http.Request, name string) {
	// Inits the default value
	var c bool = false

	// Retrieves the cookie value from the request
	v := req.FormValue(name)

	// Sets the value of the cookie
	if v == "on" {
		c = true
	}

	// Sets the expiration date for the cookie
	e := time.Now().Add(5 * 365 * 24 * time.Hour)

	// Defines the cookie values
	cp := http.Cookie{Name: name, Value: strconv.FormatBool(c), Expires: e}

	// Stores the cookie parameter
	http.SetCookie(rw, &cp)
}

// GetCookie Returns if the cookie is activated or not
func GetCookie(req *http.Request, name string) bool {
	// Retrieves the cookie parameter
	c, _ := req.Cookie(name)

	// Returns the default value for parameter if there is no defined cookie
	if c == nil {
		return false
	}

	// Stores the cookie value to a bool type
	var cv, _ = strconv.ParseBool(c.Value)

	return cv
}
