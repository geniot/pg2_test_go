package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func render() {
	processKeys()
	redraw()
}

func processKeys() {
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
	if isRumbleSupported {
		if pressedKeysCodes.Contains(GCW_BUTTON_L2) && pressedKeysCodes.Contains(GCW_BUTTON_R2) {
			var err = haptic.RumblePlay(0.33, 500)
			if err != nil {
				panic(err)
			}
		}
	}
}

func redraw() {

	var err = surface.FillRect(nil, sdl.MapRGB(surface.Format, 16, 16, 16))
	if err != nil {
		panic(err)
	}

	for _, imgEl := range imageElements {
		if imgEl.displayOnPress == sdl.K_UNKNOWN ||
			pressedKeysCodes.Contains(imgEl.displayOnPress) {
			var err = imgEl.surface.Blit(
				nil,
				surface,
				&sdl.Rect{X: imgEl.offsetX, Y: imgEl.offsetY, W: imgEl.surface.W, H: imgEl.surface.H})
			if err != nil {
				panic(err)
			}
		}
	}

	drawMessages()
	drawJoystick()

	err = window.UpdateSurface()
	if err != nil {
		panic(err)
	}
}

func drawJoystick() {
	var x = int32(joystick.Axis(0) / 5461)
	var y = int32(joystick.Axis(1) / 5461)
	var jImgEl = joystickImageElements[0]
	if x != 0 || y != 0 {
		jImgEl = joystickImageElements[1]
	}
	var err = jImgEl.surface.Blit(
		nil,
		surface,
		&sdl.Rect{X: jImgEl.offsetX + x, Y: jImgEl.offsetY + y, W: jImgEl.surface.W, H: jImgEl.surface.H})
	if err != nil {
		panic(err)
	}
}

func drawMessages() {
	drawText(MSG_0, 10, 180, 255, 255, 0)
	drawText(MSG_1, 10, 190, 255, 255, 0)

	//last detected key
	var textWidth1 = drawText(MSG_2, TEXT_OFFSET_X, 160, 0, 255, 255)
	var textWidth2 = drawText(fmt.Sprintf("%d [0x%04X]", lastPressedKey, lastPressedKey), TEXT_OFFSET_X+TEXT_PADDING_X+textWidth1, 160, 255, 255, 255)
	var value, ok = keyNames[lastPressedKey]
	if ok {
		drawText(value, TEXT_OFFSET_X+TEXT_PADDING_X+textWidth1+TEXT_PADDING_X+textWidth2, 160, 0, 255, 255)
	} else {
		drawText("Not defined", TEXT_OFFSET_X+TEXT_PADDING_X+textWidth1+TEXT_PADDING_X+textWidth2, 160, 0, 255, 255)
	}

	//drawText(MSG_3, 10, 200, 255, 255, 0)
	drawText(MSG_4, 197, 20, 255, 255, 255)
	if isRumbleSupported {
		drawText(MSG_5, 10, 200, 255, 255, 0)
	}
}

func drawText(txt string, x int32, y int32, fR uint8, fG uint8, fB uint8) int32 {
	var text, err = font.RenderUTF8Blended(txt, sdl.Color{R: fR, G: fG, B: fB, A: 255})
	if err != nil {
		return 0
	}
	defer text.Free()
	err = text.Blit(nil, surface, &sdl.Rect{X: x, Y: y, W: 0, H: 0})
	if err != nil {
		panic(err)
	}
	return text.W
}
