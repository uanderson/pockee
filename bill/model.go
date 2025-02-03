package bill

type CreateBill struct {
	Description string  `json:"description" validate:"required,min=2,max=255"`
	Type        string  `json:"type" validate:"required,oneof=PIX BOLETO"`
	DueAt       string  `json:"dueAt" validate:"required,datetime=2006-01-02"`
	Amount      float64 `json:"amount" validate:"required,gt=0,lt=9999999999"`
	ContactID   string  `json:"contactId" validate:"required,len=20"`
	CategoryID  string  `json:"categoryId" validate:"required,len=20"`
	Barcode     *string `json:"barcode" validate:"required_if=Type BOLETO,omitnil,len=44"`
}

type CreateRecurringBill struct {
	Description string  `json:"description" validate:"required,min=2,max=255"`
	Type        string  `json:"type" validate:"required,oneof=PIX"`
	StartAt     string  `json:"startAt" validate:"required,datetime=2006-01-02"`
	EndAt       *string `json:"endAt" validate:"omitnil,datetime=2006-01-02"`
	Amount      float64 `json:"amount" validate:"required,gt=0,lt=9999999999"`
	Interval    string  `json:"interval" validate:"required,oneof=MONTHLY"`
	ContactID   string  `json:"contactId" validate:"required,len=20"`
	CategoryID  string  `json:"categoryId" validate:"required,len=20"`
}
