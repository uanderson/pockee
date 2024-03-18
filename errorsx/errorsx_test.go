package errorsx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidationErrorError(t *testing.T) {
	err := ValidationError{
		Field:   "testField",
		Message: "testMessage",
	}

	assert.Equal(t, "testMessage", err.Error())
}

func TestResponseErrorWithValidationError(t *testing.T) {
	err := ResponseError{
		Error: "testError",
		Validation: []ValidationError{
			{
				Field:   "testField",
				Message: "testMessage",
			},
		},
	}

	assert.Equal(t, "testError", err.Error)
	assert.Equal(t, "testField", err.Validation[0].Field)
	assert.Equal(t, "testMessage", err.Validation[0].Message)
}

func TestResponseErrorWithoutValidationError(t *testing.T) {
	err := ResponseError{
		Error: "testError",
	}

	assert.Equal(t, "testError", err.Error)
	assert.Nil(t, err.Validation)
}
