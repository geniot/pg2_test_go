package dev

import (
	"fmt"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/glb"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"strconv"
	"strings"
)

type HandheldDeviceImpl struct {
	haptic            *sdl.Haptic
	joystick          *sdl.Joystick
	isRumbleSupported bool
}

func (handhelpDevice HandheldDeviceImpl) GetJoystickAxis(axis int) int16 {
	return handhelpDevice.joystick.Axis(axis)
}

func (handhelpDevice HandheldDeviceImpl) Stop() {
	handhelpDevice.joystick.Close()
	handhelpDevice.haptic.Close()
	closeCommon()
}

func (handhelpDevice HandheldDeviceImpl) UpdateBatteryStatus() {
	dat, err := os.ReadFile("/sys/class/power_supply/usb/online")
	if err != nil {
		fmt.Println(err.Error())
	}
	isCharging := strings.TrimSpace(string(dat)) == "1"
	voltage, err := os.ReadFile("/sys/class/power_supply/battery/voltage_now")
	if err != nil {
		fmt.Println(err.Error())
	}
	pct, err := strconv.Atoi(strings.TrimSpace(string(voltage)))

	if isCharging {
		pct = ((pct - glb.MIN_VOLTAGE) - glb.USB_VOLTAGE) * 100 / (glb.MAX_VOLTAGE - glb.MIN_VOLTAGE)
	} else {
		pct = (pct - glb.MIN_VOLTAGE) * 100 / (glb.MAX_VOLTAGE - glb.MIN_VOLTAGE)
	}

	if pct > 100 {
		pct = 100
	}
	if pct < 0 {
		pct = 0
	}

	//voltage jumps a little but percentage cannot go up if we are not charging
	if !isCharging && pct > ctx.PowerInformation.Pct {
		pct = ctx.PowerInformation.Pct
	}
	//we cannot go down when charging
	if isCharging && pct < ctx.PowerInformation.Pct {
		pct = ctx.PowerInformation.Pct
	}

	ctx.PowerInformation.Pct = pct
	ctx.PowerInformation.IsCharging = isCharging
}

func (handhelpDevice HandheldDeviceImpl) UpdateDiskStatus() {
	updateDiskInfo("/usr/local/home", &ctx.DiskInfo1)
	updateDiskInfo("/media/sdcard/", &ctx.DiskInfo2)
}

/*
see: https://github.com/libsdl-org/SDL/issues/6744
*/
func (handhelpDevice HandheldDeviceImpl) UpdateVolume() {
	//cmd := exec.Command("amixer", "sget", "Master")
	//res, _ := cmd.Output()
	//lines := strings.Split(string(res), "\n")
	//lastLine := lines[len(lines)-2]
	//percentSplit := strings.Split(lastLine, "[")[1]
	//percentStr := strings.Split(percentSplit, "%")[0]
	//currentVolume, _ = strconv.Atoi(percentStr)
}

func (handhelpDevice HandheldDeviceImpl) ProcessKeyActions() {
	if ctx.PressedKeysCodes.Contains(glb.GCW_BUTTON_L1) &&
		ctx.PressedKeysCodes.Contains(glb.GCW_BUTTON_START) {
		ctx.Loop.Stop()
	}
	if ctx.PressedKeysCodes.Contains(glb.GCW_BUTTON_L1) &&
		ctx.PressedKeysCodes.Contains(glb.GCW_BUTTON_X) {
		if mix.Playing(-1) != 1 {
			ctx.Application.PlaySound()
		}
	}
	if ctx.PressedKeysCodes.Contains(glb.GCW_BUTTON_L2) &&
		ctx.PressedKeysCodes.Contains(glb.GCW_BUTTON_R2) {
		handhelpDevice.rumble()
	}
}

func (handhelpDevice HandheldDeviceImpl) GetWindowPosAndSize() (int32, int32, int32, int32) {
	return 0, 0, 320, 240
}

func (handhelpDevice HandheldDeviceImpl) GetWindowState() uint32 {
	return sdl.WINDOW_SHOWN | sdl.WINDOW_BORDERLESS
}

func (handhelpDevice HandheldDeviceImpl) IsRumbleSupported() bool {
	return handhelpDevice.isRumbleSupported
}

func (handhelpDevice HandheldDeviceImpl) rumble() {
	err := handhelpDevice.haptic.RumblePlay(0.33, 500)
	if err != nil {
		println(err.Error())
	}
}

func NewHandheldDevice() HandheldDeviceImpl {
	device := HandheldDeviceImpl{}
	device.init()
	return device
}

func (handhelpDevice *HandheldDeviceImpl) init() {
	initCommon()
	numHaptics, err := sdl.NumHaptics()
	if err != nil {
		panic(err)
	}
	if numHaptics > 0 {
		println("Haptics: " + strconv.Itoa(numHaptics))
		println(sdl.HapticName(0))
		handhelpDevice.haptic, err = sdl.HapticOpen(0)
		if err != nil {
			panic(err)
		}
		err = handhelpDevice.haptic.RumbleInit()
		if err != nil {
			panic(err)
		}
		handhelpDevice.isRumbleSupported, _ = handhelpDevice.haptic.RumbleSupported()
	}
	numJoysticks := sdl.NumJoysticks()
	if numJoysticks > 0 {
		println("Joysticks: " + strconv.Itoa(numJoysticks))
		println(sdl.JoystickNameForIndex(0))
		handhelpDevice.joystick = sdl.JoystickOpen(0)
	}
	sdl.JoystickEventState(sdl.ENABLE)
}
