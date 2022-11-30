package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func render() {
	//err := renderer.Clear()
	//if err != nil {
	//	panic(err)
	//}
	processKeyActions()
	redraw()

	//renderer.Present()
	var err = window.UpdateSurface()
	if err != nil {
		panic(err)
	}
}

func processKeyActions() {
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

	//background
	var err = surface.FillRect(nil, sdl.MapRGB(surface.Format, 16, 16, 16))
	if err != nil {
		panic(err)
	}

	//images
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

	drawJoystick()
	drawMessages()
	drawBattery()
}

func drawBattery() {
	var bImgEl = batteryImageElements[0]
	if powerInfo.powerState == sdl.POWERSTATE_CHARGING {
		bImgEl = batteryImageElements[1]
	} else {
		drawText(fmt.Sprintf("%2d%%", powerInfo.pct), 279, 120, 255, 255, 255)
	}
	var err = bImgEl.surface.Blit(
		nil,
		surface,
		&sdl.Rect{X: bImgEl.offsetX, Y: bImgEl.offsetY, W: bImgEl.surface.W, H: bImgEl.surface.H})
	if err != nil {
		panic(err)
	}
	var batteryLevelHeight = int32(powerInfo.pct * 39 / 100)
	if powerInfo.pct > 24 {
		err = surface.FillRect(&sdl.Rect{X: bImgEl.offsetX + 1, Y: bImgEl.offsetY + 43 - batteryLevelHeight, W: 14, H: batteryLevelHeight}, sdl.MapRGB(surface.Format, 64, 192, 64)) //green
	} else {
		err = surface.FillRect(&sdl.Rect{X: bImgEl.offsetX + 1, Y: bImgEl.offsetY + 43 - batteryLevelHeight, W: 14, H: batteryLevelHeight}, sdl.MapRGB(surface.Format, 192, 64, 64)) //red
	}
}

func drawJoystick() {
	var axisX = joystick.Axis(0)
	var axisY = joystick.Axis(1)

	var jPosX = int32(axisX / 5461)
	var jPosY = int32(axisY / 5461)
	var cPosX = 32767/1130 + int32(axisX/1130)
	var cPosY = 32767/1130 + int32(axisY/1598)

	var jImgEl = joystickImageElements[0]
	if jPosX != 0 || jPosY != 0 {
		jImgEl = joystickImageElements[1]
	}
	//left joystick in PG2v2 cannot be pressed, just moved, so I cannot test it
	if pressedKeysCodes.Contains(GCW_BUTTON_L3) {
		jImgEl = joystickImageElements[2]
	}
	var err = jImgEl.surface.Blit(
		nil,
		surface,
		&sdl.Rect{X: jImgEl.offsetX + jPosX, Y: jImgEl.offsetY + jPosY, W: jImgEl.surface.W, H: jImgEl.surface.H})
	if err != nil {
		panic(err)
	}

	drawText(fmt.Sprintf("%.2f", float32(axisX)/32767.0), 131, 69, 255, 0, 255)
	drawText(fmt.Sprintf("%.2f", float32(axisY)/32767.0), 131, 79, 255, 0, 255)

	err = surface.FillRect(&sdl.Rect{X: 131 - J_CROSS_WIDTH/2 + cPosX, Y: 62 + cPosY, W: J_CROSS_WIDTH, H: 1}, sdl.MapRGB(surface.Format, 255, 0, 255))
	err = surface.FillRect(&sdl.Rect{X: 131 + cPosX, Y: 62 - J_CROSS_WIDTH/2 + cPosY, W: 1, H: J_CROSS_WIDTH}, sdl.MapRGB(surface.Format, 255, 0, 255))
	//err = renderer.SetDrawColor(255, 0, 255, 255)
	//err = renderer.DrawLine(131-3+cPosX, 70+cPosY, 131+4+cPosX, 70+cPosY)
	//err = renderer.DrawLine(131+cPosX, 70-3+cPosY, 131+cPosX, 70+4+cPosY)

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
