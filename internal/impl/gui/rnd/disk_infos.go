package rnd

import (
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

type DiskInfos struct {
	imageElements []ImageElement
}

func NewDiskInfos() api.IRenderable {
	return DiskInfos{initImageElements(api.DiskImages)}
}

func (diskInfos DiskInfos) Render() {
	diskInfos.renderImageElement(2)
	diskInfos.renderImageElement(3)

	if ctx.DiskInfo1.IsDiskAvailable {
		diskInfos.renderImageElement(0)

		textTexture1, w1, h1 := diskInfos.genTexture(ctx.DiskInfo1.FreeDiskSpace, api.COLOR_RED)
		textTexture2, w2, h2 := diskInfos.genTexture(" / ", api.COLOR_GREEN)
		textTexture3, w3, h3 := diskInfos.genTexture(ctx.DiskInfo1.MaxDiskSpace, api.COLOR_GRAY)

		ctx.Renderer.Copy(textTexture1, nil, &sdl.Rect{X: 120 - w1 - w2 - w3, Y: 20, W: w1, H: h1})
		ctx.Renderer.Copy(textTexture2, nil, &sdl.Rect{X: 120 - w2 - w3, Y: 20, W: w2, H: h2})
		ctx.Renderer.Copy(textTexture3, nil, &sdl.Rect{X: 120 - w3, Y: 20, W: w3, H: h3})
	}
	if ctx.DiskInfo2.IsDiskAvailable {
		diskInfos.renderImageElement(1)

		textTexture1, w1, h1 := diskInfos.genTexture(ctx.DiskInfo2.FreeDiskSpace, api.COLOR_RED)
		textTexture2, w2, h2 := diskInfos.genTexture(" / ", api.COLOR_GREEN)
		textTexture3, w3, h3 := diskInfos.genTexture(ctx.DiskInfo2.MaxDiskSpace, api.COLOR_GRAY)

		ctx.Renderer.Copy(textTexture1, nil, &sdl.Rect{X: 197, Y: 20, W: w1, H: h1})
		ctx.Renderer.Copy(textTexture2, nil, &sdl.Rect{X: 197 + w1, Y: 20, W: w2, H: h2})
		ctx.Renderer.Copy(textTexture3, nil, &sdl.Rect{X: 197 + w1 + w2, Y: 20, W: w3, H: h3})

	}
}

func (diskInfos DiskInfos) genTexture(text string, color sdl.Color) (*sdl.Texture, int32, int32) {
	textSurface, _ := ctx.Font.RenderUTF8Blended(text, color)
	defer textSurface.Free()
	textTexture, _ := ctx.Renderer.CreateTextureFromSurface(textSurface)
	return textTexture, textSurface.W, textSurface.H
}

func (diskInfos DiskInfos) renderImageElement(index int) {
	ctx.Renderer.Copy(
		diskInfos.imageElements[index].texture,
		nil,
		&sdl.Rect{
			diskInfos.imageElements[index].offsetX,
			diskInfos.imageElements[index].offsetY,
			diskInfos.imageElements[index].width,
			diskInfos.imageElements[index].height})
}
