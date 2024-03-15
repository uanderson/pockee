package main

import (
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/firebase"
	"github.com/uanderson/pockee/setting"
	"github.com/uanderson/pockee/validation"
	"net/http"
	"os"
)

// ServiceContainer is a container for all services
type ServiceContainer struct {
	settingService *setting.Service
}

var appDatabase *database.Database
var appFirebase *firebase.Firebase
var appServices ServiceContainer

func main() {
	appFirebase = firebase.New()
	appDatabase = database.New()

	initServices()
	initScheduling()
	initServer()
}

func initServices() {
	appServices = ServiceContainer{
		settingService: setting.NewService(appDatabase),
	}
}

func initScheduling() {
	// nothing scheduled yet
}

func initServer() {
	e := echo.New()
	e.Validator = validation.NewEchoValidator()

	e.Use(validation.ErrorMiddleware)

	e.GET("/status", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	setting.NewApi(e, appFirebase, appServices.settingService).Serve()

	address := os.Getenv("ADDRESS")
	e.Logger.Fatal(e.Start(address))
}
