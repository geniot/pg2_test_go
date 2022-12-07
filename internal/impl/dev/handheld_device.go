package dev

type HandheldDeviceImpl struct {
}

func (h HandheldDeviceImpl) IsRumbleSupported() bool {
	return false
}

func (h HandheldDeviceImpl) Rumble() {
	//TODO implement me
	panic("implement me")
}

func (h HandheldDeviceImpl) PlaySound() {
	//TODO implement me
	panic("implement me")
}

func NewHandheldDevice() HandheldDeviceImpl {
	return HandheldDeviceImpl{}
}
