package directory

import (
    "net/http"
    "html/template"

    "go-file-explorer/app/common"

    "github.com/gorilla/mux"
)

func GetViewPage(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title  string
        DirectoryId string
    }

    params := mux.Vars(req)
    directoryId := params["id"]

    p := Page{
        Title: "directory_view",
        DirectoryId: directoryId,
    }

    common.Templates = template.Must(template.ParseFiles("templates/directory/view.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}