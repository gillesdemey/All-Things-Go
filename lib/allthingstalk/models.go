package allthingstalk

type IODevice struct {
	Name string
	Pin  byte
	Id   string
}

type Broker struct {
	Uri string
}

type Connection struct {
	Address string
}

type Device struct {
	DeviceId string
	ClientId string
}
