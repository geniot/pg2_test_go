package dev

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/glb"
	"github.com/itchyny/volume-go"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
)

type DesktopDeviceImpl struct {
}

func (desktopDevice DesktopDeviceImpl) GetJoystickAxis(axis int) int16 {
	return 0 //no joystick on desktop
}

func (desktopDevice DesktopDeviceImpl) Stop() {
	closeCommon()
}

func (desktopDevice DesktopDeviceImpl) UpdateBatteryStatus() {

}

func (desktopDevice DesktopDeviceImpl) UpdateDiskStatus() {
	if runtime.GOOS == "windows" {
		updateDiskInfo("C:\\", &ctx.DiskInfo1)
		//can be uncommented for debugging
		updateDiskInfo("C:\\", &ctx.DiskInfo2)
	} else {
		updateDiskInfo("/", &ctx.DiskInfo1)
		//can be uncommented for debugging
		updateDiskInfo("/", &ctx.DiskInfo2)
	}
}

func (desktopDevice DesktopDeviceImpl) UpdateVolume() {
	newVolume, err := volume.GetVolume()
	if err != nil {
		println(err.Error())
	} else {
		ctx.CurrentVolume = newVolume
	}
}

func (desktopDevice DesktopDeviceImpl) ProcessKeyActions() {
	if ctx.PressedKeysCodes.Contains(sdl.K_q) {
		ctx.Loop.Stop()
	}
	if ctx.PressedKeysCodes.Contains(sdl.K_s) {
		if mix.Playing(-1) != 1 {
			ctx.Application.PlaySound()
		}
	}
}

func (desktopDevice DesktopDeviceImpl) GetWindowPosAndSize() (int32, int32, int32, int32) {
	return int32(ctx.Config.Get(glb.WINDOW_XPOS_KEY)),
		int32(ctx.Config.Get(glb.WINDOW_YPOS_KEY)),
		320, 240
	//int32(ctx.Config.Get(mdl.WINDOW_WIDTH_KEY)),
	//int32(ctx.Config.Get(mdl.WINDOW_HEIGHT_KEY))
}

func (desktopDevice DesktopDeviceImpl) GetWindowState() uint32 {
	//return ctx.Config.Get(mdl.WINDOW_STATE_KEY)
	return sdl.WINDOW_SHOWN
}

func NewDesktopDevice() DesktopDeviceImpl {
	device := DesktopDeviceImpl{}
	device.init()
	return device
}

func (desktopDevice DesktopDeviceImpl) init() {
	initCommon()
}
