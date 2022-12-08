package glb

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
)

const (
	APP_NAME       = "PG2 Test"
	APP_VERSION    = "0.5"
	CONF_FILE_NAME = ".pg2_test.properties"
	FONT_FILE_NAME = "pixelberry.ttf"
	FONT_SIZE      = 8
	TEXT_OFFSET_X  = 10
	TEXT_PADDING_X = 4
	J_CROSS_WIDTH  = 7
	J_CROSS_HEIGHT = 7

	MAX_VOLTAGE = 4150000
	MIN_VOLTAGE = 3330000
	USB_VOLTAGE = 110000 //PG2V2
	//USB_VOLTAGE    = 80000 //PG2

	//joystick cross is bound to this rectangle
	SMALL_SCREEN_X1       = 124 + J_CROSS_WIDTH/2
	SMALL_SCREEN_X2       = 196 - J_CROSS_WIDTH/2
	SMALL_SCREEN_Y1       = 63 + J_CROSS_WIDTH/2
	SMALL_SCREEN_Y2       = 115 - J_CROSS_WIDTH/2
	SMALL_SCREEN_WIDTH    = SMALL_SCREEN_X2 - SMALL_SCREEN_X1
	SMALL_SCREEN_HEIGHT   = SMALL_SCREEN_Y2 - SMALL_SCREEN_Y1
	SMALL_SCREEN_X_CENTER = float32(SMALL_SCREEN_X1 + SMALL_SCREEN_WIDTH/2)
	SMALL_SCREEN_Y_CENTER = float32(SMALL_SCREEN_Y1 + SMALL_SCREEN_HEIGHT/2)

	MSG_0 = "Press L1 + START to exit."
	MSG_1 = "Press L1 + X to play a sound."
	MSG_2 = "Last detected key:"
	//MSG_3 = "Press POWER + R1 to de/activate mouse."
	//MSG_4 = "reading..."
	MSG_5 = "Press L2 + R2 to rumble."

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

	COLOR_RED    = sdl.Color{R: 192, G: 64, B: 64, A: 255}
	COLOR_GREEN  = sdl.Color{R: 64, G: 192, B: 64, A: 255}
	COLOR_GRAY   = sdl.Color{R: 192, G: 192, B: 192, A: 255}
	COLOR_WHITE  = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	COLOR_PURPLE = sdl.Color{R: 255, G: 0, B: 255, A: 255}
	COLOR_YELLOW = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	COLOR_BLUE   = sdl.Color{R: 0, G: 255, B: 255, A: 255}
)

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func Bytes(s uint64) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1000, sizes)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	//https://emptycharacter.com/
	f := "%.0f%s"
	//if val < 10 {
	//	f = "%.1f%s"
	//}

	return fmt.Sprintf(f, val, suffix)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}
