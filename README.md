## Small go plugin to control a Phillips Hue Bridge
----
Requires your Hue REST-Api username to be set as env variable.

```bash
HUE_USERNAME=<username>
``` 

```go
func main() {
	bridge, _ := hue.Setup()

	bridge.ChangeLight(2, hue.LightState{On: hue.VPtrs(true), Brightness: hue.VPtrs[uint8](255), Hue: hue.VPtrs[uint16](40000)})
}

```

i guess thats it