## Small go plugin to control a Phillips Hue Bridge
----
Requires your Hue REST-Api username to be set as env variable.

```bash
HUE_USERNAME=<username>
``` 

```go
func main() {
	h := bridge.HueBridge{}
	h.Find()
    
	h.ChangeLight(2, bridge.LightState{On: bridge.VPtrs(true), Brightness: bridge.VPtrs[uint8](255), Hue: bridge.VPtrs[uint16](40000)})
}

```

i guess thats it