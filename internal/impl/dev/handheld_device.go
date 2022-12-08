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

func (device HandheldDeviceImpl) GetJoystickAxis(axis int) int16 {
	return device.joystick.Axis(axis)
}

func (device HandheldDeviceImpl) Stop() {
	closeCommon()
	device.joystick.Close()
	device.haptic.Close()
}

func (device HandheldDeviceImpl) UpdateBatteryStatus() {
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

func (device HandheldDeviceImpl) UpdateDiskStatus() {
	updateDiskInfo("/usr/local/home", &ctx.DiskInfo1)
	updateDiskInfo("/media/sdcard/", &ctx.DiskInfo2)
}

/*
see: https://github.com/libsdl-org/SDL/issues/6744
*/
func (device HandheldDeviceImpl) UpdateVolume() {
	//cmd := exec.Command("amixer", "sget", "Master")
	//res, _ := cmd.Output()
	//lines := strings.Split(string(res), "\n")
	//lastLine := lines[len(lines)-2]
	//percentSplit := strings.Split(lastLine, "[")[1]
	//percentStr := strings.Split(percentSplit, "%")[0]
	//currentVolume, _ = strconv.Atoi(percentStr)
}

func (device HandheldDeviceImpl) ProcessKeyActions() {
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
