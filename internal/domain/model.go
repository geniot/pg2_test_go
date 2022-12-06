package domain

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type Model struct {
	IsRumbleSupported bool
	AudioChunk        *mix.Chunk
	Haptic            *sdl.Haptic
	Joystick          *sdl.Joystick
}

func NewModel() *Model {
	return &Model{false,
		nil, nil, nil}
}
