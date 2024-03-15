package util

import "context"

func GetUserID(ctx context.Context) string {
	return ctx.Value("userID").(string)
}
