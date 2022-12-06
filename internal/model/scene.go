package model

import (
	"geniot.com/geniot/pg2_test_go/internal/utils"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"io"
	"unsafe"
)

type Scene struct {
	imageElements mapset.Set[ImageElement]
}

func NewScene() *Scene {
	s, _ := img.Load("media/bmp_24.bmp")
	println(s.Pitch)
	println(s.Format)

	file, _ := utils.MEDIA_LIST.Open("media/bmp_24.bmp")
	stat, _ := file.Stat()
	buf := make([]byte, stat.Size())
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
	}

	iEl := ImageElement{}
	iEl.offsetX = 90
	iEl.offsetY = 50
	iEl.displayOnPress = sdl.K_UNKNOWN
	ptr := unsafe.Pointer(&buf[0])
	iEl.surface, _ = sdl.CreateRGBSurfaceWithFormatFrom(
		ptr,
		200,
		200,
		24,
		600,
		390076419,
	)

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
