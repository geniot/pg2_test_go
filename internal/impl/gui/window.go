package gui

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/imm"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type Window struct {
	application *ApplicationImpl
	sdlWindow   *sdl.Window
	sdlRenderer *sdl.Renderer
}

func NewWindow(app *ApplicationImpl) *Window {
	wnd, _ := sdl.CreateWindow(
		imm.APP_NAME+" "+imm.APP_VERSION,
		int32(ctx.Config.Get(imm.WINDOW_XPOS_KEY)),
		int32(ctx.Config.Get(imm.WINDOW_YPOS_KEY)),
		int32(ctx.Config.Get(imm.WINDOW_WIDTH_KEY)),
		int32(ctx.Config.Get(imm.WINDOW_HEIGHT_KEY)),
		ctx.Config.Get(imm.WINDOW_STATE_KEY))

	rnd, _ := sdl.CreateRenderer(wnd, -1,
		sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_ACCELERATED)
	//sdl.RENDERER_ACCELERATED)
	//srf, _ := wnd.GetSurface()
	//wnd.UpdateSurface()
	//
	//srf.FillRect(&sdl.Rect{0, 0, srf.H, srf.W}, sdl.MapRGB(srf.Format, 16, 16, 16))
	//rnd.SetDrawColor(16, 16, 16, 255)
	//rnd.Clear()

	w := Window{app, wnd, rnd}

	sdl.AddEventWatchFunc(w.resizingEventWatcher, nil)

	return &w
}

func (window Window) resizingEventWatcher(event sdl.Event, data interface{}) bool {
	switch t := event.(type) {
	case *sdl.WindowEvent:
		if t.Event == sdl.WINDOWEVENT_RESIZED {
			window.application.loop.renderLoop.Run()
		}
		break
	}
	return false
}

func (window Window) OnBeforeClose() {
	window.SaveWindowState()
}

func (window Window) SaveWindowState() {
	width, height := window.sdlWindow.GetSize()
	xPos, yPos := window.sdlWindow.GetPosition()
	windowState := window.sdlWindow.GetFlags()
	ctx.Config.Set(imm.WINDOW_STATE_KEY, strconv.FormatInt(int64(windowState), 10))

	if windowState&sdl.WINDOW_MAXIMIZED <= 0 {
		ctx.Config.Set(imm.WINDOW_WIDTH_KEY, strconv.FormatInt(int64(width), 10))
		ctx.Config.Set(imm.WINDOW_HEIGHT_KEY, strconv.FormatInt(int64(height), 10))
		ctx.Config.Set(imm.WINDOW_XPOS_KEY, strconv.FormatInt(int64(xPos), 10))
		ctx.Config.Set(imm.WINDOW_YPOS_KEY, strconv.FormatInt(int64(yPos), 10))
	}

	ctx.Config.Save()
}
