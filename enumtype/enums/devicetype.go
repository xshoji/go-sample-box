package enums

//go:generate go run golang.org/x/tools/cmd/stringer -type=DeviceType
type DeviceType int

const (
	DeviceTypeNull DeviceType = iota
	DeviceTypeIos
	DeviceTypeAndroid
	DeviceTypeWindows
	DeviceTypeLinux
)
