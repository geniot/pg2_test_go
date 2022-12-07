package dev

type DesktopDeviceImpl struct {
}

func NewDesktopDevice() *DesktopDeviceImpl {
	return &DesktopDeviceImpl{}
}

func (device DesktopDeviceImpl) IsRumbleSupported() bool {
	return false
}

func (device DesktopDeviceImpl) Rumble() {
	println("Rumble is not supported on this device.")
}
