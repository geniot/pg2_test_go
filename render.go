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
	drawDisks()
	drawVolume()
}

func drawDisks() {
	diskImageElements[2].surface.Blit(nil, surface, &sdl.Rect{X: diskImageElements[2].offsetX, Y: diskImageElements[2].offsetY, W: diskImageElements[2].surface.W, H: diskImageElements[2].surface.H})
	diskImageElements[3].surface.Blit(nil, surface, &sdl.Rect{X: diskImageElements[3].offsetX, Y: diskImageElements[3].offsetY, W: diskImageElements[3].surface.W, H: diskImageElements[3].surface.H})
	if diskInfos[0].isDiskAvailable {
		diskImageElements[0].surface.Blit(nil, surface, &sdl.Rect{X: diskImageElements[0].offsetX, Y: diskImageElements[0].offsetY, W: diskImageElements[0].surface.W, H: diskImageElements[0].surface.H})
		var text1, _ = font.RenderUTF8Blended(diskInfos[0].freeDiskSpace, COLOR_RED)
		var text2, _ = font.RenderUTF8Blended(" / ", COLOR_GREEN)
		var text3, _ = font.RenderUTF8Blended(diskInfos[0].maxDiskSpace, COLOR_GRAY)
		defer text1.Free()
		defer text2.Free()
		defer text3.Free()
		text1.Blit(nil, surface, &sdl.Rect{X: 120 - text1.W - text2.W - text3.W, Y: 20, W: 0, H: 0})
		text2.Blit(nil, surface, &sdl.Rect{X: 120 - text2.W - text3.W, Y: 20, W: 0, H: 0})
		text3.Blit(nil, surface, &sdl.Rect{X: 120 - text3.W, Y: 20, W: 0, H: 0})
	}
	if diskInfos[1].isDiskAvailable {
		diskImageElements[1].surface.Blit(nil, surface, &sdl.Rect{X: diskImageElements[1].offsetX, Y: diskImageElements[1].offsetY, W: diskImageElements[1].surface.W, H: diskImageElements[1].surface.H})
		var txtWidth1 = drawText(diskInfos[1].freeDiskSpace, 197, 20, COLOR_RED)
		var txtWidth2 = drawText(" / ", 197+txtWidth1, 20, COLOR_GREEN)
		drawText(diskInfos[1].maxDiskSpace, 197+txtWidth1+txtWidth2, 20, COLOR_GRAY)
	}
}

func drawVolume() {
	var volumeLevelHeight = int32(currentVolume * 39 / 100)
	volumeImageElements[0].surface.Blit(nil, surface, &sdl.Rect{X: volumeImageElements[0].offsetX, Y: volumeImageElements[0].offsetY, W: volumeImageElements[0].surface.W, H: volumeImageElements[0].surface.H})
	surface.FillRect(&sdl.Rect{X: volumeImageElements[0].offsetX + 1, Y: volumeImageElements[0].offsetY + 43 - volumeLevelHeight, W: 14, H: volumeLevelHeight}, sdl.MapRGB(surface.Format, 64, 192, 64)) //green
	drawText(fmt.Sprintf("%2d%%", currentVolume), 24, 120, COLOR_WHITE)
	drawText("VOL", 20, 55, COLOR_WHITE)
}

func drawBattery() {
	var bImgEl = batteryImageElements[0]
	var bChImgEl = batteryImageElements[1]
	drawText(fmt.Sprintf("%2d%%", powerInfo.pct), 279, 120, COLOR_WHITE)
	drawText("POW", 275, 55, COLOR_WHITE)

	bImgEl.surface.Blit(
		nil,
		surface,
		&sdl.Rect{X: bImgEl.offsetX, Y: bImgEl.offsetY, W: bImgEl.surface.W, H: bImgEl.surface.H})

	var batteryLevelHeight = int32(powerInfo.pct * 39 / 100)
	if powerInfo.pct > 24 {
		surface.FillRect(&sdl.Rect{X: bImgEl.offsetX + 1, Y: bImgEl.offsetY + 43 - batteryLevelHeight, W: 14, H: batteryLevelHeight}, sdl.MapRGB(surface.Format, 64, 192, 64)) //green
	} else {
		surface.FillRect(&sdl.Rect{X: bImgEl.offsetX + 1, Y: bImgEl.offsetY + 43 - batteryLevelHeight, W: 14, H: batteryLevelHeight}, sdl.MapRGB(surface.Format, 192, 64, 64)) //red
	}

	if powerInfo.isCharging {
		bChImgEl.surface.Blit(
			nil,
			surface,
			&sdl.Rect{X: bChImgEl.offsetX, Y: bChImgEl.offsetY, W: bChImgEl.surface.W, H: bChImgEl.surface.H})

	}
}

