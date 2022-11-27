package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func main() {

	var window *sdl.Window
	var font *ttf.Font
	var surface *sdl.Surface
	var text *sdl.Surface
	var err error
	var pressedKeysCodes = mapset.NewSet[sdl.Keycode]()
	var pngImage *sdl.Surface

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
		"pg2_test_go",
		int32(If(numVideoDisplay > 1, SECOND_SCREEN_X_OFFSET, sdl.WINDOWPOS_UNDEFINED)),
		sdl.WINDOWPOS_UNDEFINED,
		SCREEN_WIDTH, SCREEN_HEIGHT,
		sdl.WINDOW_SHOWN|sdl.WINDOW_BORDERLESS)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err = window.GetSurface()
	if err != nil {
		panic(err)
	}

	surface.FillRect(nil, sdl.MapRGB(surface.Format, 16, 16, 16))

	// Load the font for our text
	if font, err = ttf.OpenFont(FONT_PATH, FONT_SIZE); err != nil {
		panic(err)
	}
	defer font.Close()

	if pngImage, err = img.Load("media/pg2_back.png"); err != nil {
		panic(err)
	}
	defer pngImage.Free()

	//var colorKey, _ = pngImage.GetColorKey()
	//var surf, _ = sdl.CreateRGBSurface(colorKey, pngImage.W, pngImage.H, 16, 0, 0, 0, 0)
	//pngImage.Blit(nil, surf, nil)
	pngImage.Blit(nil, surface, &sdl.Rect{X: 90, Y: 50, W: pngImage.W, H: pngImage.H})
	//surf.SetColorKey(true, sdl.MapRGB(surface.Format, 255, 0, 255))
	//defer surf.Free()

	// Create a red text with the font
	if text, err = font.RenderUTF8Blended("Hello, World!", sdl.Color{R: 255, G: 0, B: 0, A: 255}); err != nil {
		return
	}
	defer text.Free()

	// Draw the text around the center of the window
	if err = text.Blit(nil, surface, &sdl.Rect{X: SCREEN_WIDTH/2 - text.W/2, Y: SCREEN_HEIGHT/2 - text.H/2, W: 0, H: 0}); err != nil {
		return
	}

	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.KeyboardEvent:
				keyCode := t.Keysym.Sym
				if t.State == sdl.PRESSED {
					pressedKeysCodes.Add(keyCode)
				} else { // if t.State == sdl.RELEASED {
					pressedKeysCodes.Remove(keyCode)
				}
				if pressedKeysCodes.Contains(sdl.K_q) ||
					(pressedKeysCodes.Contains(GCW_BUTTON_L1) && pressedKeysCodes.Contains(GCW_BUTTON_START)) {
					running = false
				}
				redraw()
			}
		}

		//sdl.Delay(16)
	}

}

func redraw() {
	//rect := sdl.Rect{0, 0, SCREEN_WIDTH, SCREEN_HEIGHT}
	//surface.FillRect(&rect, sdl.MapRGB(surface.Format, 16, 16, 16))
	//
	//var txt = strconv.Itoa(int(keyCode)) + ":" + string(keyCode)
	//text, err = font.RenderUTF8Blended(txt, sdl.Color{R: 255, G: 0, B: 0, A: 255})
	//text.Blit(nil, surface, &sdl.Rect{X: SCREEN_WIDTH/2 - text.W/2, Y: SCREEN_HEIGHT/2 - text.H/2, W: 0, H: 0})
	//
	//window.UpdateSurface()
}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}
