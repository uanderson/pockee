package echox

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/uanderson/pockee/errorsx"
)

type EchoValidator struct {
	Validator *validator.Validate
}

func NewEchoValidator() *EchoValidator {
	return &EchoValidator{Validator: validator.New()}
}

func (ev *EchoValidator) Validate(target interface{}) error {
	return ev.Validator.Struct(target)
}

func RequestContext(ctx echo.Context) context.Context {
	userID := ctx.Get("uid")
	if userID == nil {
		return ctx.Request().Context()
	}

	return context.WithValue(ctx.Request().Context(), "userID", userID.(string))
}

func GetUserID(ctx context.Context) string {
	return ctx.Value("userID").(string)
}

func BindAndValidate(ctx echo.Context, input interface{}) error {
	if err := ctx.Bind(input); err != nil {
		return errorsx.InvalidInputData
	}

	if err := ctx.Validate(input); err != nil {
		return err
	}

	return nil
}
