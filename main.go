package main

import (
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/account"
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

type ServiceContainer struct {
	accountService  *account.Service
	billService     *bill.Service
	categoryService *category.Service
	contactService  *contact.Service
	settingService  *setting.Service
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
		accountService:  account.NewService(appDatabase),
		billService:     bill.NewService(appDatabase),
		categoryService: category.NewService(appDatabase),
		contactService:  contact.NewService(appDatabase),
		settingService:  setting.NewService(appDatabase),
	}
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

	account.NewApi(e, appFirebase, appServices.accountService).Serve()
	bill.NewApi(e, appFirebase, appServices.billService).Serve()
	category.NewApi(e, appFirebase, appServices.categoryService).Serve()
	contact.NewApi(e, appFirebase, appServices.contactService).Serve()
	setting.NewApi(e, appFirebase, appServices.settingService).Serve()

	address := os.Getenv("ADDRESS")
	e.Logger.Fatal(e.Start(address))
}
