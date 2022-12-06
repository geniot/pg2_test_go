package mdl

import (
	"geniot.com/geniot/pg2_test_go/internal/imm"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	imageElements []*ImageElement
}

func NewScene(renderer *sdl.Renderer) *Scene {
	l := imm.GetImageDescriptorsLength()
	imgElements := make([]*ImageElement, l)
	for i := 0; i < l; i++ {
		oX, oY, fN, dO := imm.GetImageDescriptorPropsByIndex(i)
		iEl := NewImageElement(renderer, fN, oX, oY, dO)
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
