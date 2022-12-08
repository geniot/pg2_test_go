package gui

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/glb"
	"geniot.com/geniot/pg2_test_go/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type Window struct {
	sdlWindow   *sdl.Window
	iconSurface *sdl.Surface
}

func NewWindow() *Window {
	w := Window{}

	xPos, yPos, width, height := ctx.Device.GetWindowPosAndSize()
	w.sdlWindow, _ = sdl.CreateWindow(
		glb.APP_NAME+" "+glb.APP_VERSION,
		xPos, yPos, width, height,
		ctx.Device.GetWindowState())

	w.iconSurface, _ = img.LoadRW(resources.GetResource("pg2test.png"), true)
	w.sdlWindow.SetIcon(w.iconSurface)

	ctx.Renderer, _ = sdl.CreateRenderer(w.sdlWindow, -1,
		sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_ACCELERATED)
	//sdl.RENDERER_ACCELERATED)
	//srf, _ := wnd.GetSurface()
	//wnd.UpdateSurface()
	//
	//srf.FillRect(&sdl.Rect{0, 0, srf.H, srf.W}, sdl.MapRGB(srf.Format, 16, 16, 16))
	//rnd.SetDrawColor(16, 16, 16, 255)
	//rnd.Clear()

	sdl.AddEventWatchFunc(w.resizingEventWatcher, nil)

	return &w
}

func (window Window) resizingEventWatcher(event sdl.Event, data interface{}) bool {
	switch t := event.(type) {
	case *sdl.WindowEvent:
		if t.Event == sdl.WINDOWEVENT_RESIZED {
			ctx.RenderLoop.Run()
		}
		break
	}
	return false
}

func (window Window) SaveWindowState() {
	width, height := window.sdlWindow.GetSize()
	xPos, yPos := window.sdlWindow.GetPosition()
	windowState := window.sdlWindow.GetFlags()
	ctx.Config.Set(glb.WINDOW_STATE_KEY, strconv.FormatInt(int64(windowState), 10))

	if windowState&sdl.WINDOW_MAXIMIZED <= 0 {
		ctx.Config.Set(glb.WINDOW_WIDTH_KEY, strconv.FormatInt(int64(width), 10))
		ctx.Config.Set(glb.WINDOW_HEIGHT_KEY, strconv.FormatInt(int64(height), 10))
		ctx.Config.Set(glb.WINDOW_XPOS_KEY, strconv.FormatInt(int64(xPos), 10))
		ctx.Config.Set(glb.WINDOW_YPOS_KEY, strconv.FormatInt(int64(yPos), 10))
	}

	ctx.Config.Save()
}
