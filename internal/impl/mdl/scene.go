package mdl

type Scene struct {
	imageElements []*ImageElement
}

func NewScene() *Scene {
	l := GetImageDescriptorsLength()
	imgElements := make([]*ImageElement, l)
	for i := 0; i < l; i++ {
		oX, oY, fN, dO := GetImageDescriptorPropsByIndex(i)
		iEl := NewImageElement(fN, oX, oY, dO)
		imgElements[i] = iEl
	}
	return &Scene{imgElements}
}

func (scene Scene) Render() {
	for i := range scene.imageElements {
		scene.imageElements[i].Render()
	}
}
