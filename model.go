package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type ImageElement struct {
	offsetX        int32
	offsetY        int32
	imageName      string
	surface        *sdl.Surface
	displayOnPress sdl.Keycode
}

func initImageElements() {
	imageElements = []ImageElement{
		{
			offsetX:        90,
			offsetY:        50,
			imageName:      "media/pg2_back.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        226,
			offsetY:        73,
			imageName:      "media/info_btna.png",
			displayOnPress: GCW_BUTTON_A,
		},
		{
			offsetX:        218,
			offsetY:        90,
			imageName:      "media/info_btnb.png",
			displayOnPress: GCW_BUTTON_B,
		},
		{
			offsetX:        218,
			offsetY:        64,
			imageName:      "media/info_btnx.png",
			displayOnPress: GCW_BUTTON_X,
		},
		{
			offsetX:        211,
			offsetY:        82,
			imageName:      "media/info_btny.png",
			displayOnPress: GCW_BUTTON_Y,
		},
		{
			offsetX:        71,
			offsetY:        76,
			imageName:      "media/info_padup.png",
			displayOnPress: GCW_BUTTON_UP,
		},
		{
			offsetX:        59,
			offsetY:        92,
			imageName:      "media/info_paddown.png",
			displayOnPress: GCW_BUTTON_DOWN,
		},
		{
			offsetX:        63,
			offsetY:        73,
			imageName:      "media/info_padleft.png",
			displayOnPress: GCW_BUTTON_LEFT,
		},
		{
			offsetX:        58,
			offsetY:        81,
			imageName:      "media/info_padright.png",
			displayOnPress: GCW_BUTTON_RIGHT,
		},
	}
	for i, imageElement := range imageElements {
		var pngImage, err = img.Load(imageElement.imageName)
		if err != nil {
			panic(err)
		}
		imageElements[i].surface = pngImage
	}
}

func closeImageElements() {
	for i, _ := range imageElements {
		imageElements[i].surface.Free()
	}
}
