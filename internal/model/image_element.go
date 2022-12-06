package model

import (
	"github.com/veandco/go-sdl2/sdl"
)

type ImageElement struct {
	offsetX        int32
	offsetY        int32
	width          int32
	height         int32
	texture        *sdl.Texture
	displayOnPress sdl.Keycode
}

func NewImageElement() *ImageElement {
	return &ImageElement{}
}

func (iEl ImageElement) Render(renderer *sdl.Renderer) {
	srcRect := sdl.Rect{0, 0, iEl.width, iEl.height}
	dstRect := sdl.Rect{iEl.offsetY, iEl.offsetY, iEl.width, iEl.height}
	renderer.Copy(iEl.texture, &srcRect, &dstRect)
}
