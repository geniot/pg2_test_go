package loop

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLoop struct {
	lastPressedKey sdl.Keycode
}

func NewEventLoop() *EventLoop {
	return &EventLoop{sdl.K_UNKNOWN}
}

func (eventLoop EventLoop) Run() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {

		case *sdl.KeyboardEvent:
			if t.Repeat > 0 {
				break
			}
			eventLoop.lastPressedKey = t.Keysym.Sym
			if t.State == sdl.PRESSED {
				ctx.PressedKeysCodes.Add(eventLoop.lastPressedKey)
			} else { // if t.State == sdl.RELEASED {
				ctx.PressedKeysCodes.Remove(eventLoop.lastPressedKey)
			}
			break

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_CLOSE {
				ctx.Window.SaveWindowState()
			}
			break

		case *sdl.QuitEvent:
			ctx.Application.Stop()
			break
		}
	}
	ctx.Device.ProcessKeyActions()
}
