package echox

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

type TestStruct struct {
	Name string `validate:"required"`
}

func TestRequestContextWithUserID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("uid", "testUserID")

	ctx := RequestContext(c)

	userID, ok := ctx.Value("userID").(string)

	assert.True(t, ok)
	assert.Equal(t, "testUserID", userID)
}

func TestRequestContextWithoutUserID(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctx := RequestContext(c)

	userID, ok := ctx.Value("userID").(string)

	assert.False(t, ok)
	assert.Equal(t, "", userID)
}

func TestBindAndValidateWithInvalidInput(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := BindAndValidate(c, &TestStruct{})

	assert.Error(t, err)
}

func TestBindAndValidateWithValidInput(t *testing.T) {
	e := echo.New()
	e.Validator = NewEchoValidator()
	req := httptest.NewRequest(echo.POST, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := BindAndValidate(c, &TestStruct{Name: "Test"})

	assert.NoError(t, err)
}
