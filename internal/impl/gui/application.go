package gui

import (
	"fmt"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/dev"
	"geniot.com/geniot/pg2_test_go/internal/impl/gui/loop"
	"geniot.com/geniot/pg2_test_go/internal/impl/mdl"
	"geniot.com/geniot/pg2_test_go/resources"
	"github.com/robfig/cron/v3"
	"github.com/veandco/go-sdl2/mix"
)

type ApplicationImpl struct {
	audioChunk *mix.Chunk
}

func (app *ApplicationImpl) Stop() {
	ctx.Device.Stop()
	ctx.Loop.Stop()
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

	go ctx.Device.UpdateBatteryStatus()
	go ctx.Device.UpdateDiskStatus()
	go ctx.Device.UpdateVolume()

	c := cron.New()
	_, err := c.AddFunc("@every 1s", ctx.Device.UpdateBatteryStatus)
	_, err = c.AddFunc("@every 1s", ctx.Device.UpdateDiskStatus)
	_, err = c.AddFunc("@every 1s", ctx.Device.UpdateVolume)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.Start()

	ctx.Loop.Start()
}

func (app *ApplicationImpl) PlaySound() {
	app.audioChunk.Play(1, 0)
}
