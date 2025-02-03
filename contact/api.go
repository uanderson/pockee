package contact

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
	api.echo.GET("/contacts", api.firebase.Protect(api.getContacts))
	api.echo.POST("/contacts.create", api.firebase.Protect(api.createContact))
	api.echo.POST("/contacts.update", api.firebase.Protect(api.updateContact))
	api.echo.POST("/contacts.delete", api.firebase.Protect(api.deleteContact))
}

func (api *Api) getContacts(ctx echo.Context) (err error) {
	contacts, err := api.service.GetContacts(echox.RequestContext(ctx))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, ToContactOutputs(contacts))
}

func (api *Api) createContact(ctx echo.Context) error {
	var input CreateContactInput

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.CreateContact(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (api *Api) updateContact(ctx echo.Context) error {
	var input UpdateContactInput

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.UpdateContact(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (api *Api) deleteContact(ctx echo.Context) error {
	var input DeleteContactInput

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.DeleteContact(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
