package main

import (
	"flag"
	"fmt"
	"net/http"
	"user/api-demo/http-helpers"

	"github.com/gorilla/mux"
)

var muxFlag *string

func init() {
	muxFlag = flag.String("mux", "", "[gorilla,default]")
}

func main() {
	flag.Parse()

	muxer := getMuxFromArgs()

	switch muxer.(type) {
	case *mux.Router:
		fmt.Println("Using gorilla router!")
	case *http.ServeMux:
		fmt.Println("Using default mux!")
	}

	http.ListenAndServe(":8080", muxer)
}

func getMuxFromArgs() http.Handler {
	if *muxFlag == "gorilla" {
		return helpers.GetGorillaMux()
	}

	return helpers.GetDefaultMux()
}
