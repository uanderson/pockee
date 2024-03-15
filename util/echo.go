package util

import (
	"context"
	"github.com/labstack/echo/v4"
)

func EchoContext(c echo.Context) context.Context {
	userID := c.Get("uid").(string)
	return context.WithValue(c.Request().Context(), "userID", userID)
}
