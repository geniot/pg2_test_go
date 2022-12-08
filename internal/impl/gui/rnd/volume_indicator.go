package rnd

import (
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
)

type VolumeIndicator struct {
	imageElements []ImageElement
}

func NewVolumeIndicator() api.IRenderable {
	return VolumeIndicator{initImageElements(api.VolumeImages)}
}

func (volumeIndicator VolumeIndicator) Render() {
	if ctx.CurrentVolume >= 0 {
		//volumeLevelHeight := int32(ctx.CurrentVolume * 39 / 100)
		renderImageElement(volumeIndicator.imageElements, 0)
		//bak.volumeImageElements[0].surface.Blit(nil, bak.surface, &sdl.Rect{X: bak.volumeImageElements[0].offsetX, Y: bak.volumeImageElements[0].offsetY, W: bak.volumeImageElements[0].surface.W, H: bak.volumeImageElements[0].surface.H})
		//bak.surface.FillRect(&sdl.Rect{X: bak.volumeImageElements[0].offsetX + 1, Y: bak.volumeImageElements[0].offsetY + 43 - volumeLevelHeight, W: 14, H: volumeLevelHeight}, sdl.MapRGB(bak.surface.Format, 64, 192, 64)) //green
		//drawText(fmt.Sprintf("%2d%%", bak.currentVolume), 24, 120, COLOR_WHITE)
		//drawText("VOL", 20, 55, COLOR_WHITE)
	}
}
