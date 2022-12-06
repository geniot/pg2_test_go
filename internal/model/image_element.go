package model

import (
	"geniot.com/geniot/pg2_test_go/resources"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type ImageElement struct {
	fileName       string
	offsetX        int32
	offsetY        int32
	width          int32
	height         int32
	displayOnPress sdl.Keycode
	texture        *sdl.Texture
}

func NewImageElement(renderer *sdl.Renderer, fN string, oX, oY int32, dO sdl.Keycode) *ImageElement {
	file, _ := resources.MEDIA_LIST.Open("media/" + fN)
	stat, _ := file.Stat()
	buf := make([]byte, stat.Size())
	file.Read(buf)
	rwops, _ := sdl.RWFromMem(buf)
	surface, _ := img.LoadRW(rwops, true)
	defer surface.Free()
	txt, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}
	return &ImageElement{fN, oX, oY, surface.W, surface.H, dO, txt}
}

func (iEl ImageElement) Render(renderer *sdl.Renderer, pressedKeysCodes mapset.Set[sdl.Keycode]) {
	if iEl.displayOnPress == sdl.K_UNKNOWN ||
		pressedKeysCodes.Contains(iEl.displayOnPress) {
		dstRect := sdl.Rect{iEl.offsetX, iEl.offsetY, iEl.width, iEl.height}
		renderer.Copy(iEl.texture, nil, &dstRect)
	}
}
