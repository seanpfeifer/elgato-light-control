package elgato

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	DefaultPort     = 9123
	LightEndpoint   = "http://%s:%d/elgato/lights"
	ContentTypeJSON = "application/json"
)

type Lights struct {
	NumberOfLights uint8          `json:"numberOfLights"`
	Lights         []LightOptions `json:"lights"`
}

func (l *Lights) Toggle() {
	for i := range l.Lights {
		l.Lights[i].Toggle()
	}
}

type LightOptions struct {
	On          uint8  `json:"on"`
	Brightness  uint8  `json:"brightness"`
	Temperature uint16 `json:"temperature"`
}

func (lo *LightOptions) Toggle() {
	// If a light is 1, turn it to zero. If zero, turn to 1
	lo.On = 1 - lo.On
}

func GetLightInfo(address string, port uint16) (*Lights, error) {
	url := fmt.Sprintf(LightEndpoint, address, port)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var lights Lights
	err = json.Unmarshal(b, &lights)
	return &lights, err
}

func UpdateLightOptions(address string, port uint16, lights *Lights) error {
	url := fmt.Sprintf(LightEndpoint, address, port)

	buf, err := json.Marshal(lights)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", ContentTypeJSON)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}
