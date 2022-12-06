package model

import (
	"geniot.com/geniot/pg2_test_go/internal/utils"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type Scene struct {
	imageElements mapset.Set[*ImageElement]
}

func NewScene(renderer *sdl.Renderer) *Scene {
	set := mapset.NewSet[*ImageElement]()
	for i, _ := range utils.IMAGE_DESCRIPTORS {
		iEl := NewImageElement(
			renderer,
			utils.IMAGE_DESCRIPTORS[i].ImageName,
			utils.IMAGE_DESCRIPTORS[i].OffsetX,
			utils.IMAGE_DESCRIPTORS[i].OffsetY,
			utils.IMAGE_DESCRIPTORS[i].DisplayOnPress)
		set.Add(iEl)
	}
	return &Scene{set}
}

func (scene Scene) Render(renderer *sdl.Renderer, pressedKeysCodes mapset.Set[sdl.Keycode]) {
	scene.imageElements.Each(func(iEl *ImageElement) bool {
		iEl.Render(renderer, pressedKeysCodes)
		return false
	})
}
