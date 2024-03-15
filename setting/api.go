package setting

import (
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/firebase"
	"github.com/uanderson/pockee/util"
	"net/http"
)

// Api holds access to echo and service
type Api struct {
	echo     *echo.Echo
	firebase *firebase.Firebase
	service  *Service
}

// NewApi creates a new instance of Api
func NewApi(echo *echo.Echo, firebase *firebase.Firebase, service *Service) *Api {
	return &Api{echo, firebase, service}
}

// Serve hookup all the endpoints
func (api *Api) Serve() {
	api.echo.POST("/users.settings.update", api.firebase.Protect(api.updateSetting))
}

// updateSetting validates and delegate the update of the user setting to the service
func (api *Api) updateSetting(c echo.Context) (err error) {
	var input UpdateSettingInput
	if err = c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	if err = c.Validate(&input); err != nil {
		return
	}

	setting, err := api.service.UpdateUserSetting(util.EchoContext(c), &input)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, setting)
}
