package dev

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/mdl"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type DesktopDeviceImpl struct {
}

func (device DesktopDeviceImpl) Stop() {
	closeCommon()
}

func (device DesktopDeviceImpl) UpdateBatteryStatus() {

}

func (device DesktopDeviceImpl) UpdateDiskStatus() {

}

func (device DesktopDeviceImpl) UpdateVolume() {

}

func (device DesktopDeviceImpl) ProcessKeyActions() {
	if ctx.PressedKeysCodes.Contains(sdl.K_q) {
		ctx.Loop.Stop()
	}
	if ctx.PressedKeysCodes.Contains(sdl.K_s) {
		if mix.Playing(-1) != 1 {
			ctx.Application.PlaySound()
		}
	}
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

func NewDesktopDevice() DesktopDeviceImpl {
	device := DesktopDeviceImpl{}
	device.init()
	return device
}

func (device DesktopDeviceImpl) init() {
	initCommon()
}
