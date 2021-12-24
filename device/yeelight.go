package device

import (
	"github.com/danilarff86/miio-go/capability"
	"github.com/danilarff86/miio-go/common"
)

type Yeelight struct {
	Device
	*capability.Light
	*capability.Power
}

func NewYeelight(device Device) *Yeelight {
	dev := &Yeelight{
		Device: device,
		Power:  capability.NewPower(device, device.Outbound()),
		Light:  capability.NewLight(device, device.Outbound()),
	}
	go dev.refresh()
	return dev
}

func (p *Yeelight) refresh() {
	for range p.RefreshThrottle() {
		_ = p.Power.Update()
		_ = p.Light.Update()
	}

	common.Log.Debug("Device refresh closed.")
}
