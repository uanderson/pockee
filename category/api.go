package category

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
	api.echo.GET("/categories", api.firebase.Protect(api.getCategories))
	api.echo.POST("/categories.create", api.firebase.Protect(api.createCategory))
	api.echo.POST("/categories.delete", api.firebase.Protect(api.deleteCategory))
	api.echo.POST("/categories.update", api.firebase.Protect(api.updateCategory))
}

func (api *Api) getCategories(ctx echo.Context) (err error) {
	categories, err := api.service.GetCategories(echox.RequestContext(ctx))
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, ToCategoryOutputs(categories))
}

func (api *Api) createCategory(ctx echo.Context) error {
	var input CreateCategoryInput

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.CreateCategory(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (api *Api) deleteCategory(ctx echo.Context) error {
	var input DeleteCategoryInput

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.DeleteCategory(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (api *Api) updateCategory(ctx echo.Context) error {
	var input UpdateCategoryInput

	err := echox.BindAndValidate(ctx, &input)
	if err != nil {
		return err
	}

	err = api.service.UpdateCategory(echox.RequestContext(ctx), input)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusNoContent)
}
