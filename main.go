package main

import (
	att "github.com/gillesdemey/All-Things-Go/lib/allthingstalk"
	gp "github.com/gillesdemey/All-Things-Go/lib/grovepi"
	"log"
	"time"
)

func main() {

	device := att.Device{
		DeviceId:  "jEGkKxDP9INnpsXciju0r9M",
		ClientId:  "5454c20f30721aacc441ae6a",
		ClientKey: "f4jjslau4l4",
	}

	device.Setup()

	log.Printf("Device: %+v\n", device)

	/**
	 * Create the sensors
	 */
	button := att.IODevice{
		Name:        "Button",
		Pin:         2,
		Id:          "my-button",
		Type:        "sensor",
		Description: "This is my button",
		Profile: att.Profile{
			Type: "bool",
		},
	}

	led := att.IODevice{
		Name:        "LED",
		Pin:         3,
		Id:          "my-led",
		Type:        "actuator",
		Description: "This is my LED",
		Profile: att.Profile{
			Type: "int",
		},
	}

	/**
	 * Register IO Devices
	 */
	device.RegisterAsset(&button)
	device.RegisterAsset(&led)

	/**
	 * Initialize the GrovePi shield
	 */
	grovepi := *gp.InitGrovePi(0x04)

	grovepi.PinMode(button.Pin, "input")
	grovepi.PinMode(led.Pin, "output")

	log.Printf("GrovePi: %+v\n", grovepi)

	/**
	 * Send LED commands
	 */

	var status byte

	status, _ = grovepi.DigitalRead(led.Pin)

	status = invertByte(status)
	grovepi.DigitalWrite(led.Pin, status)

	time.Sleep(500 * time.Millisecond)

	status = invertByte(status)
	grovepi.DigitalWrite(led.Pin, status)

}

func invertByte(b byte) byte {
	return byte(int(b) ^ 1)
}
