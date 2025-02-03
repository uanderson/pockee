package bill

import (
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/echox"
	"github.com/uanderson/pockee/firebase"
	"net/http"
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
	api.echo.POST("/bills.create", api.firebase.Protect(api.createBill))
	api.echo.POST("/recurring-bills.create", api.firebase.Protect(api.createRecurringBill))
}

func (api *Api) createBill(ctx echo.Context) error {
	var input CreateBill

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.CreateBill(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (api *Api) createRecurringBill(ctx echo.Context) error {
	var input CreateRecurringBill

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.CreateRecurringBill(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
