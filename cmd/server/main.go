package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
        }
	}()

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {fmt.Fprintf(w, "Привет Мир")}).Methods("GET")

    configServer(router).ListenAndServe()
}

func configServer (router *mux.Router) *http.Server {
	return 	&http.Server{
        Handler:      router,
        Addr:         ":8000",
        WriteTimeout: 5 * time.Second,
        ReadTimeout:  5 * time.Second,
		MaxHeaderBytes: 1 << 20,
    }
}