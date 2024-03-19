package errorsx

import (
	"errors"
)

// Application errors
var InputValidationFailed = errors.New("error::inputValidationFailed")
var InvalidAuthorizationToken = errors.New("error::invalidAuthorizationToken")
var InvalidInputData = errors.New("error::invalidInputData")
var MissingAuthorizationToken = errors.New("error::missingAuthorizationToken")
var UnauthorizedAccess = errors.New("error::unauthorizedAccess")

// Category errors
var CategoryNotFound = errors.New("error::categoryNotFound")

// Settings errors
var SettingNotFound = errors.New("error::settingNotFound")

type ValidationError struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type ResponseError struct {
	Error      string            `json:"error"`
	Validation []ValidationError `json:"validation,omitempty"`
}

func (err ValidationError) Error() string {
	return err.Message
}
