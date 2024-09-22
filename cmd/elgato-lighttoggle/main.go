package main

import (
	"log"
	"net"

	"github.com/seanpfeifer/elgato-light-control/elgato"
	"github.com/seanpfeifer/rigging/logging"
)

func main() {
	ifaces, err := net.Interfaces()
	logging.FatalIfError(err, "getting interfaces")
	// I'm just going to assume the interface I need is the first one
	devices, err := elgato.FindDevices(elgato.NameRingLight, ifaces[0])
	logging.FatalIfError(err, "finding lights")
	if len(devices) < 1 {
		return
	}

	// https://github.com/hashicorp/mdns/issues/80

	address := devices[0].IP.String()
	port := devices[0].Port

	lights, err := elgato.GetLightInfo(address, port)
	logging.FatalIfError(err, "getting light info")
	log.Printf("%+v", lights)

	lights.Toggle()
	err = elgato.UpdateLightOptions(address, port, lights)
	logging.FatalIfError(err, "toggling lights")
}
