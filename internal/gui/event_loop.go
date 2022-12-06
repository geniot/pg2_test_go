package gui

import (
	"geniot.com/geniot/pg2_test_go/internal/imm"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/mix"
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
	eventLoop.processKeyActions()
}

func (eventLoop EventLoop) processKeyActions() {
	if eventLoop.pressedKeysCodes.Contains(sdl.K_q) ||
		(eventLoop.pressedKeysCodes.Contains(imm.GCW_BUTTON_L1) &&
			eventLoop.pressedKeysCodes.Contains(imm.GCW_BUTTON_START)) {
		eventLoop.application.loop.isRunning.UnSet()
	}
	if eventLoop.pressedKeysCodes.Contains(imm.GCW_BUTTON_L1) &&
		eventLoop.pressedKeysCodes.Contains(imm.GCW_BUTTON_X) &&
		mix.Playing(-1) != 1 {
		_, err := eventLoop.application.model.AudioChunk.Play(1, 0)
		if err != nil {
			panic(err)
		}
	}
	if eventLoop.application.model.IsRumbleSupported {
		if eventLoop.pressedKeysCodes.Contains(imm.GCW_BUTTON_L2) &&
			eventLoop.pressedKeysCodes.Contains(imm.GCW_BUTTON_R2) {
			var err = eventLoop.application.model.Haptic.RumblePlay(0.33, 500)
			if err != nil {
				println(err.Error())
			}
		}
	}
}
