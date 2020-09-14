package main

import (
    "fmt"
    "flag"
    "net/http"
    "time"

    "go-file-explorer/app/common"
    "go-file-explorer/app/home"
    "go-file-explorer/app/directory"

    "github.com/golang/glog"
    "github.com/gorilla/mux"
)

func main() {
    flag.Parse()
    defer glog.Flush()

    router := mux.NewRouter()
    http.Handle("/", httpInterceptor(router))

    router.HandleFunc("/", home.GetHomePage).Methods("GET")
    // router.HandleFunc("/directory{_:/?}", directory.GetHomePage).Methods("GET")
    router.HandleFunc("/directory{_:/?}", directory.GetDirectories).Methods("GET")

    router.HandleFunc("/directory/view/{id:[0-9]+}", directory.GetViewPage).Methods("GET")
    router.HandleFunc("/directory/{id:[0-9]+}", directory.GetViewPage).Methods("GET")

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    err := http.ListenAndServe(":8080", nil)

    if err != nil {
        fmt.Println(err)
    }
}

func httpInterceptor(router http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        startTime := time.Now()

        router.ServeHTTP(w, req)

        finishTime := time.Now()
        elapsedTime := finishTime.Sub(startTime)

        switch req.Method {
        case "GET":
            // We may not always want to StatusOK, but for the sake of
            // this example we will
            common.LogAccess(w, req, elapsedTime)
        case "POST":
            // here we might use http.StatusCreated
        }

    })
}
