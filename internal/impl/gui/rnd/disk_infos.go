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
	renderImageElement(diskInfos.imageElements, 2)
	renderImageElement(diskInfos.imageElements, 3)

	if ctx.DiskInfo1.IsDiskAvailable {
		renderImageElement(diskInfos.imageElements, 0)

		textTexture1, w1, h1 := genTexture(ctx.DiskInfo1.FreeDiskSpace, api.COLOR_RED)
		textTexture2, w2, h2 := genTexture(" / ", api.COLOR_GREEN)
		textTexture3, w3, h3 := genTexture(ctx.DiskInfo1.MaxDiskSpace, api.COLOR_GRAY)

		ctx.Renderer.Copy(textTexture1, nil, &sdl.Rect{X: 120 - w1 - w2 - w3, Y: 20, W: w1, H: h1})
		ctx.Renderer.Copy(textTexture2, nil, &sdl.Rect{X: 120 - w2 - w3, Y: 20, W: w2, H: h2})
		ctx.Renderer.Copy(textTexture3, nil, &sdl.Rect{X: 120 - w3, Y: 20, W: w3, H: h3})
	}
	if ctx.DiskInfo2.IsDiskAvailable {
		renderImageElement(diskInfos.imageElements, 1)

		textTexture1, w1, h1 := genTexture(ctx.DiskInfo2.FreeDiskSpace, api.COLOR_RED)
		textTexture2, w2, h2 := genTexture(" / ", api.COLOR_GREEN)
		textTexture3, w3, h3 := genTexture(ctx.DiskInfo2.MaxDiskSpace, api.COLOR_GRAY)

		ctx.Renderer.Copy(textTexture1, nil, &sdl.Rect{X: 197, Y: 20, W: w1, H: h1})
		ctx.Renderer.Copy(textTexture2, nil, &sdl.Rect{X: 197 + w1, Y: 20, W: w2, H: h2})
		ctx.Renderer.Copy(textTexture3, nil, &sdl.Rect{X: 197 + w1 + w2, Y: 20, W: w3, H: h3})

	}
}
