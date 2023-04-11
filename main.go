package main

import (
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/exchange"
	"github.com/uanderson/pockee/firebase"
	"github.com/uanderson/pockee/pocketsmith"
	"github.com/uanderson/pockee/setting"
	"github.com/uanderson/pockee/validation"
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
	elEcho := echo.New()
	elEcho.Validator = validation.NewEchoValidator()

	elEcho.Use(validation.ErrorMiddleware)

	elEcho.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	setting.NewApi(elEcho).Serve()

	address := os.Getenv("ADDRESS")
	elEcho.Logger.Fatal(elEcho.Start(address))
}
