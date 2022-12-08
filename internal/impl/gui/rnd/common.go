package rnd

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

func renderImageElement(imageElements []ImageElement, index int) {
	ctx.Renderer.Copy(
		imageElements[index].texture,
		nil,
		&sdl.Rect{
			imageElements[index].offsetX,
			imageElements[index].offsetY,
			imageElements[index].width,
			imageElements[index].height})
}

func genTexture(text string, color sdl.Color) (*sdl.Texture, int32, int32) {
	textSurface, _ := ctx.Font.RenderUTF8Blended(text, color)
	defer textSurface.Free()
	textTexture, _ := ctx.Renderer.CreateTextureFromSurface(textSurface)
	return textTexture, textSurface.W, textSurface.H
}
