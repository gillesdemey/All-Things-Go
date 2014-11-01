package main

import (
	att "github.com/gillesdemey/All-Things-Go/lib/allthingstalk"
	gp "github.com/gillesdemey/All-Things-Go/lib/grovepi"
	"log"
	"time"
)

func main() {

	deviceId := "jEGkKxDP9INnpsXciju0r9M"
	clientId := "5454c20f30721aacc441ae6a"
	// clientSecret := "f4jjslau4l4"

	device := att.Device{
		DeviceId: deviceId,
		ClientId: clientId,
	}

	log.Printf("Device: %+v\n", device)

	/**
	 * Create the sensors
	 */
	button := att.IODevice{
		Name: "Button",
		Pin:  2,
		Id:   "1",
	}

	diode := att.IODevice{
		Name: "Diode",
		Pin:  3,
		Id:   "2",
	}

	/**
	 * Initialize the GrovePi shield
	 */
	grovepi := *gp.InitGrovePi(0x04)

	grovepi.PinMode(button.Pin, "input")
	grovepi.PinMode(diode.Pin, "output")

	log.Printf("GrovePi: %+v\n", grovepi)

	/**
	 * Send LED commands
	 */

	var status byte

	status, _ = grovepi.DigitalRead(diode.Pin)

	status = invertByte(status)
	grovepi.DigitalWrite(diode.Pin, status)

	time.Sleep(500 * time.Millisecond)

	status = invertByte(status)
	grovepi.DigitalWrite(diode.Pin, status)

}

func invertByte(b byte) byte {
	return byte(int(b) ^ 1)
}
