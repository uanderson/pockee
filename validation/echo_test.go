package validation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestStruct struct {
	Name string `validate:"required"`
}

func TestEchoValidate(t *testing.T) {
	ev := NewEchoValidator()
	err := ev.Validate(&TestStruct{Name: "Test"})

	assert.NoError(t, err)
}

func TestEchoValidateWithError(t *testing.T) {
	ev := NewEchoValidator()
	err := ev.Validate(&TestStruct{})

	assert.Error(t, err)
}
