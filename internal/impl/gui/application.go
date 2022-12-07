package gui

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/dev"
	"geniot.com/geniot/pg2_test_go/internal/impl/gui/loop"
	"geniot.com/geniot/pg2_test_go/internal/impl/mdl"
	"geniot.com/geniot/pg2_test_go/resources"
	"github.com/veandco/go-sdl2/mix"
)

type ApplicationImpl struct {
	audioChunk *mix.Chunk
}

func NewApplication() *ApplicationImpl {
	return &ApplicationImpl{}
}

func (app *ApplicationImpl) Start() {
	ctx.Config = NewConfig()
	ctx.Device = dev.NewDevice()
	ctx.Window = NewWindow()

	ctx.Loop = loop.NewLoop()
	ctx.EventLoop = loop.NewEventLoop()
	ctx.PhysicsLoop = loop.NewPhysicsLoop()
	ctx.RenderLoop = loop.NewRenderLoop()

	ctx.CurrentScene = mdl.NewScene()

	app.audioChunk, _ = mix.LoadWAVRW(resources.GetResource("tone.wav"), true)

	ctx.Loop.Start()
}

func (app *ApplicationImpl) PlaySound() {
	app.audioChunk.Play(1, 0)
}
