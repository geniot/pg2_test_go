package loop

import "geniot.com/geniot/pg2_test_go/internal/ctx"

type RenderLoop struct {
}

func NewRenderLoop() *RenderLoop {
	return &RenderLoop{}
}

func (renderLoop RenderLoop) Run() {
	ctx.Renderer.SetDrawColor(16, 16, 16, 255)
	ctx.Renderer.Clear()
	ctx.CurrentScene.Render()
	ctx.Renderer.Present()
}
