package rnd

import (
	"container/list"
	"geniot.com/geniot/pg2_test_go/internal/api"
)

type Scene struct {
	renderables *list.List
}

func NewScene() *Scene {
	l := list.New()
	for i := range api.ButtonImages {
		iEl := NewImageElement(
			api.ButtonImages[i].ImageName,
			api.ButtonImages[i].OffsetX,
			api.ButtonImages[i].OffsetY,
			api.ButtonImages[i].DisplayOnPress)
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
