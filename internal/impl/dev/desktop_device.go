package dev

type DesktopDeviceImpl struct {
}

func (device DesktopDeviceImpl) PlaySound() {
	//TODO implement me
	panic("implement me")
}

func NewDesktopDevice() DesktopDeviceImpl {
	return DesktopDeviceImpl{}
}

func (device DesktopDeviceImpl) IsRumbleSupported() bool {
	return false
}

func (device DesktopDeviceImpl) Rumble() {
	println("Rumble is not supported on this device.")
}
