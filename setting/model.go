package setting

// UpdateSettingInput holds information on which data should be updated for a specific setting
type UpdateSettingInput struct {
	Key   string `json:"key" validate:"required,max=255"`
	Value string `json:"value" validate:"required,max=65535"`
}
