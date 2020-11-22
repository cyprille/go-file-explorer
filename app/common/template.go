/*
 * Project: Go File Explorer
 * File: template.go
 * ---
 * Created: 3/11/2020
 * Author: Cyprille Chauvry
 * ---
 * License: MIT License
 * Copyright © 2020 Cyprille Chauvry
 */

package common

import (
	"html/template"
)

// Templates Allows access to the template pointer from the entire application
var Templates *template.Template

// LayoutPath Defines the path of the layout template page
const LayoutPath string = "templates/layout.html"
