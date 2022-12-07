package dev

import (
	"geniot.com/geniot/pg2_test_go/internal/api"
	"runtime"
)

func NewDevice() api.IDevice {
	if runtime.GOOS == "windows" {
		return NewDesktopDevice()
	} else {
		return NewHandheldDevice()
	}
}
