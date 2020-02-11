package enums

//go:generate stringer -type=DeviceType
type DeviceType int

const (
	DeviceTypeNull DeviceType = iota
	DeviceTypeIos
	DeviceTypeAndroid
	DeviceTypeWindows
	DeviceTypeLinux
)
