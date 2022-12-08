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

func getTextWidth(text string) int32 {
	width, _, _ := ctx.Font.SizeUTF8(text)
	return int32(width)
}

func genTexture(text string, color sdl.Color) (*sdl.Texture, int32, int32) {
	textSurface, _ := ctx.Font.RenderUTF8Blended(text, color)
	defer textSurface.Free()
	textTexture, _ := ctx.Renderer.CreateTextureFromSurface(textSurface)
	return textTexture, textSurface.W, textSurface.H
}

func drawText(txt string, x int32, y int32, color sdl.Color) int32 {
	textTexture, width, height := genTexture(txt, color)
	ctx.Renderer.Copy(textTexture, nil,
		&sdl.Rect{X: x, Y: y, W: width, H: height})
	return width
}
