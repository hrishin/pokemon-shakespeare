package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/hrishin/pokemon-shakespeare/pkg/pokemon"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("pokemon")

func main() {
	wait := time.Second * 15
	port := 5000

	r := mux.NewRouter()
	r.HandleFunc("/pokemon/{name}", pokemon.GetDescriptionHandler)

	srv := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", port),
		Handler: r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Infof("starting the server and listening on port %d", port)
		if err := srv.ListenAndServe(); err != nil {
			log.Error(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Info("stopping the server")
	os.Exit(0)
}
