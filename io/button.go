package io

// Returns an IODevice pointer with the correct Type and Profile
// for a Button
func NewButton(config *Config) *IODevice {

	return &IODevice{
		Config: config,
		Type:   "sensor",
		Profile: Profile{
			Type: "bool",
		},
	}

}
