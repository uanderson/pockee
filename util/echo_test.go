package util

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoContext(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.Set("uid", "testUserID")

	ctx := EchoContext(c)

	userID, ok := ctx.Value("userID").(string)

	assert.True(t, ok)
	assert.Equal(t, "testUserID", userID)
}
