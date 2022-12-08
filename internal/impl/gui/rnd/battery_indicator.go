package rnd

import (
	"fmt"
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

type BatteryIndicator struct {
	imageElements []ImageElement
}

func NewBatteryIndicator() api.IRenderable {
	return BatteryIndicator{initImageElements(api.BatteryImages)}
}

func (batteryIndicator BatteryIndicator) Render() {
	drawText(fmt.Sprintf("%2d%%", ctx.PowerInformation.Pct), 279, 120, api.COLOR_WHITE)
	drawText("POW", 275, 55, api.COLOR_WHITE)
	renderImageElement(batteryIndicator.imageElements, 0)
	batteryLevelHeight := int32(ctx.PowerInformation.Pct * 39 / 100)
	if ctx.PowerInformation.Pct > 24 {
		ctx.Renderer.SetDrawColor(64, 192, 64, 255) //green
	} else {
		ctx.Renderer.SetDrawColor(192, 64, 64, 255) //red
	}
	ctx.Renderer.FillRect(
		&sdl.Rect{
			X: batteryIndicator.imageElements[0].offsetX + 1,
			Y: batteryIndicator.imageElements[0].offsetY + 43 - batteryLevelHeight,
			W: 14,
			H: batteryLevelHeight})

	if ctx.PowerInformation.IsCharging {
		renderImageElement(batteryIndicator.imageElements, 1)
	}
}
