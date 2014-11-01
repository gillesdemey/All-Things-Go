package allthingstalk

type Profile struct {
	Type string `json:"type"`
}

type IODevice struct {
	Name        string  `json:"name"`
	Pin         byte    `json:"-"`
	Id          string  `json:"deviceId"`
	Description string  `json:"description,omitempty"`
	Type        string  `json:"is"`
	Profile     Profile `json:"profile"`
}

type RegisterPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Type        string  `json:"is"`
	Profile     Profile `json:"profile"`
	DeviceId    string  `json:"deviceId"`
}

type Broker struct {
	Uri string
}

type Device struct {
	DeviceId  string
	ClientId  string
	ClientKey string
}
