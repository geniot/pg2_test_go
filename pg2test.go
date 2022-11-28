package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
)

var window *sdl.Window
var surface *sdl.Surface
var font *ttf.Font
var lastPressedKey sdl.Keycode
var pressedKeysCodes = mapset.NewSet[sdl.Keycode]()
var imageElements []ImageElement
var audioChunk *mix.Chunk
var joystick *sdl.Joystick
var haptic *sdl.Haptic

func main() {

	initAll()
	defer closeAll()

	redraw()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
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
				if pressedKeysCodes.Contains(sdl.K_q) ||
					(pressedKeysCodes.Contains(GCW_BUTTON_L1) && pressedKeysCodes.Contains(GCW_BUTTON_START)) {
					running = false
				}
				if pressedKeysCodes.Contains(GCW_BUTTON_L1) && pressedKeysCodes.Contains(GCW_BUTTON_X) && mix.Playing(-1) != 1 {
					_, err := audioChunk.Play(1, 0)
					if err != nil {
						panic(err)
					}
				}
				if pressedKeysCodes.Contains(GCW_BUTTON_L2) && pressedKeysCodes.Contains(GCW_BUTTON_R2) {
					//var err = joystick.Rumble(0, 0xFFFF, 3000)
					var err = haptic.RumblePlay(1.0, 3000)
					if err != nil {
						panic(err)
					}
				}
				redraw()
				break
			}
		}

		//sdl.Delay(16)
	}

}

func initAll() {
	err := ttf.Init()
	err = sdl.Init(sdl.INIT_JOYSTICK | sdl.INIT_AUDIO | sdl.INIT_VIDEO)
	haptic, err = sdl.HapticOpen(0)
	err = haptic.RumbleInit()
	err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	sdl.JoystickEventState(sdl.ENABLE)
	joystick = sdl.JoystickOpen(0)
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
		panic(err)
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

func redraw() {

	var err = surface.FillRect(nil, sdl.MapRGB(surface.Format, 16, 16, 16))
	if err != nil {
		panic(err)
	}

	for i, imageElement := range imageElements {
		if imageElement.displayOnPress == sdl.K_UNKNOWN ||
			pressedKeysCodes.Contains(imageElement.displayOnPress) {
			var err = imageElement.surface.Blit(
				nil,
				surface,
				&sdl.Rect{X: imageElement.offsetX, Y: imageElement.offsetY, W: imageElements[i].surface.W, H: imageElements[i].surface.H})
			if err != nil {
				panic(err)
			}
		}
	}

	drawText(MSG_0, 10, 180, 255, 255, 0)
	drawText(MSG_1, 10, 190, 255, 255, 0)
	//last detected key
	drawText(MSG_2, 10, 160, 0, 255, 255)
	drawText(fmt.Sprintf("%d [0x%04X]", lastPressedKey, lastPressedKey), 110, 160, 255, 255, 255)

	//drawText(MSG_3, 10, 200, 255, 255, 0)
	drawText(MSG_4, 197, 20, 255, 255, 255)
	drawText(MSG_5, 10, 200, 255, 255, 0)

	err = window.UpdateSurface()
	if err != nil {
		panic(err)
	}
}

func drawText(txt string, x int32, y int32, fR uint8, fG uint8, fB uint8) {
	var text, err = font.RenderUTF8Blended(txt, sdl.Color{R: fR, G: fG, B: fB, A: 255})
	if err != nil {
		return
	}
	defer text.Free()
	err = text.Blit(nil, surface, &sdl.Rect{X: x, Y: y, W: 0, H: 0})
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
