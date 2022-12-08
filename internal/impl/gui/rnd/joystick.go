package rnd

import (
	"fmt"
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
)

type Joystick struct {
	imageElements []ImageElement
}

func NewJoystick() api.IRenderable {
	return Joystick{initImageElements(api.JoystickImages)}
}

func (joystick Joystick) Render() {
	axisX := ctx.Device.GetJoystickAxis(0)
	axisY := ctx.Device.GetJoystickAxis(1)

	axisPctX := float32(axisX) * 100 / 32767.0
	axisPctY := float32(axisY) * 100 / 32767.0

	jPosX := int32(axisX / 5461)
	jPosY := int32(axisY / 5461)

	var cPosX = (glb.SMALL_SCREEN_WIDTH / 2 * axisPctX / 50) / 2
	var cPosY = (glb.SMALL_SCREEN_HEIGHT / 2 * axisPctY / 50) / 2

	var jImgEl = joystick.imageElements[0]
	if jPosX != 0 || jPosY != 0 {
		jImgEl = joystick.imageElements[1]
	}
	//left joystick in PG2v2 cannot be pressed, just moved, so I cannot test it
	if ctx.PressedKeysCodes.Contains(glb.GCW_BUTTON_L3) {
		jImgEl = joystick.imageElements[2]
	}

	ctx.Renderer.Copy(
		jImgEl.texture,
		nil,
		&sdl.Rect{
			X: jImgEl.offsetX + jPosX,
			Y: jImgEl.offsetY + jPosY,
			W: jImgEl.width, H: jImgEl.height})

	drawText(fmt.Sprintf("%.2f", float32(axisX)/32767.0), 131, 69, glb.COLOR_PURPLE)
	drawText(fmt.Sprintf("%.2f", float32(axisY)/32767.0), 131, 79, glb.COLOR_PURPLE)

	//white rectangle can be used for debugging the small screen area
	//err = surface.FillRect(&sdl.Rect{X: SMALL_SCREEN_X1, Y: SMALL_SCREEN_Y1, W: SMALL_SCREEN_WIDTH, H: SMALL_SCREEN_HEIGHT}, sdl.MapRGB(surface.Format, 255, 255, 255))

	ctx.Renderer.SetDrawColor(255, 0, 255, 255) //fuchsia
	ctx.Renderer.DrawLine(
		int32(glb.SMALL_SCREEN_X_CENTER-glb.J_CROSS_WIDTH/2+cPosX),
		int32(glb.SMALL_SCREEN_Y_CENTER+cPosY),
		int32(glb.SMALL_SCREEN_X_CENTER+glb.J_CROSS_WIDTH/2+cPosX),
		int32(glb.SMALL_SCREEN_Y_CENTER+cPosY))
	ctx.Renderer.DrawLine(
		int32(glb.SMALL_SCREEN_X_CENTER+cPosX),
		int32(glb.SMALL_SCREEN_Y_CENTER-glb.J_CROSS_HEIGHT/2+cPosY),
		int32(glb.SMALL_SCREEN_X_CENTER+cPosX),
		int32(glb.SMALL_SCREEN_Y_CENTER+glb.J_CROSS_HEIGHT/2+cPosY))
}
