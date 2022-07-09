package main

import (
	"log"

	"github.com/seanpfeifer/elgato-light-control/elgato"
	"github.com/seanpfeifer/rigging/logging"
)

func main() {
	devices, err := elgato.FindDevices(elgato.NameRingLight)
	logging.FatalIfError(err, "finding lights")
	if len(devices) < 1 {
		return
	}

	address := devices[0].IP.String()
	port := devices[0].Port

	lights, err := elgato.GetLightInfo(address, port)
	logging.FatalIfError(err, "getting light info")
	log.Printf("%+v", lights)

	lights.Toggle()
	err = elgato.UpdateLightOptions(address, port, lights)
	logging.FatalIfError(err, "toggling lights")
}
