package bak

import (
	"fmt"
	"geniot.com/geniot/pg2_test_go/bak"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"runtime"
)

func render() {
	//err := renderer.Clear()
	//if err != nil {
	//	panic(err)
	//}
	processKeyActions()
	redraw()

	//renderer.Present()
	var err = bak.window.UpdateSurface()
	if err != nil {
		panic(err)
	}
}

func processKeyActions() {
	if bak.pressedKeysCodes.Contains(sdl.K_q) ||
		(bak.pressedKeysCodes.Contains(GCW_BUTTON_L1) && bak.pressedKeysCodes.Contains(GCW_BUTTON_START)) {
		bak.running = false
	}
	if bak.pressedKeysCodes.Contains(GCW_BUTTON_L1) && bak.pressedKeysCodes.Contains(GCW_BUTTON_X) && mix.Playing(-1) != 1 {
		_, err := bak.audioChunk.Play(1, 0)
		if err != nil {
			panic(err)
		}
	}
	if bak.isRumbleSupported {
		if bak.pressedKeysCodes.Contains(GCW_BUTTON_L2) && bak.pressedKeysCodes.Contains(GCW_BUTTON_R2) {
			var err = bak.haptic.RumblePlay(0.33, 500)
			if err != nil {
				println(err.Error())
			}
		}
	}
}

