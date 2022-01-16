package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func handle() *mux.Router {
	return mux.NewRouter()
}

func main() {
	address := os.Getenv("ADDRESS")

	if len(address) == 0 {
		address = ":8000"
	}

	server := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      handle(),
	}

	log.Printf("Server running: %s\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
