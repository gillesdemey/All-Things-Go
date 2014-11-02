package io

/**
 * Returns an IODevice pointer with the correct Type and Profile
 * for an LED
 */
func NewLED(config *Config) *IODevice {

	return &IODevice{
		Name:        config.Name,
		Pin:         config.Pin,
		Id:          config.Id,
		Description: config.Description,
		Type:        "actuator",
		Profile: Profile{
			Type: "int",
		},
	}

}
