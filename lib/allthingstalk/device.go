package allthingstalk

import (
	"github.com/gillesdemey/All-Things-Go/lib/allthingstalk/io"
	"log"
)

type Device struct {
	DeviceId  string
	ClientId  string
	ClientKey string
	IODevices []*io.IODevice
}

/**
 * Constructor for a new Device
 */
func NewDevice(device *Device) *Device {

	device.Setup()

	return device
}

/**
 * Add a generic IODevice
 */
func (device *Device) AddIODevice(ioDevice *io.IODevice) *io.IODevice {
	device.IODevices = append(device.IODevices, ioDevice)
	device.RegisterAsset(ioDevice)

	return ioDevice
}

/**
 * Add an LED to the device configuration
 */
func (device *Device) AddLED(config *io.Config) *io.IODevice {

	led := io.NewLED(config)
	device.AddIODevice(led)

	return led
}

/**
 * Add a Button to the device configuration
 */
func (device *Device) AddButton(config *io.Config) *io.IODevice {

	button := io.NewButton(config)
	device.AddIODevice(button)

	return button
}

/**
 * Set up the appropriate Broker socket connection
 */
func (device *Device) Setup() {

	_, err := NewBroker(device)

	if err != nil {
		log.Fatalf("Could not connect to broker: %s\n", err)
	}

	log.Printf("Successfuly connected to broker")
}
