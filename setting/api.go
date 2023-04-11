package setting

import (
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/firebase"
	"net/http"
)

// Api holds access to echo and service
type Api struct {
	echo    *echo.Echo
	service *Service
}

// NewApi creates a new instance of Api
func NewApi(e *echo.Echo) *Api {
	return &Api{echo: e, service: NewService()}
}

// Serve hookup all the endpoints
func (api *Api) Serve() {
	api.echo.POST("/users.settings.update", firebase.Protect(api.updateSetting))
}

// updateSetting validates and delegate the update of the user
// setting to the service
func (api *Api) updateSetting(c echo.Context) (err error) {
	var input UpdateSettingInput
	if err = c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	if err = c.Validate(&input); err != nil {
		return
	}

	setting, err := api.service.UpdateUserSetting(&input, c.Get("uid").(string))
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, setting)
}
