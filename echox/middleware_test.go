package echox

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockFieldError struct {
	validator.FieldError
	tag   string
	field string
}

func (e MockFieldError) Tag() string {
	return e.tag
}

func (e MockFieldError) Field() string {
	return e.field
}

func TestErrorMiddlewareWithValidationError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	next := func(ec echo.Context) error {
		return validator.ValidationErrors{
			MockFieldError{field: "bar", tag: "foo"},
		}
	}

	middleware := ErrorMiddleware(next)
	_ = middleware(c)

	responseBody := rec.Body.String()

	assert.Equal(t, "{\"error\":\"inputValidationFailed\",\"validation\":[{\"field\":\"bar\",\"message\":\"foo\"}]}\n", responseBody)
}

func TestErrorMiddlewareWithPrefixedError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	next := func(ec echo.Context) error {
		return errors.New("error::foo")
	}

	middleware := ErrorMiddleware(next)
	_ = middleware(c)

	responseBody := rec.Body.String()

	assert.Equal(t, "{\"error\":\"foo\"}\n", responseBody)
}

func TestErrorMiddlewareWithDefaultError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	next := func(ec echo.Context) error {
		return errors.New("foo")
	}

	middleware := ErrorMiddleware(next)
	_ = middleware(c)

	responseBody := rec.Body.String()

	assert.Equal(t, "{\"error\":\"Blast it all to the moon! The gremlins are at it again ಠ_ಠ\"}\n", responseBody)
}

func TestErrorMiddlewareWithoutError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	next := func(ec echo.Context) error {
		return nil
	}

	middleware := ErrorMiddleware(next)
	err := middleware(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
