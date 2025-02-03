package main

import (
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/bill"
	"github.com/uanderson/pockee/category"
	"github.com/uanderson/pockee/contact"
	"github.com/uanderson/pockee/database"
	"github.com/uanderson/pockee/echox"
	"github.com/uanderson/pockee/firebase"
	"github.com/uanderson/pockee/setting"
	"net/http"
	"os"
)

var billService *bill.Service
var categoryService *category.Service
var contactService *contact.Service
var settingService *setting.Service

var appDatabase *database.Database
var appFirebase *firebase.Firebase

func main() {
	appFirebase = firebase.New()
	appDatabase = database.New()

	initServices()
	initScheduling()
	initServer()
}

func initServices() {
	categoryService = category.NewService(appDatabase)
	contactService = contact.NewService(appDatabase)
	settingService = setting.NewService(appDatabase)
	billService = bill.NewService(appDatabase, categoryService, contactService)
}

func initScheduling() {
	// nothing scheduled yet
}

func initServer() {
	e := echo.New()
	e.Validator = echox.NewEchoValidator()
	e.Use(echox.ErrorMiddleware)

	e.GET("/status", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "OK")
	})

	bill.NewApi(e, appFirebase, billService).Serve()
	category.NewApi(e, appFirebase, categoryService).Serve()
	contact.NewApi(e, appFirebase, contactService).Serve()
	setting.NewApi(e, appFirebase, settingService).Serve()

	address := os.Getenv("POCKEE_ADDRESS")
	e.Logger.Fatal(e.Start(address))
}
