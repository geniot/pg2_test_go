package main

import (
	"geniot.com/geniot/pg2_test_go/bak"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var window *sdl.Window
var surface *sdl.Surface

// var renderer *sdl.Renderer
var font *ttf.Font
var lastPressedKey sdl.Keycode
var pressedKeysCodes = mapset.NewSet[sdl.Keycode]()
var imageElements []bak.ImageElement
var joystickImageElements []bak.ImageElement
var batteryImageElements []bak.ImageElement
var diskImageElements []bak.ImageElement
var volumeImageElements []bak.ImageElement
var audioChunk *mix.Chunk
var joystick *sdl.Joystick
var haptic *sdl.Haptic
var isRumbleSupported bool
var keyNames map[sdl.Keycode]string
var running bool
var powerInfo bak.PowerInfo
var diskInfos [2]bak.DiskInfo
var currentVolume int

func main() {

	bak.initAll()
	defer bak.closeAll()

	bak.render()

	running = true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {

			case *sdl.QuitEvent:
				running = false
				break

			case *sdl.KeyboardEvent:
				if t.Repeat > 0 {
					break
				}
				lastPressedKey = t.Keysym.Sym
				if t.State == sdl.PRESSED {
					pressedKeysCodes.Add(lastPressedKey)
				} else { // if t.State == sdl.RELEASED {
					pressedKeysCodes.Remove(lastPressedKey)
				}
				break
			}
		}

		bak.render()
		sdl.Delay(1000 / 60)
	}

}