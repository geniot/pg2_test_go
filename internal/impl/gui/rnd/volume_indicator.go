package rnd

import (
	"fmt"
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

type VolumeIndicator struct {
	imageElements []ImageElement
}

func NewVolumeIndicator() api.IRenderable {
	return VolumeIndicator{initImageElements(api.VolumeImages)}
}

func (volumeIndicator VolumeIndicator) Render() {
	if ctx.CurrentVolume >= 0 {
		volumeLevelHeight := int32(ctx.CurrentVolume * 39 / 100)
		renderImageElement(volumeIndicator.imageElements, 0)
		ctx.Renderer.SetDrawColor(64, 192, 64, 255) //green
		ctx.Renderer.FillRect(
			&sdl.Rect{
				X: volumeIndicator.imageElements[0].offsetX + 1,
				Y: volumeIndicator.imageElements[0].offsetY + 43 - volumeLevelHeight,
				W: 14, H: volumeLevelHeight})
		drawText(fmt.Sprintf("%2d%%", ctx.CurrentVolume), 24, 120, api.COLOR_WHITE)
		drawText("VOL", 20, 55, api.COLOR_WHITE)
	}
}
