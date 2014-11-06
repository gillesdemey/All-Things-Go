package allthingstalk

import (
	"bytes"
	"encoding/json"
	io "github.com/gillesdemey/All-Things-Go/lib/allthingstalk/io"
	"log"
	"net/http"
)

var httpClient *http.Client = &http.Client{}
var httpUri = "http://beta.smartliving.io"

// Registers your IODevice to the platform
func (device *Device) RegisterAsset(ioDevice *io.IODevice) {

	json, _ := json.Marshal(io.Registration{
		Name:        ioDevice.Name,
		Description: ioDevice.Description,
		Type:        ioDevice.Type,
		Profile:     ioDevice.Profile,
		DeviceId:    device.DeviceId,
	})

	req, err := http.NewRequest(
		"PUT",
		buildAssetUri(device, ioDevice),
		bytes.NewReader(json))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Auth-ClientKey", device.ClientKey)
	req.Header.Set("Auth-ClientId", device.ClientId)

	_, err = httpClient.Do(req)

	if err != nil {
		log.Fatalf("Error registering asset: %s\n", err)
	}

	log.Printf("Successfuly registered IODevice: %+s\n", ioDevice.Name)
}
