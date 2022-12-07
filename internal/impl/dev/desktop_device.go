package dev

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/mdl"
	"github.com/veandco/go-sdl2/sdl"
)

type DesktopDeviceImpl struct {
}

func (device DesktopDeviceImpl) GetWindowPosAndSize() (int32, int32, int32, int32) {
	return int32(ctx.Config.Get(mdl.WINDOW_XPOS_KEY)),
		int32(ctx.Config.Get(mdl.WINDOW_YPOS_KEY)),
		320, 240
	//int32(ctx.Config.Get(mdl.WINDOW_WIDTH_KEY)),
	//int32(ctx.Config.Get(mdl.WINDOW_HEIGHT_KEY))
}

func (device DesktopDeviceImpl) GetWindowState() uint32 {
	//return ctx.Config.Get(mdl.WINDOW_STATE_KEY)
	return sdl.WINDOW_SHOWN
}

func (device DesktopDeviceImpl) PlaySound() {
	//TODO implement me
	panic("implement me")
}

func NewDesktopDevice() DesktopDeviceImpl {
	device := DesktopDeviceImpl{}
	device.init()
	return device
}

func (device DesktopDeviceImpl) IsRumbleSupported() bool {
	return false
}

func (device DesktopDeviceImpl) Rumble() {
	println("Rumble is not supported on this device.")
}

func (device DesktopDeviceImpl) init() {
	initCommon()
}
