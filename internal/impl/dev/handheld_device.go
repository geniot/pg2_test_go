package dev

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/mdl"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type HandheldDeviceImpl struct {
	haptic            *sdl.Haptic
	joystick          *sdl.Joystick
	isRumbleSupported bool
}

func (device HandheldDeviceImpl) ProcessKeyActions() {
	if ctx.PressedKeysCodes.Contains(mdl.GCW_BUTTON_L1) &&
		ctx.PressedKeysCodes.Contains(mdl.GCW_BUTTON_START) {
		ctx.Loop.Stop()
	}
	if ctx.PressedKeysCodes.Contains(mdl.GCW_BUTTON_L1) &&
		ctx.PressedKeysCodes.Contains(mdl.GCW_BUTTON_X) {
		if mix.Playing(-1) != 1 {
			ctx.Application.PlaySound()
		}
	}
	if ctx.PressedKeysCodes.Contains(mdl.GCW_BUTTON_L2) &&
		ctx.PressedKeysCodes.Contains(mdl.GCW_BUTTON_R2) {
		device.rumble()
	}
}

func (device HandheldDeviceImpl) GetWindowPosAndSize() (int32, int32, int32, int32) {
	return 0, 0, 320, 240
}

func (h HandheldDeviceImpl) GetWindowState() uint32 {
	return sdl.WINDOW_SHOWN | sdl.WINDOW_BORDERLESS
}

func (h HandheldDeviceImpl) IsRumbleSupported() bool {
	return h.isRumbleSupported
}

func (h HandheldDeviceImpl) rumble() {
	//TODO implement me
	panic("implement me")
}

func NewHandheldDevice() HandheldDeviceImpl {
	device := HandheldDeviceImpl{}
	device.init()
	return device
}

func (device HandheldDeviceImpl) init() {
	initCommon()
	numHaptics, err := sdl.NumHaptics()
	if err != nil {
		panic(err)
	}
	if numHaptics > 0 {
		println("Haptics: " + strconv.Itoa(numHaptics))
		println(sdl.HapticName(0))
		device.haptic, err = sdl.HapticOpen(0)
		if err != nil {
			panic(err)
		}
		err = device.haptic.RumbleInit()
		if err != nil {
			panic(err)
		}
		device.isRumbleSupported, _ = device.haptic.RumbleSupported()
	}
	numJoysticks := sdl.NumJoysticks()
	if numJoysticks > 0 {
		println("Joysticks: " + strconv.Itoa(numJoysticks))
		println(sdl.JoystickNameForIndex(0))
		device.joystick = sdl.JoystickOpen(0)
	}
	sdl.JoystickEventState(sdl.ENABLE)
}
