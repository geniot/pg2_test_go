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

func drawText(txt string, x int32, y int32, color sdl.Color) int32 {
	textSurface, _ := ctx.Font.RenderUTF8Blended(txt, color)
	defer textSurface.Free()
	textTexture, _ := ctx.Renderer.CreateTextureFromSurface(textSurface)
	ctx.Renderer.Copy(textTexture, nil,
		&sdl.Rect{X: x, Y: y, W: textSurface.W, H: textSurface.H})
	defer textTexture.Destroy()
	return textSurface.W
}
