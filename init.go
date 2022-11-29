package main

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"strconv"
	"fmt"
)

func initAll() {
	err := ttf.Init()
	if err != nil {
		panic(err)
	}
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	numHaptics, err := sdl.NumHaptics()
	if err != nil {
		panic(err)
	}
	if numHaptics > 0 {
		println("Haptics: " + strconv.Itoa(numHaptics))
		println(sdl.HapticName(0))
		haptic, err = sdl.HapticOpen(0)
		if err != nil {
			panic(err)
		}
		err = haptic.RumbleInit()
		if err != nil {
			panic(err)
		}
		isRumbleSupported, _ = haptic.RumbleSupported()
	}
	err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		println(err)
	}

	numJoysticks := sdl.NumJoysticks()
	if numJoysticks > 0 {
		println("Joysticks: " + strconv.Itoa(numJoysticks))
		println(sdl.JoystickNameForIndex(0))
		joystick = sdl.JoystickOpen(0)
	}

	sdl.JoystickEventState(sdl.ENABLE)

	//_, err = sdl.ShowCursor(0)
	numVideoDisplay, err := sdl.GetNumVideoDisplays()
	window, err = sdl.CreateWindow(
		"pg2_test_go",
		int32(If(numVideoDisplay > 1, SECOND_SCREEN_X_OFFSET, sdl.WINDOWPOS_UNDEFINED)),
		sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH, SCREEN_HEIGHT,
		sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS)
	surface, err = window.GetSurface()
	font, err = ttf.OpenFont(FONT_PATH, FONT_SIZE)

	initImageElements()
	data, err := os.ReadFile("media/tone.wav")
	audioChunk, err = mix.QuickLoadWAV(data)

	if err != nil {
		fmt.Println(err)
	}
}

func closeAll() {
	err := window.Destroy()
	closeImageElements()
	audioChunk.Free()
	joystick.Close()
	haptic.Close()
	font.Close()
	ttf.Quit()
	sdl.Quit()
	mix.CloseAudio()

	if err != nil {
		panic(err)
	}
}
