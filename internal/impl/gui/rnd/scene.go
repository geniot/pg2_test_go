package rnd

import (
	"container/list"
	"fmt"
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/glb"
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
	l.PushBack(NewVolumeIndicator())
	l.PushBack(NewBatteryIndicator())
	l.PushBack(NewJoystick())
	return &Scene{l}
}

func (scene Scene) Render() {
	for e := scene.renderables.Front(); e != nil; e = e.Next() {
		e.Value.(api.IRenderable).Render()
	}
	scene.renderMessages()
}

func (scene Scene) renderMessages() {
	drawText(glb.MSG_0, 10, 180, glb.COLOR_YELLOW)
	drawText(glb.MSG_1, 10, 190, glb.COLOR_YELLOW)

	//last detected key
	var textWidth1 = drawText(glb.MSG_2, glb.TEXT_OFFSET_X, 160, glb.COLOR_BLUE)
	var textWidth2 = drawText(
		fmt.Sprintf("%d [0x%04X]",
			ctx.LastPressedKey, ctx.LastPressedKey),
		glb.TEXT_OFFSET_X+glb.TEXT_PADDING_X+textWidth1,
		160, glb.COLOR_WHITE)

	var value, ok = glb.KeyNames[ctx.LastPressedKey]
	if ok {
		drawText(
			value,
			glb.TEXT_OFFSET_X+glb.TEXT_PADDING_X+textWidth1+glb.TEXT_PADDING_X+textWidth2,
			160, glb.COLOR_BLUE)
	} else {
		drawText("Not defined",
			glb.TEXT_OFFSET_X+glb.TEXT_PADDING_X+textWidth1+glb.TEXT_PADDING_X+textWidth2,
			160, glb.COLOR_BLUE)
	}

	//drawText(MSG_3, 10, 200, 255, 255, 0)
	//drawText(MSG_4, 197, 20, 255, 255, 255)
	drawText(glb.MSG_5, 10, 200, glb.COLOR_YELLOW)
}
