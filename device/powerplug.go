package device

import (
	"github.com/danilarff86/miio-go/capability"
	"github.com/danilarff86/miio-go/common"
)

type PowerPlug struct {
	Device
	*capability.Power
}

func NewPowerPlug(device Device) *PowerPlug {
	dev := &PowerPlug{
		Device: device,
		Power:  capability.NewPower(device, device.Outbound()),
	}
	go dev.refresh()
	return dev
}

func (p *PowerPlug) refresh() {
	for range p.RefreshThrottle() {
		_ = p.Power.Update()
	}

	common.Log.Debug("Device refresh closed.")
}
