package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorStruct(t *testing.T) {
	err := Error{Message: "Test error message"}

	assert.Equal(t, "Test error message", err.Error())
}

func TestNewErrorStruct(t *testing.T) {
	err := NewError("Test error message")

	assert.Equal(t, "Test error message", err.Message)
}

func TestConvertEmptyValidationErrors(t *testing.T) {
	ve := validator.ValidationErrors{}

	errors := convertValidationErrors(ve)

	assert.Equal(t, 0, len(errors.Errors))
}
