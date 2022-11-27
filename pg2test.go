package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"strconv"
)

const (
	fontPath           = "media/pixelberry.ttf"
	fontSize           = 14
	width              = 320
	height             = 240
	secondScreenOffset = 2500 //used for testing/debugging
)

func main() {

	var window *sdl.Window
	var font *ttf.Font
	var surface *sdl.Surface
	var text *sdl.Surface
	var err error
	var pressedKeysCodes = mapset.NewSet[sdl.Keycode]()

	if err = ttf.Init(); err != nil {
		return
	}
	defer ttf.Quit()

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	numVideoDisplay, err := sdl.GetNumVideoDisplays()

	window, err = sdl.CreateWindow(
		"test",
		int32(If(numVideoDisplay > 1, secondScreenOffset, sdl.WINDOWPOS_UNDEFINED)),
		sdl.WINDOWPOS_UNDEFINED,
		width, height,
		sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(nil, 0xffffffff)

	// Load the font for our text
	if font, err = ttf.OpenFont(fontPath, fontSize); err != nil {
		panic(err)
	}
	defer font.Close()

	// Create a red text with the font
	if text, err = font.RenderUTF8Blended("Hello, World!", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
		return
	}
	defer text.Free()

	// Draw the text around the center of the window
	if err = text.Blit(nil, surface, &sdl.Rect{X: width/2 - text.W/2, Y: height/2 - text.H/2, W: 0, H: 0}); err != nil {
		return
	}

	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				if t.State == sdl.RELEASED {
					pressedKeysCodes.Remove(keyCode)
				} else if t.State == sdl.PRESSED {
					pressedKeysCodes.Add(keyCode)
				}

				rect := sdl.Rect{0, 0, width, height}
				surface.FillRect(&rect, 0xffffffff)
				var txt = strconv.Itoa(int(keyCode)) + ":" + string(keyCode)
				text, err = font.RenderUTF8Blended(txt, sdl.Color{R: 255, G: 0, B: 0, A: 255})
				text.Blit(nil, surface, &sdl.Rect{X: width/2 - text.W/2, Y: height/2 - text.H/2, W: 0, H: 0})
				window.UpdateSurface()

				if pressedKeysCodes.Contains(sdl.K_q) ||
					(pressedKeysCodes.Contains(GCW_BUTTON_L1) && pressedKeysCodes.Contains(GCW_BUTTON_START)) {
					running = false
				}
			}
		}

		//sdl.Delay(16)
	}

}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}
