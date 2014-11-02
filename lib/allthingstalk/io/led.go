package io

/**
 * Returns an IODevice pointer with the correct Type and Profile
 * for an LED
 */
func NewLED(config *Config) *IODevice {

	return &IODevice{
		Config: config,
		Type:   "actuator",
		Profile: Profile{
			Type: "int",
		},
	}

}
