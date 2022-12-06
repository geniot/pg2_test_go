package model

import (
	"geniot.com/geniot/pg2_test_go/resources"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"io"
)

type Scene struct {
	imageElements mapset.Set[ImageElement]
}

func NewScene(renderer *sdl.Renderer) *Scene {
	file, _ := resources.MEDIA_LIST.Open("media/pg2_back.png")

	stat, _ := file.Stat()
	buf := make([]byte, stat.Size())
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
	}

	rwops, _ := sdl.RWFromMem(buf)

	iEl := ImageElement{}
	iEl.offsetX = 90
	iEl.offsetY = 50
	iEl.displayOnPress = sdl.K_UNKNOWN

	surface, _ := img.LoadRW(rwops, true)
	iEl.width = surface.W
	iEl.height = surface.H
	defer surface.Free()
	txt, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}
	iEl.texture = txt

	set := mapset.NewSet[ImageElement]()
	set.Add(iEl)
	return &Scene{set}
}

func (scene Scene) Render(renderer *sdl.Renderer) {
	scene.imageElements.Each(func(iEl ImageElement) bool {
		iEl.Render(renderer)
		return false
	})
}