func redraw() {

	//background
	var err = bak.surface.FillRect(nil, sdl.MapRGB(bak.surface.Format, 16, 16, 16))
	if err != nil {
		panic(err)
	}

	//images
	for _, imgEl := range bak.imageElements {
		if imgEl.displayOnPress == sdl.K_UNKNOWN ||
			bak.pressedKeysCodes.Contains(imgEl.displayOnPress) {
			var err = imgEl.surface.Blit(
				nil,
				bak.surface,
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
	bak.diskImageElements[2].surface.Blit(nil, bak.surface, &sdl.Rect{X: bak.diskImageElements[2].offsetX, Y: bak.diskImageElements[2].offsetY, W: bak.diskImageElements[2].surface.W, H: bak.diskImageElements[2].surface.H})
	bak.diskImageElements[3].surface.Blit(nil, bak.surface, &sdl.Rect{X: bak.diskImageElements[3].offsetX, Y: bak.diskImageElements[3].offsetY, W: bak.diskImageElements[3].surface.W, H: bak.diskImageElements[3].surface.H})
	if bak.diskInfos[0].isDiskAvailable {
		bak.diskImageElements[0].surface.Blit(nil, bak.surface, &sdl.Rect{X: bak.diskImageElements[0].offsetX, Y: bak.diskImageElements[0].offsetY, W: bak.diskImageElements[0].surface.W, H: bak.diskImageElements[0].surface.H})
		var text1, _ = bak.font.RenderUTF8Blended(bak.diskInfos[0].freeDiskSpace, COLOR_RED)
		var text2, _ = bak.font.RenderUTF8Blended(" / ", COLOR_GREEN)
		var text3, _ = bak.font.RenderUTF8Blended(bak.diskInfos[0].maxDiskSpace, COLOR_GRAY)
		defer text1.Free()
		defer text2.Free()
		defer text3.Free()
		text1.Blit(nil, bak.surface, &sdl.Rect{X: 120 - text1.W - text2.W - text3.W, Y: 20, W: 0, H: 0})
		text2.Blit(nil, bak.surface, &sdl.Rect{X: 120 - text2.W - text3.W, Y: 20, W: 0, H: 0})
		text3.Blit(nil, bak.surface, &sdl.Rect{X: 120 - text3.W, Y: 20, W: 0, H: 0})
	}
	if bak.diskInfos[1].isDiskAvailable {
		bak.diskImageElements[1].surface.Blit(nil, bak.surface, &sdl.Rect{X: bak.diskImageElements[1].offsetX, Y: bak.diskImageElements[1].offsetY, W: bak.diskImageElements[1].surface.W, H: bak.diskImageElements[1].surface.H})
		var txtWidth1 = drawText(bak.diskInfos[1].freeDiskSpace, 197, 20, COLOR_RED)
		var txtWidth2 = drawText(" / ", 197+txtWidth1, 20, COLOR_GREEN)
		drawText(bak.diskInfos[1].maxDiskSpace, 197+txtWidth1+txtWidth2, 20, COLOR_GRAY)
	}
}

func drawVolume() {
	if runtime.GOOS == "windows" {
		var volumeLevelHeight = int32(bak.currentVolume * 39 / 100)
		bak.volumeImageElements[0].surface.Blit(nil, bak.surface, &sdl.Rect{X: bak.volumeImageElements[0].offsetX, Y: bak.volumeImageElements[0].offsetY, W: bak.volumeImageElements[0].surface.W, H: bak.volumeImageElements[0].surface.H})
		bak.surface.FillRect(&sdl.Rect{X: bak.volumeImageElements[0].offsetX + 1, Y: bak.volumeImageElements[0].offsetY + 43 - volumeLevelHeight, W: 14, H: volumeLevelHeight}, sdl.MapRGB(bak.surface.Format, 64, 192, 64)) //green
		drawText(fmt.Sprintf("%2d%%", bak.currentVolume), 24, 120, COLOR_WHITE)
		drawText("VOL", 20, 55, COLOR_WHITE)
	}
}

func drawBattery() {
	var bImgEl = bak.batteryImageElements[0]
	var bChImgEl = bak.batteryImageElements[1]
	drawText(fmt.Sprintf("%2d%%", bak.powerInfo.pct), 279, 120, COLOR_WHITE)
	drawText("POW", 275, 55, COLOR_WHITE)

	bImgEl.surface.Blit(
		nil,
		bak.surface,
		&sdl.Rect{X: bImgEl.offsetX, Y: bImgEl.offsetY, W: bImgEl.surface.W, H: bImgEl.surface.H})

	var batteryLevelHeight = int32(bak.powerInfo.pct * 39 / 100)
	if bak.powerInfo.pct > 24 {
		bak.surface.FillRect(&sdl.Rect{X: bImgEl.offsetX + 1, Y: bImgEl.offsetY + 43 - batteryLevelHeight, W: 14, H: batteryLevelHeight}, sdl.MapRGB(bak.surface.Format, 64, 192, 64)) //green
	} else {
		bak.surface.FillRect(&sdl.Rect{X: bImgEl.offsetX + 1, Y: bImgEl.offsetY + 43 - batteryLevelHeight, W: 14, H: batteryLevelHeight}, sdl.MapRGB(bak.surface.Format, 192, 64, 64)) //red
	}

	if bak.powerInfo.isCharging {
		bChImgEl.surface.Blit(
			nil,
			bak.surface,
			&sdl.Rect{X: bChImgEl.offsetX, Y: bChImgEl.offsetY, W: bChImgEl.surface.W, H: bChImgEl.surface.H})

	}
}

func drawJoystick() {
	var axisX = bak.joystick.Axis(0)
	var axisY = bak.joystick.Axis(1)

	var axisPctX = float32(axisX) * 100 / 32767.0
	var axisPctY = float32(axisY) * 100 / 32767.0

	var jPosX = int32(axisX / 5461)
	var jPosY = int32(axisY / 5461)

	var cPosX = (SMALL_SCREEN_WIDTH / 2 * axisPctX / 50) / 2
	var cPosY = (SMALL_SCREEN_HEIGHT / 2 * axisPctY / 50) / 2

	var jImgEl = bak.joystickImageElements[0]
	if jPosX != 0 || jPosY != 0 {
		jImgEl = bak.joystickImageElements[1]
	}
	//left joystick in PG2v2 cannot be pressed, just moved, so I cannot test it
	if bak.pressedKeysCodes.Contains(GCW_BUTTON_L3) {
		jImgEl = bak.joystickImageElements[2]
	}
	var err = jImgEl.surface.Blit(
		nil,
		bak.surface,
		&sdl.Rect{X: jImgEl.offsetX + jPosX, Y: jImgEl.offsetY + jPosY, W: jImgEl.surface.W, H: jImgEl.surface.H})
	if err != nil {
		panic(err)
	}

	drawText(fmt.Sprintf("%.2f", float32(axisX)/32767.0), 131, 69, COLOR_PURPLE)
	drawText(fmt.Sprintf("%.2f", float32(axisY)/32767.0), 131, 79, COLOR_PURPLE)

	//white rectangle can be used for debugging the small screen area
	//err = surface.FillRect(&sdl.Rect{X: SMALL_SCREEN_X1, Y: SMALL_SCREEN_Y1, W: SMALL_SCREEN_WIDTH, H: SMALL_SCREEN_HEIGHT}, sdl.MapRGB(surface.Format, 255, 255, 255))

	err = bak.surface.FillRect(&sdl.Rect{X: int32(SMALL_SCREEN_X_CENTER - J_CROSS_WIDTH/2 + cPosX), Y: int32(SMALL_SCREEN_Y_CENTER + cPosY), W: J_CROSS_WIDTH, H: 1}, sdl.MapRGB(bak.surface.Format, 255, 0, 255))
	err = bak.surface.FillRect(&sdl.Rect{X: int32(SMALL_SCREEN_X_CENTER + cPosX), Y: int32(SMALL_SCREEN_Y_CENTER - J_CROSS_WIDTH/2 + cPosY), W: 1, H: J_CROSS_WIDTH}, sdl.MapRGB(bak.surface.Format, 255, 0, 255))

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
	var textWidth2 = drawText(fmt.Sprintf("%d [0x%04X]", bak.lastPressedKey, bak.lastPressedKey), TEXT_OFFSET_X+TEXT_PADDING_X+textWidth1, 160, COLOR_WHITE)
	var value, ok = bak.keyNames[bak.lastPressedKey]
	if ok {
		drawText(value, TEXT_OFFSET_X+TEXT_PADDING_X+textWidth1+TEXT_PADDING_X+textWidth2, 160, COLOR_BLUE)
	} else {
		drawText("Not defined", TEXT_OFFSET_X+TEXT_PADDING_X+textWidth1+TEXT_PADDING_X+textWidth2, 160, COLOR_BLUE)
	}

	//drawText(MSG_3, 10, 200, 255, 255, 0)
	//drawText(MSG_4, 197, 20, 255, 255, 255)
	drawText(MSG_5, 10, 200, COLOR_YELLOW)
}

func drawText(txt string, x int32, y int32, color sdl.Color) int32 {
	var text, err = bak.font.RenderUTF8Blended(txt, color)
	if err != nil {
		return 0
	}
	defer text.Free()
	err = text.Blit(nil, bak.surface, &sdl.Rect{X: x, Y: y, W: 0, H: 0})
	if err != nil {
		panic(err)
	}
	return text.W
}
