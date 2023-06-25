package hue

// State represents the current state of a light.
type State struct {
	On         bool      `json:"on"`        // Whether the light is on or off
	Brightness uint8     `json:"bri"`       // Brightness level (0-255)
	Hue        uint16    `json:"hue"`       // Hue value (0-65535)
	Saturation uint8     `json:"sat"`       // Saturation level (0-255)
	Effect     string    `json:"effect"`    // Currently active effect
	XY         []float64 `json:"xy"`        // XY color coordinates
	ColorTemp  uint16    `json:"ct"`        // Color temperature
	Alert      string    `json:"alert"`     // Alert state (none, select, lselect)
	ColorMode  string    `json:"colormode"` // Color mode (hs, xy, ct)
	Mode       string    `json:"mode"`      // Light mode (homeautomation, normal)
	Reachable  bool      `json:"reachable"` // Whether the light is reachable
}

// SWUpdate represents the software update information of a light.
type SWUpdate struct {
	State       string `json:"state"`       // Software update state (notupdatable, noupdates, downloading, readytoinstall, installing)
	LastInstall string `json:"lastinstall"` // Last installation date/time
}

// Capabilities represents the capabilities of a light.
type Capabilities struct {
	Certified bool `json:"certified"` // Whether the light is certified
	Control   struct {
		MinDimLevel    uint8          `json:"mindimlevel"`    // Minimum dimming level
		MaxLumen       uint16         `json:"maxlumen"`       // Maximum luminosity
		ColorGamutType string         `json:"colorgamuttype"` // Type of color gamut supported
		ColorGamut     [][]float64    `json:"colorgamut"`     // Color gamut values
		ColorTempRange map[string]int `json:"ct"`             // Range of color temperatures supported
	} `json:"control"` // Control capabilities
	Streaming struct {
		Renderer bool `json:"renderer"` // Whether the light can act as a renderer
		Proxy    bool `json:"proxy"`    // Whether the light can act as a proxy
	} `json:"streaming"` // Streaming capabilities
}

// Config represents the configuration of a light.
type Config struct {
	Archetype string `json:"archetype"` // Archetype type
	Function  string `json:"function"`  // Function type
	Direction string `json:"direction"` // Direction type
	Startup   struct {
		Mode       string `json:"mode"`       // Startup mode
		Configured bool   `json:"configured"` // Whether startup is configured
	} `json:"startup"` // Startup configuration
}

// Light represents a Hue light.
type Light struct {
	State        State        `json:"state"`            // Current state of the light
	SWUpdate     SWUpdate     `json:"swupdate"`         // Software update information
	Type         string       `json:"type"`             // Light type
	Name         string       `json:"name"`             // Light name
	ModelID      string       `json:"modelid"`          // Model ID
	Manufacturer string       `json:"manufacturername"` // Manufacturer name
	ProductName  string       `json:"productname"`      // Product name
	Capabilities Capabilities `json:"capabilities"`     // Capabilities of the light
	Config       Config       `json:"config"`           // Configuration of the light
	UniqueID     string       `json:"uniqueid"`         // Unique identifier
	SWVersion    string       `json:"swversion"`        // Software version
	SWConfigID   string       `json:"swconfigid"`       // Software configuration ID
	ProductID    string       `json:"productid"`        // Product ID
}

// LightState represents the desired state to be set for a light.
type LightState struct {
	On         *bool      `json:"on,omitempty"`        // Whether the light should be turned on or off
	Brightness *uint8     `json:"bri,omitempty"`       // Brightness level (0-255)
	Hue        *uint16    `json:"hue,omitempty"`       // Hue value (0-65535)
	Saturation *uint8     `json:"sat,omitempty"`       // Saturation level (0-255)
	Effect     *string    `json:"effect,omitempty"`    // Effect to be activated
	XY         *[]float64 `json:"xy,omitempty"`        // XY color coordinates
	ColorTemp  *uint16    `json:"ct,omitempty"`        // Color temperature
	Alert      *string    `json:"alert,omitempty"`     // Alert state (none, select, lselect)
	ColorMode  *string    `json:"colormode,omitempty"` // Color mode (hs, xy, ct)
	Mode       *string    `json:"mode,omitempty"`      // Light mode (homeautomation, normal)
	Reachable  *bool      `json:"reachable,omitempty"` // Whether the light is reachable
}
