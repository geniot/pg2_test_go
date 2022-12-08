package loop

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLoop struct {
}

func NewEventLoop() *EventLoop {
	return &EventLoop{}
}

func (eventLoop EventLoop) Run() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {

		case *sdl.KeyboardEvent:
			if t.Repeat > 0 {
				break
			}
			ctx.LastPressedKey = t.Keysym.Sym
			if t.State == sdl.PRESSED {
				ctx.PressedKeysCodes.Add(t.Keysym.Sym)
			} else { // if t.State == sdl.RELEASED {
				ctx.PressedKeysCodes.Remove(t.Keysym.Sym)
			}
			break

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_CLOSE {
				ctx.Window.SaveWindowState()
			}
			break

		case *sdl.QuitEvent:
			ctx.Loop.Stop()
			break
		}
	}
	ctx.Device.ProcessKeyActions()
}
