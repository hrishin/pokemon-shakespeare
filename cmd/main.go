package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hrishin/pokemon-shakespeare/pkg/pokemon"
)

func main() {
	wait := time.Second * 15

	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{name}", pokemon.GetDescriptionHandler)

	srv := &http.Server{
		Addr:    "0.0.0.0:5000",
		Handler: r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	os.Exit(0)
}
