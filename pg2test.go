package main

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var window *sdl.Window
var surface *sdl.Surface
var font *ttf.Font
var pressedKeysCodes = mapset.NewSet[sdl.Keycode]()
var imageElements []ImageElement

func main() {

	initAll()
	defer closeAll()

	redraw()

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

func initAll() {
	err := ttf.Init()
	err = sdl.Init(sdl.INIT_EVERYTHING)
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
	if err != nil {
		panic(err)
	}
}

func closeAll() {
	closeImageElements()
	font.Close()
	ttf.Quit()
	sdl.Quit()
	err := window.Destroy()

	if err != nil {
		panic(err)
	}
}

func redraw() {

	var err = surface.FillRect(nil, sdl.MapRGB(surface.Format, 16, 16, 16))
	if err != nil {
		panic(err)
	}

	for i, imageElement := range imageElements {
		imageElements[i].surface.Blit(
			nil,
			surface,
			&sdl.Rect{X: imageElement.offsetX, Y: imageElement.offsetY, W: imageElements[i].surface.W, H: imageElements[i].surface.H})
	}
	//var colorKey, _ = pngImage.GetColorKey()
	//var surf, _ = sdl.CreateRGBSurface(colorKey, pngImage.W, pngImage.H, 16, 0, 0, 0, 0)
	//pngImage.Blit(nil, surf, nil)
	//pngImage.Blit(nil, surface, &sdl.Rect{X: 90, Y: 50, W: pngImage.W, H: pngImage.H})
	//surf.SetColorKey(true, sdl.MapRGB(surface.Format, 255, 0, 255))
	//defer surf.Free()

	// Create a red text with the font
	//text, err := font.RenderUTF8Blended("Hello, World!", sdl.Color{R: 255, G: 0, B: 0, A: 255})
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Draw the text around the center of the window
	//if err := text.Blit(nil, surface, &sdl.Rect{X: SCREEN_WIDTH/2 - text.W/2, Y: SCREEN_HEIGHT/2 - text.H/2, W: 0, H: 0}); err != nil {
	//	panic(err)
	//}
	//rect := sdl.Rect{0, 0, SCREEN_WIDTH, SCREEN_HEIGHT}
	//surface.FillRect(&rect, sdl.MapRGB(surface.Format, 16, 16, 16))
	//
	//var txt = strconv.Itoa(int(keyCode)) + ":" + string(keyCode)
	//text, err = font.RenderUTF8Blended(txt, sdl.Color{R: 255, G: 0, B: 0, A: 255})
	//text.Blit(nil, surface, &sdl.Rect{X: SCREEN_WIDTH/2 - text.W/2, Y: SCREEN_HEIGHT/2 - text.H/2, W: 0, H: 0})
	//
	err = window.UpdateSurface()
	if err != nil {
		panic(err)
	}
}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}
