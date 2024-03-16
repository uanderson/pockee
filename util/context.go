package util

import (
	"context"
	"errors"
)

func GetUserID(c context.Context) (string, error) {
	value := c.Value("userID")

	if value == nil {
		return "", errors.New("userID not found in context")
	}

	return value.(string), nil
}
