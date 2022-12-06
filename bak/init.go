package bak

import (
	"fmt"
	"geniot.com/geniot/pg2_test_go/bak"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"runtime"
	"strconv"
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
		bak.haptic, err = sdl.HapticOpen(0)
		if err != nil {
			panic(err)
		}
		err = bak.haptic.RumbleInit()
		if err != nil {
			panic(err)
		}
		bak.isRumbleSupported, _ = bak.haptic.RumbleSupported()
	}
	err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		println(err.Error())
	}

	numJoysticks := sdl.NumJoysticks()
	if numJoysticks > 0 {
		println("Joysticks: " + strconv.Itoa(numJoysticks))
		println(sdl.JoystickNameForIndex(0))
		bak.joystick = sdl.JoystickOpen(0)
	}

	sdl.JoystickEventState(sdl.ENABLE)

	//_, err = sdl.ShowCursor(0)
	//numVideoDisplay, err := sdl.GetNumVideoDisplays()
	var windowProps uint32 = sdl.WINDOW_SHOWN | sdl.WINDOW_BORDERLESS
	if runtime.GOOS == "windows" {
		windowProps = sdl.WINDOW_SHOWN
	}

	bak.window, err = sdl.CreateWindow(
		"pg2_test_go",
		//int32(If(numVideoDisplay > 1, SECOND_SCREEN_X_OFFSET, sdl.WINDOWPOS_UNDEFINED)),
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH, SCREEN_HEIGHT,
		windowProps)

	//renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	bak.surface, err = bak.window.GetSurface()

	bak.font, err = ttf.OpenFont(FONT_PATH, FONT_SIZE)

	initArrays()

	loadImages(bak.imageElements)
	loadImages(bak.joystickImageElements)
	loadImages(bak.batteryImageElements)
	loadImages(bak.diskImageElements)
	loadImages(bak.volumeImageElements)

	data, err := os.ReadFile("media/tone.wav")
	bak.audioChunk, err = mix.QuickLoadWAV(data)

	if err != nil {
		fmt.Println(err.Error())
	}

	//todo: https://stackoverflow.com/questions/37135193/how-to-set-default-values-in-go-structs
	bak.powerInfo = PowerInfo{100, false}
	bak.diskInfos[0] = DiskInfo{false, "", ""}
	bak.diskInfos[1] = DiskInfo{false, "", ""}

	go updateBatteryStatus()
	go updateDiskStatus()
	go updateVolume()

	c := cron.New()
	_, err = c.AddFunc("@every 1s", updateBatteryStatus)
	_, err = c.AddFunc("@every 1s", updateDiskStatus)
	_, err = c.AddFunc("@every 1s", updateVolume)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Start()
}

func closeAll() {
	err := bak.window.Destroy()
	//err = renderer.Destroy()

	freeImageElements(bak.imageElements)
	freeImageElements(bak.joystickImageElements)
	freeImageElements(bak.batteryImageElements)
	freeImageElements(bak.diskImageElements)
	freeImageElements(bak.volumeImageElements)

	bak.audioChunk.Free()
	bak.joystick.Close()
	bak.haptic.Close()
	bak.font.Close()
	ttf.Quit()
	sdl.Quit()
	mix.CloseAudio()

	if err != nil {
		panic(err)
	}
}
