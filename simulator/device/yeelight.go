package device

import (
	"github.com/GijsvanDulmen/miio-go/simulator/capability"
)

type SimulatedYeelight struct {
	*BaseDevice
}

func NewSimulatedYeelight(baseDevice *BaseDevice) *SimulatedYeelight {
	baseDevice.AddCapability(&capability.Info{
		Model: "yeelink.light.color1",
	})
	baseDevice.AddCapability(&capability.Power{})
	baseDevice.AddCapability(&capability.Light{})

	return &SimulatedYeelight{
		BaseDevice: baseDevice,
	}
}
