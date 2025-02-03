package setting

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
	api.echo.POST("/settings.delete", api.firebase.Protect(api.deleteSetting))
	api.echo.POST("/settings.update", api.firebase.Protect(api.updateSetting))
}

func (api *Api) deleteSetting(ctx echo.Context) error {
	var input DeleteSettingInput

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.DeleteSetting(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (api *Api) updateSetting(ctx echo.Context) error {
	var input UpdateSettingInput

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.UpdateUserSetting(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
