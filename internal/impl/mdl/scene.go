package mdl

type Scene struct {
	imageElements []*ImageElement
}

func NewScene() *Scene {
	imgElements := make([]*ImageElement, len(ButtonImages))
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
