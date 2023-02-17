package main

import (
	// "encoding/json"
	// "fmt"
	"log"
	"net/http"
	"rest/controller"

	"github.com/gorilla/mux"
)

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)

	log.Fatal(http.ListenAndServe(":3001", router))
}

func main() {
	controller.Start()
}
