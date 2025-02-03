package setting

type DeleteSettingInput struct {
	Key string `json:"key" validate:"required,max=255"`
}

type UpdateSettingInput struct {
	Key   string `json:"key" validate:"required,max=255"`
	Value string `json:"value" validate:"required,max=65535"`
}
