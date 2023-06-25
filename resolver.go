package hue_controller

import (
	"context"
	"errors"

	"github.com/grandcat/zeroconf"
)

type mDNSController struct {
	Resolver    *zeroconf.Resolver
	BridgeEntry *zeroconf.ServiceEntry
}

func (m *mDNSController) SearchHue() (*HueBridge, error) {
	entries := make(chan *zeroconf.ServiceEntry)
	b := &HueBridge{}
	if err := m.Resolver.Browse(context.TODO(), "_hue._tcp", "local.", entries); err != nil {
		return nil, err
	}
	for entry := range entries {
		m.BridgeEntry = entry
		b = &HueBridge{IPs: entry.AddrIPv4, Instance: entry.Instance, Port: entry.Port, Text: entry.Text}
		return b, nil
	}
	return nil, errors.New("not found")
}

func (m *mDNSController) Init() error {
	r, err := zeroconf.NewResolver(nil)
	if err != nil {
		return err
	}
	m.Resolver = r
	return nil
}
