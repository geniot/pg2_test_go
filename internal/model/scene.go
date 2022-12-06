package model

import (
	"geniot.com/geniot/pg2_test_go/internal/utils"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	imageElements []*ImageElement
}

func NewScene(renderer *sdl.Renderer) *Scene {
	imgElements := make([]*ImageElement, len(utils.IMAGE_DESCRIPTORS))
	for i := range utils.IMAGE_DESCRIPTORS {
		iEl := NewImageElement(
			renderer,
			utils.IMAGE_DESCRIPTORS[i].ImageName,
			utils.IMAGE_DESCRIPTORS[i].OffsetX,
			utils.IMAGE_DESCRIPTORS[i].OffsetY,
			utils.IMAGE_DESCRIPTORS[i].DisplayOnPress)
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
