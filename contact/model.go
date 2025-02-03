package contact

import (
	"github.com/uanderson/pockee/contact/dao"
	"github.com/uanderson/pockee/model"
)

type CreateContactInput struct {
	Name   string  `json:"name" validate:"required,min=2,max=255"`
	Email  *string `json:"email" validate:"omitnil,email,max=255"`
	Phone  *string `json:"phone" validate:"omitnil,e164,max=20"`
	PixKey *string `json:"pixKey" validate:"omitnil,max=255"`
}

type UpdateContactInput struct {
	CreateContactInput
	model.IdentifiableInput
}

type DeleteContactInput struct {
	model.IdentifiableInput
}

type ContactOutput struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Email *string `json:"email,omitempty"`
	Phone *string `json:"phone,omitempty"`
}

func ToContactOutput(contact dao.Contact) ContactOutput {
	return ContactOutput{
		ID:    contact.ID,
		Name:  contact.Name,
		Email: contact.Email,
		Phone: contact.Phone,
	}
}

func ToContactOutputs(contacts []dao.Contact) []ContactOutput {
	outputs := make([]ContactOutput, len(contacts))
	for i, b := range contacts {
		outputs[i] = ToContactOutput(b)
	}

	return outputs
}
