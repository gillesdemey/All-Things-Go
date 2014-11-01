package allthingstalk

import (
	"net"
)

func GetInterfaces() ([]net.Interface, error) {

	ifaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	return ifaces, nil
}
