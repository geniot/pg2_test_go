package gui

type RenderLoop struct {
	application *ApplicationImpl
}

func NewRenderLoop(app *ApplicationImpl) *RenderLoop {
	return &RenderLoop{app}
}

func (renderLoop RenderLoop) Run() {
	window := renderLoop.application.window
	window.sdlRenderer.SetDrawColor(16, 16, 16, 255)
	window.sdlRenderer.Clear()
	renderLoop.application.scene.Render(window.sdlRenderer, renderLoop.application.loop.eventLoop.pressedKeysCodes)
	window.sdlRenderer.Present()
}
