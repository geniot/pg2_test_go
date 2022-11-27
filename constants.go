package main

import "github.com/veandco/go-sdl2/sdl"

const (
	FONT_PATH              = "media/pixelberry.ttf"
	FONT_SIZE              = 14
	SCREEN_WIDTH           = 320
	SCREEN_HEIGHT          = 240
	SECOND_SCREEN_X_OFFSET = 2500 //used for testing/debugging

	GCW_BUTTON_UP    = sdl.K_UP
	GCW_BUTTON_DOWN  = sdl.K_DOWN
	GCW_BUTTON_LEFT  = sdl.K_LEFT
	GCW_BUTTON_RIGHT = sdl.K_RIGHT

	GCW_BUTTON_A = sdl.K_LCTRL
	GCW_BUTTON_B = sdl.K_LALT
	GCW_BUTTON_X = sdl.K_SPACE
	GCW_BUTTON_Y = sdl.K_LSHIFT

	GCW_BUTTON_L1 = sdl.K_TAB
	GCW_BUTTON_R1 = sdl.K_BACKSPACE

	GCW_BUTTON_L2 = sdl.K_RSHIFT
	GCW_BUTTON_R2 = sdl.K_RALT

	//GCW_BUTTON_L2 = sdl.K_PAGEUP
	//GCW_BUTTON_R2 = sdl.K_PAGEDOWN

	GCW_BUTTON_SELECT = sdl.K_ESCAPE
	GCW_BUTTON_START  = sdl.K_RETURN
	GCW_BUTTON_MENU   = sdl.K_RCTRL

	GCW_BUTTON_L3    = sdl.K_KP_DIVIDE
	GCW_BUTTON_R3    = sdl.K_KP_PERIOD
	GCW_BUTTON_POWER = sdl.K_HOME
)
