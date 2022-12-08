package gui

import (
	"fmt"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/glb"
	"geniot.com/geniot/pg2_test_go/internal/impl/dev"
	"geniot.com/geniot/pg2_test_go/internal/impl/gui/loop"
	"geniot.com/geniot/pg2_test_go/internal/impl/gui/rnd"
	"geniot.com/geniot/pg2_test_go/resources"
	"github.com/robfig/cron/v3"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/ttf"
)

type ApplicationImpl struct {
	audioChunk *mix.Chunk
	cron       *cron.Cron
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

	ctx.CurrentScene = rnd.NewScene()

	ctx.Font, _ = ttf.OpenFontRW(resources.GetResource(glb.FONT_FILE_NAME), 1, glb.FONT_SIZE)
	app.audioChunk, _ = mix.LoadWAVRW(resources.GetResource("tone.wav"), true)

	go ctx.Device.UpdateBatteryStatus()
	go ctx.Device.UpdateDiskStatus()
	go ctx.Device.UpdateVolume()

	app.cron = cron.New()
	_, err := app.cron.AddFunc("@every 1s", ctx.Device.UpdateBatteryStatus)
	_, err = app.cron.AddFunc("@every 1s", ctx.Device.UpdateDiskStatus)
	_, err = app.cron.AddFunc("@every 1s", ctx.Device.UpdateVolume)
	if err != nil {
		fmt.Println(err.Error())
	}
	app.cron.Start()

	ctx.Loop.Start()

	//graceful shutdown :) we let the loop finish all rendering/processing
	app.Stop()
}

func (app *ApplicationImpl) PlaySound() {
	app.audioChunk.Play(1, 0)
}

func (app *ApplicationImpl) Stop() {
	app.cron.Stop()
	app.audioChunk.Free()
	ctx.Font.Close()
	ctx.Device.Stop()
}