func drawJoystick() {
	var axisX = joystick.Axis(0)
	var axisY = joystick.Axis(1)

	var axisPctX = float32(axisX) * 100 / 32767.0
	var axisPctY = float32(axisY) * 100 / 32767.0

	var jPosX = int32(axisX / 5461)
	var jPosY = int32(axisY / 5461)

	var cPosX = (SMALL_SCREEN_WIDTH / 2 * axisPctX / 50) / 2
	var cPosY = (SMALL_SCREEN_HEIGHT / 2 * axisPctY / 50) / 2

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

	drawText(fmt.Sprintf("%.2f", float32(axisX)/32767.0), 131, 69, COLOR_PURPLE)
	drawText(fmt.Sprintf("%.2f", float32(axisY)/32767.0), 131, 79, COLOR_PURPLE)

	//white rectangle can be used for debugging the small screen area
	//err = surface.FillRect(&sdl.Rect{X: SMALL_SCREEN_X1, Y: SMALL_SCREEN_Y1, W: SMALL_SCREEN_WIDTH, H: SMALL_SCREEN_HEIGHT}, sdl.MapRGB(surface.Format, 255, 255, 255))

	err = surface.FillRect(&sdl.Rect{X: int32(SMALL_SCREEN_X_CENTER - J_CROSS_WIDTH/2 + cPosX), Y: int32(SMALL_SCREEN_Y_CENTER + cPosY), W: J_CROSS_WIDTH, H: 1}, sdl.MapRGB(surface.Format, 255, 0, 255))
	err = surface.FillRect(&sdl.Rect{X: int32(SMALL_SCREEN_X_CENTER + cPosX), Y: int32(SMALL_SCREEN_Y_CENTER - J_CROSS_WIDTH/2 + cPosY), W: 1, H: J_CROSS_WIDTH}, sdl.MapRGB(surface.Format, 255, 0, 255))

	//correct way of drawing lines, caused blinking in clear-present, using FillRect for now
	//err = renderer.SetDrawColor(255, 0, 255, 255)
	//err = renderer.DrawLine(131-3+cPosX, 70+cPosY, 131+4+cPosX, 70+cPosY)
	//err = renderer.DrawLine(131+cPosX, 70-3+cPosY, 131+cPosX, 70+4+cPosY)

}

func drawMessages() {
	drawText(MSG_0, 10, 180, COLOR_YELLOW)
	drawText(MSG_1, 10, 190, COLOR_YELLOW)

	//last detected key
	var textWidth1 = drawText(MSG_2, TEXT_OFFSET_X, 160, COLOR_BLUE)
	var textWidth2 = drawText(fmt.Sprintf("%d [0x%04X]", lastPressedKey, lastPressedKey), TEXT_OFFSET_X+TEXT_PADDING_X+textWidth1, 160, COLOR_WHITE)
	var value, ok = keyNames[lastPressedKey]
	if ok {
		drawText(value, TEXT_OFFSET_X+TEXT_PADDING_X+textWidth1+TEXT_PADDING_X+textWidth2, 160, COLOR_BLUE)
	} else {
		drawText("Not defined", TEXT_OFFSET_X+TEXT_PADDING_X+textWidth1+TEXT_PADDING_X+textWidth2, 160, COLOR_BLUE)
	}

	//drawText(MSG_3, 10, 200, 255, 255, 0)
	//drawText(MSG_4, 197, 20, 255, 255, 255)
	if isRumbleSupported {
		drawText(MSG_5, 10, 200, COLOR_YELLOW)
	}
}

func drawText(txt string, x int32, y int32, color sdl.Color) int32 {
	var text, err = font.RenderUTF8Blended(txt, color)
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
