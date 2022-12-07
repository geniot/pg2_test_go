package gui

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/dev"
	"geniot.com/geniot/pg2_test_go/internal/impl/gui/loop"
	"geniot.com/geniot/pg2_test_go/internal/impl/mdl"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type ApplicationImpl struct {
	IsRumbleSupported bool
	AudioChunk        *mix.Chunk
	Haptic            *sdl.Haptic
	Joystick          *sdl.Joystick
}

func NewApplication() *ApplicationImpl {
	return &ApplicationImpl{}
}

func (app ApplicationImpl) Start() {
	sdl.Init(sdl.INIT_EVERYTHING)

	ctx.Config = NewConfig()
	ctx.Device = dev.NewDesktopDevice()
	ctx.Window = NewWindow()

	ctx.Loop = loop.NewLoop()
	ctx.EventLoop = loop.NewEventLoop()
	ctx.PhysicsLoop = loop.NewPhysicsLoop()
	ctx.RenderLoop = loop.NewRenderLoop()

	ctx.CurrentScene = mdl.NewScene()

	ctx.Loop.Start()
}

func (app ApplicationImpl) PlaySound() {
	app.AudioChunk.Play(1, 0)
}
