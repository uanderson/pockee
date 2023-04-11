package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Errors define a custom error struct to hold validation error messages
type Errors struct {
	Errors []Error `json:"errors"`
}

// Error define a custom error struct to hold validation error information
type Error struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

func (err Error) Error() string {
	return err.Message
}

// NewError creates a new error with a message only
func NewError(message string) Error {
	return Error{Message: message}
}

// ErrorMiddleware acts on every request and tries to translate errors
// to a more standardized structure
func ErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			switch err.(type) {
			case validator.ValidationErrors:
				return c.JSON(http.StatusBadRequest, convertValidationErrors(err.(validator.ValidationErrors)))
			case Error:
				return c.JSON(http.StatusBadRequest, err)
			default:
				return err
			}
		}
		return nil
	}
}

// convertValidationErrors converts any validator.ValidationErrors into
// the common errors generated as response to the api
func convertValidationErrors(err error) (errors Errors) {
	errors = Errors{Errors: []Error{}}

	for _, fieldError := range err.(validator.ValidationErrors) {
		errors.Errors = append(errors.Errors, Error{
			Field:   fieldError.Field(),
			Message: fieldError.Tag(),
		})
	}

	return errors
}
