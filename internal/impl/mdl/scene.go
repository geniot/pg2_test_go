package mdl

import "geniot.com/geniot/pg2_test_go/internal/api"

type Scene struct {
	imageElements []api.IRenderable
}

func NewScene() *Scene {
	//totalLength := len(ButtonImages) + len(JoystickImages)
	imgElements := make([]api.IRenderable, len(ButtonImages))
	for i := range ButtonImages {
		iEl := NewImageElement(
			ButtonImages[i].ImageName,
			ButtonImages[i].OffsetX,
			ButtonImages[i].OffsetY,
			ButtonImages[i].DisplayOnPress)
		imgElements[i] = iEl
	}
	return &Scene{imgElements}
}

func (scene Scene) Render() {
	for i := range scene.imageElements {
		scene.imageElements[i].Render()
	}
}
