package allthingstalk

import (
// "log"
// mqtt "git.eclipse.org/gitroot/paho/org.eclipse.paho.mqtt.golang.git"
)

func Connect(broker *Broker) Connection {

	uri := broker.Uri

	return Connection{
		Address: uri,
	}

}

func Setup(device *Device) {

}
