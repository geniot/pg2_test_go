package model

import (
	"geniot.com/geniot/pg2_test_go/internal/imm"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	imageElements []*ImageElement
}

func NewScene(renderer *sdl.Renderer) *Scene {
	imgElements := make([]*ImageElement, len(imm.ImageDescriptors))
	for i := range imm.ImageDescriptors {
		iEl := NewImageElement(
			renderer,
			imm.ImageDescriptors[i].ImageName,
			imm.ImageDescriptors[i].OffsetX,
			imm.ImageDescriptors[i].OffsetY,
			imm.ImageDescriptors[i].DisplayOnPress)
		imgElements[i] = iEl
	}
	return &Scene{imgElements}
}

func (scene Scene) Render(renderer *sdl.Renderer,
	pressedKeysCodes mapset.Set[sdl.Keycode]) {
	for i := range scene.imageElements {
		scene.imageElements[i].Render(renderer, pressedKeysCodes)
	}
}
