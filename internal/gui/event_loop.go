package gui

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type EventLoop struct {
	application      *Application
	pressedKeysCodes mapset.Set[sdl.Keycode]
	lastPressedKey   sdl.Keycode
}

func NewEventLoop(app *Application) *EventLoop {
	return &EventLoop{app, mapset.NewSet[sdl.Keycode](), sdl.K_UNKNOWN}
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
				eventLoop.pressedKeysCodes.Add(eventLoop.lastPressedKey)
			} else { // if t.State == sdl.RELEASED {
				eventLoop.pressedKeysCodes.Remove(eventLoop.lastPressedKey)
			}
			break

		case *sdl.WindowEvent:
			if t.Event == sdl.WINDOWEVENT_CLOSE {
				eventLoop.application.window.OnBeforeClose()
			}
			break

		case *sdl.QuitEvent:
			eventLoop.application.loop.isRunning.UnSet()
			break
		}
	}
}
