package main

import (
	"fmt"
	"github.com/xshoji/go-sample-box/enumtype/enums"
)

type Device struct {
	DeviceType enums.DeviceType
	Version    enums.Version
}

func main() {
	// Print
	fmt.Println("<< Print enum values >>")
	fmt.Println(enums.DeviceTypeAndroid)
	fmt.Println(enums.DeviceTypeIos)
	fmt.Println(enums.DeviceTypeWindows)
	fmt.Println(enums.Version0000)
	fmt.Println(enums.Version2017)
	fmt.Println(enums.Version2018)
	fmt.Println()

	fmt.Println("<< Print struct with enum values >>")
	device := Device{
		DeviceType: enums.DeviceTypeAndroid,
		Version:    enums.Version2018,
	}
	fmt.Println(device)
}
