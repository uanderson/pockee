package setting

type BaseSettingInput struct {
	Key string `json:"key" validate:"required,max=255"`
}

type DeleteSettingInput struct {
	BaseSettingInput
}

type UpdateSettingInput struct {
	BaseSettingInput
	Value string `json:"value" validate:"required,max=65535"`
}
