package bill

import (
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/firebase"
	"log"
)

type Api struct {
	echo     *echo.Echo
	firebase *firebase.Firebase
	service  *Service
}

func NewApi(echo *echo.Echo, firebase *firebase.Firebase, service *Service) *Api {
	return &Api{echo, firebase, service}
}

func (api *Api) Serve() {
	log.Println("serving bills")
}
