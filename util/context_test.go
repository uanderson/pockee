package util

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserID(t *testing.T) {
	ctx := context.WithValue(context.Background(), "userID", "foo")
	userID := GetUserID(ctx)

	assert.Equal(t, "foo", userID)
}

func Test(t *testing.T) {
	ctx := context.Background()
	userID := GetUserID(ctx)

	assert.Equal(t, userID, "")
}
