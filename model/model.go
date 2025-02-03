package model

// IdentifiableInput represents a contact with an ID.
type IdentifiableInput struct {
	ID string `json:"id" validate:"required,len=20"`
}
