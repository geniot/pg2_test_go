package mdl

import (
	"container/list"
	"geniot.com/geniot/pg2_test_go/internal/api"
)

type Scene struct {
	renderables *list.List
}

func NewScene() *Scene {
	l := list.New()
	for i := range ButtonImages {
		iEl := NewImageElement(
			ButtonImages[i].ImageName,
			ButtonImages[i].OffsetX,
			ButtonImages[i].OffsetY,
			ButtonImages[i].DisplayOnPress)
		l.PushBack(iEl)
	}
	l.PushBack(NewDiskInfos())
	return &Scene{l}
}

func (scene Scene) Render() {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Render()
	}
}
