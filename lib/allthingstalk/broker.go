package allthingstalk

import (
	// mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var httpUri = "http://beta.smartliving.io"
var brokerUri = "http://broker.smartliving.io"

var httpClient *http.Client
var brokerClient Broker

func Connect() {
	httpClient = &http.Client{}
	log.Printf("Initialized HTTP client")
}

/**
 * Registers your IODevice to the platform
 */
func (device *Device) RegisterAsset(ioDevice *IODevice) {

	payload := RegisterPayload{
		Name:        ioDevice.Name,
		Description: ioDevice.Description,
		Type:        ioDevice.Type,
		Profile:     ioDevice.Profile,
		DeviceId:    device.DeviceId,
	}

	payloadJSON, _ := json.Marshal(payload)

	req, err := http.NewRequest("PUT", buildAssetUri(device, ioDevice), bytes.NewReader(payloadJSON))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Auth-ClientKey", device.ClientKey)
	req.Header.Set("Auth-ClientId", device.ClientId)

	_, err = httpClient.Do(req)

	if err != nil {
		log.Fatalf("Error registering asset: %s", err)
	}

	log.Printf("Successfuly registered IODevice %+s\n", ioDevice.Name)
}

func Subscribe() {

	log.Printf("Initialized Broker")
}

func (device *Device) Setup() {
	Connect()
	Subscribe()
}

func buildAssetUri(device *Device, iodev *IODevice) string {
	return fmt.Sprintf("%s/api/asset/%s%s", httpUri, device.DeviceId, iodev.Id)
}
