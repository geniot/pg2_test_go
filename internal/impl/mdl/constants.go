package mdl

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	APP_NAME       = "PG2 Test"
	APP_VERSION    = "0.5"
	CONF_FILE_NAME = ".pg2_test.properties"
	FONT_FILE_NAME = "pixelberry.ttf"
	FONT_SIZE      = 8

	WINDOW_XPOS_KEY   = "WINDOW_XPOS_KEY"
	WINDOW_YPOS_KEY   = "WINDOW_YPOS_KEY"
	WINDOW_WIDTH_KEY  = "WINDOW_WIDTH_KEY"
	WINDOW_HEIGHT_KEY = "WINDOW_HEIGHT_KEY"
	WINDOW_STATE_KEY  = "WINDOW_STATE_KEY"

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

	//GCW_BUTTON_L2 = sdl.K_RSHIFT
	//GCW_BUTTON_R2 = sdl.K_RALT

	GCW_BUTTON_L2 = sdl.K_PAGEUP
	GCW_BUTTON_R2 = sdl.K_PAGEDOWN

	GCW_BUTTON_SELECT = sdl.K_ESCAPE
	GCW_BUTTON_START  = sdl.K_RETURN
	GCW_BUTTON_MENU   = sdl.K_HOME

	GCW_VOLUMEUP   = sdl.K_VOLUMEUP
	GCW_VOLUMEDOWN = sdl.K_VOLUMEDOWN

	GCW_BUTTON_L3 = sdl.K_KP_DIVIDE
	//GCW_BUTTON_R3    = sdl.K_KP_PERIOD
	//GCW_BUTTON_POWER = sdl.K_HOME
)

var (
	KeyNames = map[sdl.Keycode]string{
		GCW_BUTTON_UP:    "UP",
		GCW_BUTTON_DOWN:  "DOWN",
		GCW_BUTTON_LEFT:  "LEFT",
		GCW_BUTTON_RIGHT: "RIGHT",

		GCW_BUTTON_A: "LCTRL",
		GCW_BUTTON_B: "LALT",
		GCW_BUTTON_X: "SPACE",
		GCW_BUTTON_Y: "LSHIFT",

		GCW_BUTTON_L1: "TAB",
		GCW_BUTTON_R1: "BACKSPACE",

		GCW_BUTTON_L2: "PAGEUP",
		GCW_BUTTON_R2: "PAGEDOWN",

		GCW_BUTTON_SELECT: "ESCAPE",
		GCW_BUTTON_START:  "RETURN",
		GCW_BUTTON_MENU:   "HOME",

		GCW_VOLUMEUP:   "VOLUMEUP",
		GCW_VOLUMEDOWN: "VOLUMEDOWN",
	}
)

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}
