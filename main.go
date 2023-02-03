package main

import (
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/exchange"
	"github.com/uanderson/pockee/firebase"
	"github.com/uanderson/pockee/pocketsmith"
	"net/http"
	"os"
)

func main() {
	firebase.Init()
	database.Init()

	schedule()
	serve()
}

func schedule() {
	exchange.Schedule()
	pocketsmith.Schedule()
}

func serve() {
	e := echo.New()

	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	address := os.Getenv("ADDRESS")
	e.Logger.Fatal(e.Start(address))
}
