package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hrishin/pokemon-shakespeare/pkg/pokemon"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/pokemon/{name}", pokemon.GetByNameHandler)
	http.Handle("/", r)

	http.ListenAndServe(":5000", nil)
}
