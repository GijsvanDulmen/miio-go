package device

import (
	"time"

	"github.com/GijsvanDulmen/miio-go/common"
	"github.com/GijsvanDulmen/miio-go/device/product"
	"github.com/GijsvanDulmen/miio-go/protocol/packet"
	"github.com/GijsvanDulmen/miio-go/protocol/transport"
)

type Device interface {
	common.Device

	Handle(*packet.Packet) error
	Close() error
	Seen() time.Time
	Provisional() bool
	SetProvisional(bool)
	GetProduct() (product.Product, error)
	Discover() error
	RefreshThrottle() <-chan struct{}
	Outbound() transport.Outbound
}
