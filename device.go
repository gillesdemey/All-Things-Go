package allthingstalk

import (
	"errors"
	"github.com/gillesdemey/All-Things-Go/io"
	"log"
)

// Predefined error types
var ErrNotFound = errors.New("NOTFOUND")

// Device structure
type Device struct {
	DeviceID  string
	ClientID  string
	ClientKey string
	IODevices []*io.IODevice
}

// Constructor for a new Device
func NewDevice(device *Device) *Device {
	device.Setup()
	return device
}

// Add a generic IODevice
func (device *Device) AddIODevice(ioDevice *io.IODevice) *io.IODevice {
	device.IODevices = append(device.IODevices, ioDevice)
	device.RegisterAsset(ioDevice)

	return ioDevice
}

// Add an LED to the device configuration
func (device *Device) NewLED(config *io.Config) *io.IODevice {

	led := io.NewLED(config)
	device.AddIODevice(led)

	return led
}

// Add a Button to the device configuration
func (device *Device) NewButton(config *io.Config) *io.IODevice {

	button := io.NewButton(config)
	device.AddIODevice(button)

	return button
}

// Searches for a device in the devices' list with the specified unique id
func (device *Device) GetIODeviceByID(id string) (*io.IODevice, error) {

	for _, ioDevice := range device.IODevices {
		if ioDevice.Id == id {
			return ioDevice, nil
		}
	}

	return nil, ErrNotFound
}

// Set up the appropriate Broker socket connection
func (device *Device) Setup() {

	_, err := NewBroker(device)

	if err != nil {
		log.Fatalf("Could not connect to broker: %s\n", err)
	}

	log.Printf("Successfuly connected to broker")
}
