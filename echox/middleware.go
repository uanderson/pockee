package echox

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/uanderson/pockee/errorsx"
	"net/http"
	"strings"
)

func ErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		err := next(ctx)
		if err == nil {
			return nil
		}

		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			return handleValidationError(ctx, validationErrors)
		}

		return handleDefaultError(ctx, err)
	}
}

func handleValidationError(ctx echo.Context, err validator.ValidationErrors) error {
	responseError := errorsx.ResponseError{
		Error:      strings.TrimPrefix(errorsx.InputValidationFailed.Error(), "error::"),
		Validation: make([]errorsx.ValidationError, len(err)),
	}

	for i, fieldError := range err {
		responseError.Validation[i] = errorsx.ValidationError{
			Field:   strings.ToLower(fieldError.Field()[:1]) + fieldError.Field()[1:],
			Message: fieldError.Tag(),
		}
	}

	return ctx.JSON(http.StatusBadRequest, responseError)
}

func handleDefaultError(ctx echo.Context, err error) error {
	errMessage := err.Error()

	if strings.HasPrefix(errMessage, "error::") {
		errMessage = strings.TrimPrefix(errMessage, "error::")
	} else {
		log.Error(err)
		errMessage = "Blast it all to the moon! The gremlins are at it again ಠ_ಠ"
	}

	return ctx.JSON(http.StatusBadRequest, errorsx.ResponseError{
		Error: errMessage,
	})
}
