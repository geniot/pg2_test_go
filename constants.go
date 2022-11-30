package main

import "github.com/veandco/go-sdl2/sdl"

const (
	FONT_PATH      = "media/pixelberry.ttf"
	FONT_SIZE      = 8
	SCREEN_WIDTH   = 320
	SCREEN_HEIGHT  = 240
	TEXT_OFFSET_X  = 10
	TEXT_PADDING_X = 4
	J_CROSS_WIDTH  = 7

	MAX_VOLTAGE = 4150000
	MIN_VOLTAGE = 3330000
	USB_VOLTAGE = 110000 //PG2V2
	//USB_VOLTAGE    = 80000 //PG2

	//joystick cross is bound to this rectangle
	SMALL_SCREEN_X1       = 123
	SMALL_SCREEN_X2       = 197
	SMALL_SCREEN_Y1       = 62
	SMALL_SCREEN_Y2       = 116
	SMALL_SCREEN_WIDTH    = SMALL_SCREEN_X2 - SMALL_SCREEN_X1
	SMALL_SCREEN_HEIGHT   = SMALL_SCREEN_Y2 - SMALL_SCREEN_Y1
	SMALL_SCREEN_X_CENTER = float32(SMALL_SCREEN_X1 + SMALL_SCREEN_WIDTH/2)
	SMALL_SCREEN_Y_CENTER = float32(SMALL_SCREEN_Y1 + SMALL_SCREEN_HEIGHT/2)

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

	//GCW_BUTTON_L2 = sdl.K_RSHIFT
	//GCW_BUTTON_R2 = sdl.K_RALT

	GCW_BUTTON_L2 = sdl.K_PAGEUP
	GCW_BUTTON_R2 = sdl.K_PAGEDOWN

	GCW_BUTTON_SELECT = sdl.K_ESCAPE
	GCW_BUTTON_START  = sdl.K_RETURN
	GCW_BUTTON_MENU   = sdl.K_HOME

	GCW_BUTTON_L3 = sdl.K_KP_DIVIDE
	//GCW_BUTTON_R3    = sdl.K_KP_PERIOD
	//GCW_BUTTON_POWER = sdl.K_HOME

	MSG_0 = "Press L1 + START to exit."
	MSG_1 = "Press L1 + X to play a sound."
	MSG_2 = "Last detected key:"
	//MSG_3 = "Press POWER + R1 to de/activate mouse."
	MSG_4 = "reading..."
	MSG_5 = "Press L2 + R2 to rumble."
)

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}
