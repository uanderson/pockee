package validation

import "github.com/go-playground/validator/v10"

// EchoValidator is a custom validator that uses go playground validator
type EchoValidator struct {
	Validator *validator.Validate
}

// NewEchoValidator returns an instance of a custom echo validator
func NewEchoValidator() *EchoValidator {
	return &EchoValidator{Validator: validator.New()}
}

// Validate any struct through playground validator
func (ev *EchoValidator) Validate(target interface{}) error {
	return ev.Validator.Struct(target)
}
