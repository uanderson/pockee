package main

import (
	"github.com/gorilla/mux"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/exchange"
	"log"
	"net/http"
	"os"
	"time"
)

func schedule() {
	exchange.Schedule()
}

func handle() *mux.Router {
	return mux.NewRouter()
}

func serve() {
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

	log.Printf("server running: %v\n", server.Addr)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	database.Init()
	schedule()
	serve()
}
