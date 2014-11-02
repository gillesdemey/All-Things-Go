package io

/**
 * Returns an IODevice pointer with the correct Type and Profile
 * for a Button
 */
func NewButton(config *Config) *IODevice {

	return &IODevice{
		Name:        config.Name,
		Pin:         config.Pin,
		Id:          config.Id,
		Description: config.Description,
		Type:        "sensor",
		Profile: Profile{
			Type: "bool",
		},
	}

}
