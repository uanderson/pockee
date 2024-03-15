package util

import "context"

func GetUserID(c context.Context) string {
	value := c.Value("userID")

	if value == nil {
		return ""
	}

	return value.(string)
}
