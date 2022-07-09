package main

import (
	"log"

	"github.com/seanpfeifer/elgato-light-control/elgato"
	"github.com/seanpfeifer/rigging/logging"
)

const (
	defaultAddress = "192.168.1.201"
)

func main() {
	lights, err := elgato.GetLightInfo(defaultAddress, elgato.DefaultPort)
	logging.FatalIfError(err, "getting light info")
	log.Printf("%+v", lights)

	lights.Toggle()
	err = elgato.UpdateLightOptions(defaultAddress, elgato.DefaultPort, lights)
	logging.FatalIfError(err, "toggling light")
}
