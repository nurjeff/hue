package hue_controller

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

// Setup initializes and sets up a connection to the Hue Bridge.
// It returns a pointer to a HueBridge struct and an error if there was a problem during setup.
// Make sure your bridge is reachable via your local network and a HUE_USERNAME environment variable is set.
// Refer here: https://developers.meethue.com/develop/get-started-2/
func Setup() (*HueBridge, error) {
	h := HueBridge{}
	if err := h.find(); err != nil {
		return nil, err
	}
	return &h, nil
}

// HueBridge represents a Hue Bridge with its properties and functionalities.
type HueBridge struct {
	IPs        []net.IP        // IP addresses of the Hue Bridge
	Instance   string          // Instance ID of the Hue Bridge
	Port       int             // Port number of the Hue Bridge
	Text       []string        // Textual descriptions of the Hue Bridge
	Controller *mDNSController // Controller for mDNS service discovery
	Lights     []Light         // List of lights connected to the Hue Bridge
	client     *http.Client    // HTTP client for making API requests
	Username   string          // Username for authentication with the Hue Bridge
}

// find discovers the Hue Bridge on the network and initializes the HueBridge struct.
// It retrieves the necessary information such as IP addresses, instance ID, port, and connected lights.
func (h *HueBridge) find() error {
	username := os.Getenv("HUE_USERNAME")
	if username == "" {
		panic(errors.New("HUE_USERNAME environment variable not set"))
	}
	h.Username = username
	h.Controller = &mDNSController{}
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	h.client = &http.Client{Transport: transport, Timeout: 5 * time.Second}
	if err := h.Controller.Init(); err != nil {
		return err
	}
	hb, err := h.Controller.SearchHue()
	if err != nil {
		return err
	}
	h.IPs = hb.IPs
	h.Instance = hb.Instance
	h.Port = hb.Port
	h.Text = hb.Text
	if err := h.FetchLights(); err != nil {
		return err
	}

	if len(h.IPs) < 1 {
		return errors.New("no IPs found")
	}
	return nil
}

// FetchLights retrieves the list of lights connected to the Hue Bridge.
func (b *HueBridge) FetchLights() error {
	resp, err := b.makeAPIRequest(http.MethodGet, "lights", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var lights map[string]Light
	err = json.NewDecoder(resp.Body).Decode(&lights)
	if err != nil {
		return err
	}

	b.Lights = make([]Light, 0, len(lights))
	for _, light := range lights {
		b.Lights = append(b.Lights, light)
	}

	return nil
}

// makeAPIRequest makes an API request to the Hue Bridge with the specified method, path, and request body.
// It returns the HTTP response and an error if there was a problem making the request.
func (b *HueBridge) makeAPIRequest(method, path string, body []byte) (*http.Response, error) {
	url := fmt.Sprintf("https://%s:%d/api/%s/%s", b.IPs[0], b.Port, b.Username, path)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	return b.client.Do(req)
}

// createJSONPayload creates a JSON payload from the given value.
// It returns the JSON payload as a byte slice and an error if there was a problem marshaling the value.
func createJSONPayload(v interface{}) ([]byte, error) {
	payload, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return payload, nil
}

// ToggleLight toggles the state of a light with the specified ID on or off.
// It takes an ID of the light (0-based indexing, the Hue Bridge uses 1-based internally) and a boolean value indicating whether to turn the light on or off.
// It returns the HTTP response and an error if there was a problem making the API request.
func (b *HueBridge) ToggleLight(id int, on bool) (*http.Response, error) {
	id++
	if id > len(b.Lights) || id < 1 {
		return nil, errors.New("ID not valid")
	}

	path := fmt.Sprintf("lights/%d/state", id)
	state := make(map[string]interface{})
	state["on"] = on
	jsonData, err := createJSONPayload(state)
	if err != nil {
		return nil, err
	}

	resp, err := b.makeAPIRequest(http.MethodPut, path, jsonData)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// ChangeLight changes the state of a light with the specified ID to the new state.
// It takes an ID of the light (0-based indexing, the Hue Bridge uses 1-based internally)  and the new state to be set.
// It returns the HTTP response and an error if there was a problem making the API request.
func (b *HueBridge) ChangeLight(id int, newState LightState) (*http.Response, error) {
	id++
	if id > len(b.Lights) || id < 1 {
		return nil, errors.New("ID not valid")
	}

	path := fmt.Sprintf("lights/%d/state", id)
	jsonData, err := json.Marshal(newState)
	if err != nil {
		return nil, err
	}

	resp, err := b.makeAPIRequest(http.MethodPut, path, jsonData)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
