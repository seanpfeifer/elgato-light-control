package elgato

import (
	"net"
	"strings"
	"sync"

	"github.com/hashicorp/mdns"
)

const (
	MDNSService   = "_elg._tcp"
	NameRingLight = "Elgato Ring Light"
)

type Device struct {
	Name string
	IP   net.IP
	Port uint16
}

func FindDevices(searchStr string, iface net.Interface) ([]Device, error) {
	var devices []Device
	entriesCh := make(chan *mdns.ServiceEntry, 4)

	// Using a WaitGroup so we can make sure our goroutine finishes processing before the return happens.
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for entry := range entriesCh {
			if strings.Contains(entry.Info, searchStr) {
				devices = append(devices, Device{
					Name: entry.Host,
					IP:   entry.AddrV4,
					Port: uint16(entry.Port),
				})
			}
		}
	}()

	params := mdns.DefaultParams(MDNSService)
	params.Entries = entriesCh
	params.DisableIPv6 = true
	// This needs to be done on Windows 11, otherwise the query will fail to find any devices
	params.Interface = &iface
	err := mdns.Query(params)
	close(entriesCh)

	// Wait until the goroutine finishes writing to `devices`.
	wg.Wait()
	return devices, err
}
