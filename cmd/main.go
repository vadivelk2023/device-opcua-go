package main

import (
	"fmt"
	"github.com/vadivelk2023/device-opcua-go"

	"github.com/vadivelk2023/device-opcua-go/internal/driver"
)

const (
	serviceName string = "device-opcua"
)

func main() {
	sd := driver.NewProtocolDriver()
	startup.Bootstrap(serviceName, device_opcua.Version, sd)
	fmt.Println("Device OpcUA")
}