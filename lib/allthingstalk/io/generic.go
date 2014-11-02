package io

/**
 * IODevice configuration structure
 */
type Config struct {
	Name        string
	Pin         byte
	Id          string
	Description string
}

/**
 * IODevice structure
 */
type IODevice struct {
	Name        string
	Pin         byte
	Id          string
	Description string
	Type        string
	Profile     Profile
}

/**
 * IODevice's profile structure
 */
type Profile struct {
	Type string `json:"type"`
}

/**
 * The payload we need to construct to register an IODevice with the API endpoint
 */
type RegisterPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Type        string  `json:"is"`
	Profile     Profile `json:"profile"`
	DeviceId    string  `json:"deviceId"`
}
