package main

import (
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type ImageElement struct {
	offsetX   int32
	offsetY   int32
	imageName string
	surface   *sdl.Surface
}

func initImageElements() {
	imageElements = []ImageElement{
		{
			offsetX:   90,
			offsetY:   50,
			imageName: "media/pg2_back.png",
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
