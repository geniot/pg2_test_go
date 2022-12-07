package gui

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/imm"
	"github.com/veandco/go-sdl2/mix"
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
				ctx.Window.OnBeforeClose()
			}
			break

		case *sdl.QuitEvent:
			ctx.Loop.Stop()
			break
		}
	}
	eventLoop.processKeyActions()
}

func (eventLoop EventLoop) processKeyActions() {
	if ctx.PressedKeysCodes.Contains(sdl.K_q) ||
		(ctx.PressedKeysCodes.Contains(imm.GCW_BUTTON_L1) &&
			ctx.PressedKeysCodes.Contains(imm.GCW_BUTTON_START)) {
		ctx.Loop.Stop()
	}
	if ctx.PressedKeysCodes.Contains(imm.GCW_BUTTON_L1) &&
		ctx.PressedKeysCodes.Contains(imm.GCW_BUTTON_X) &&
		mix.Playing(-1) != 1 {
		ctx.Application.PlaySound()
	}
	if ctx.Device.IsRumbleSupported() {
		if ctx.PressedKeysCodes.Contains(imm.GCW_BUTTON_L2) &&
			ctx.PressedKeysCodes.Contains(imm.GCW_BUTTON_R2) {
			ctx.Device.Rumble()
		}
	}
}
