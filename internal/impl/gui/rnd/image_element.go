package rnd

import (
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/resources"
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

func NewImageElement(fN string, oX, oY int32, dO sdl.Keycode) ImageElement {
	surface, _ := img.LoadRW(resources.GetResource(fN), true)
	defer surface.Free()
	txt, err := ctx.Renderer.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}
	return ImageElement{fN, oX, oY, surface.W, surface.H, dO, txt}
}

func (iEl ImageElement) Render() {
	if iEl.displayOnPress == sdl.K_UNKNOWN ||
		ctx.PressedKeysCodes.Contains(iEl.displayOnPress) {
		dstRect := sdl.Rect{iEl.offsetX, iEl.offsetY, iEl.width, iEl.height}
		ctx.Renderer.Copy(iEl.texture, nil, &dstRect)
	}
}

func initImageElements(imageDescriptors []api.ImageDescriptor) []ImageElement {
	imgElements := make([]ImageElement, len(imageDescriptors))
	for i := range imageDescriptors {
		iEl := NewImageElement(
			imageDescriptors[i].ImageName,
			imageDescriptors[i].OffsetX,
			imageDescriptors[i].OffsetY,
			imageDescriptors[i].DisplayOnPress)
		imgElements[i] = iEl
	}
	return imgElements
}
