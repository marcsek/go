package controller

import (
	"context"
	"fmt"
	"log"
	"net/http"
	controller "rest/controller/quote"

	"github.com/gorilla/mux"
)

var router *mux.Router

const port = 3001

func initHandlers() {
	router.HandleFunc("/api/quote", controller.GetQuote).Methods("GET")
	router.HandleFunc("/api/character", controller.GetCharacter).Methods("POST")
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func Start() {
	router = mux.NewRouter()
	router.Headers("AccessControl-Allow-Origin", "*")
	router.Use(Middleware)

	initHandlers()
	fmt.Printf("Sever running on port %v (http://localhost:%v)", port, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), router))
}
