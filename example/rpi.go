package main

import (
	att "github.com/gillesdemey/All-Things-Go"
	io "github.com/gillesdemey/All-Things-Go/io"
	gp "github.com/gillesdemey/All-Things-Go/lib/grovepi"
	"time"
)

func main() {

	// Create a new device
	device := att.NewDevice(&att.Device{
		DeviceId:  "jEGkKxDP9INnpsXciju0r9M",
		ClientId:  "5454c20f30721aacc441ae6a",
		ClientKey: "f4jjslau4l4",
	})

	// Add your IO devices (sensors, buttons, etc.)
	button := device.NewButton(&io.Config{
		Description: "This is my button",
		Id:          "my-button",
		Name:        "Button",
		Pin:         2,
	})

	led := device.NewLED(&io.Config{
		Description: "This is my LED",
		Id:          "my-led",
		Name:        "LED",
		Pin:         3,
	})

	// Initialize the GrovePi shield
	grovepi := *gp.InitGrovePi(0x04)

	grovepi.PinMode(button.Pin, "input")
	grovepi.PinMode(led.Pin, "output")

	// Send LED commands
	var status byte

	status, _ = grovepi.DigitalRead(led.Pin)

	status = invertBit(status)
	grovepi.DigitalWrite(led.Pin, status)

	time.Sleep(500 * time.Millisecond)

	status = invertBit(status)
	grovepi.DigitalWrite(led.Pin, status)

}

// Flippin' bits
func invertBit(b byte) byte {
	return byte(int(b) ^ 1)
}
