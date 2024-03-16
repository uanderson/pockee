package util

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPresentUserID(t *testing.T) {
	ctx := context.WithValue(context.Background(), "userID", "foo")
	userID, err := GetUserID(ctx)

	assert.NoError(t, err)
	assert.Equal(t, "foo", userID)
}

func TestMissingUserID(t *testing.T) {
	ctx := context.Background()
	userID, err := GetUserID(ctx)

	assert.Error(t, err)
	assert.Equal(t, "", userID)
}
