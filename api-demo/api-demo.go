package main

import (
    "flag"
    "fmt"
    "github.com/elauffenburger/technossus-firstfriday-go/api-demo/http-helpers"
    "net/http"

    "github.com/gorilla/mux"
)

var muxFlag *string
var portFlag *string

func init() {
    muxFlag = flag.String("mux", "", "[gorilla,default]")
    portFlag = flag.String("port", "8080", "the port to listen on")

    flag.Usage = func() {
        fmt.Println("\n------------------\nUsage of api-demo: ")
        flag.PrintDefaults()
    }
}

func main() {
    flag.Parse()

    port := *portFlag
    muxer := getMuxFromArgs()

    switch muxer.(type) {
    case *mux.Router:
        fmt.Println("Using gorilla router!")
    case *http.ServeMux:
        fmt.Println("Using default mux!")
    }

    fmt.Printf("Listening on port %s", port)
    http.ListenAndServe(fmt.Sprintf(":%s", port), muxer)
}

func getMuxFromArgs() http.Handler {
    if *muxFlag == "gorilla" {
        return helpers.GetGorillaMux()
    }

    return helpers.GetDefaultMux()
}
