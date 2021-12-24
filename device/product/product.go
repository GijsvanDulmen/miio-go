package product

import "fmt"

type Product uint16

const (
	PowerPlug Product = iota << 1
	Yeelight
	Unknown
)

func GetModel(modelName string) (Product, error) {
	switch modelName {
	case "chuangmi.plug.m1", "chuangmi.plug.hmi206":
		return PowerPlug, nil
	case "yeelink.light.color1":
		return Yeelight, nil
	default:
		return Unknown, fmt.Errorf("Unknown product for device type %s", modelName)
	}
}
