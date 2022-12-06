package model

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ImageElement struct {
	offsetX        int32
	offsetY        int32
	imageName      string
	surface        *sdl.Surface
	texture        *sdl.Texture
	displayOnPress sdl.Keycode
}

func NewImageElement() *ImageElement {
	return &ImageElement{}
}

func (iEl ImageElement) Render(renderer *sdl.Renderer) {
	txt, err := renderer.CreateTextureFromSurface(iEl.surface)
	if err != nil {
		println(err.Error())
	}
	iEl.texture = txt

	srcRect := sdl.Rect{0, 0, 200, 200}
	dstRect := sdl.Rect{0, 0, 200, 200}
	renderer.Copy(iEl.texture, &srcRect, &dstRect)
}
