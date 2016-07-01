package helpers

import (
    "github.com/elauffenburger/technossus-firstfriday-go/api-demo/http-handlers"
    "net/http"

    "github.com/gorilla/mux"
)

func GetDefaultMux() http.Handler {
    x := http.NewServeMux()
    x.Handle("/", HandlerFn(handlers.DefaultHandler).ToHandler())
    x.Handle("/home-alt", http.Handler(HandlerFn(handlers.DefaultHandler)))
    x.Handle("/parse/json", MakeHandler(handlers.ParseJSONHandlerGET))

    return x
}

func GetGorillaMux() http.Handler {
    x := mux.NewRouter()

    get := x.Methods("GET").Subrouter()
    get.Path("/").HandlerFunc(handlers.DefaultHandler)
    get.Path("/home-alt").HandlerFunc(handlers.DefaultHandler)
    get.Path("/parse/{parser}").HandlerFunc(handlers.ParseJSONHandlerGET)

    post := x.Methods("POST").Subrouter()
    post.Path("/parse/{parser}").HandlerFunc(handlers.ParseJSONHandlerPOST)

    return x
}
