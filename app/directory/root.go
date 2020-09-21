package root

import (
    "io/ioutil"
    "log"
    "net/http"
    "html/template"
    "go-file-explorer/app/common"
)

func GetDirectories(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
        Directories []string
    }

    files, error := ioutil.ReadDir(".")
    if error != nil {
        log.Fatal(error)
    }

    var directories = []string{}
    for _, f := range files {
        directories = append(directories, f.Name())
    }

    p := Page{
        Title: "directory_root",
        Directories: directories,
    }

    common.Templates = template.Must(template.ParseFiles("templates/directory/root.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err, 2)
}
