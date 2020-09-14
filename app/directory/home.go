package directory

import (
    "net/http"
    "html/template"

    "go-file-explorer/app/common"
)

func GetHomePage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }

    p := Page{
        Title: "directory_home",
    }

    common.Templates = template.Must(template.ParseFiles("templates/directory/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}