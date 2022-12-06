package utils

import (
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"path/filepath"
)

const (
	APP_NAME       = "PG2 Test"
	APP_VERSION    = "0.1"
	CONF_FILE_NAME = ".pg2_test.properties"

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

type ImageDescriptor struct {
	OffsetX        int32
	OffsetY        int32
	ImageName      string
	DisplayOnPress sdl.Keycode
}

var (
	HOME_DIR, _    = os.UserHomeDir()
	PATH_TO_CONFIG = filepath.Join(HOME_DIR, CONF_FILE_NAME)
	WINDOW_TITLE   = APP_NAME + " " + APP_VERSION

	IMAGE_DESCRIPTORS = []ImageDescriptor{
		//background
		{
			OffsetX:        90,
			OffsetY:        50,
			ImageName:      "pg2_back.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		//VOLUME DOWN
		{
			OffsetX:        169,
			OffsetY:        53,
			ImageName:      "pg2_button_vol1.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        169,
			OffsetY:        53,
			ImageName:      "pg2_button_vol1_pressed.png",
			DisplayOnPress: GCW_VOLUMEDOWN,
		},
		{
			OffsetX:        156,
			OffsetY:        40,
			ImageName:      "info_voldw.png",
			DisplayOnPress: GCW_VOLUMEDOWN,
		},
		//VOLUME UP
		{
			OffsetX:        169,
			OffsetY:        53,
			ImageName:      "pg2_button_vol2.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        169,
			OffsetY:        53,
			ImageName:      "pg2_button_vol2_pressed.png",
			DisplayOnPress: GCW_VOLUMEUP,
		},
		{
			OffsetX:        177,
			OffsetY:        40,
			ImageName:      "info_volup.png",
			DisplayOnPress: GCW_VOLUMEUP,
		},
		//A
		{
			OffsetX:        216,
			OffsetY:        77,
			ImageName:      "pg2_button_a.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        216,
			OffsetY:        77,
			ImageName:      "pg2_button_a_pressed.png",
			DisplayOnPress: GCW_BUTTON_A,
		},
		{
			OffsetX:        226,
			OffsetY:        73,
			ImageName:      "info_btna.png",
			DisplayOnPress: GCW_BUTTON_A,
		},
		//B
		{
			OffsetX:        209,
			OffsetY:        86,
			ImageName:      "pg2_button_b.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        209,
			OffsetY:        86,
			ImageName:      "pg2_button_b_pressed.png",
			DisplayOnPress: GCW_BUTTON_B,
		},
		{
			OffsetX:        218,
			OffsetY:        90,
			ImageName:      "info_btnb.png",
			DisplayOnPress: GCW_BUTTON_B,
		},
		//X
		{
			OffsetX:        209,
			OffsetY:        70,
			ImageName:      "pg2_button_x.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        209,
			OffsetY:        70,
			ImageName:      "pg2_button_x_pressed.png",
			DisplayOnPress: GCW_BUTTON_X,
		},
		{
			OffsetX:        218,
			OffsetY:        64,
			ImageName:      "info_btnx.png",
			DisplayOnPress: GCW_BUTTON_X,
		},
		//Y
		{
			OffsetX:        201,
			OffsetY:        78,
			ImageName:      "pg2_button_y.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        201,
			OffsetY:        78,
			ImageName:      "pg2_button_y_pressed.png",
			DisplayOnPress: GCW_BUTTON_Y,
		},
		{
			OffsetX:        211,
			OffsetY:        82,
			ImageName:      "info_btny.png",
			DisplayOnPress: GCW_BUTTON_Y,
		},
		//UP
		{
			OffsetX:        102,
			OffsetY:        70,
			ImageName:      "pg2_button_up.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        102,
			OffsetY:        70,
			ImageName:      "pg2_button_up_pressed.png",
			DisplayOnPress: GCW_BUTTON_UP,
		},
		{
			OffsetX:        71,
			OffsetY:        64,
			ImageName:      "info_padup.png",
			DisplayOnPress: GCW_BUTTON_UP,
		},
		//DOWN
		{
			OffsetX:        102,
			OffsetY:        84,
			ImageName:      "pg2_button_down.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        102,
			OffsetY:        84,
			ImageName:      "pg2_button_down_pressed.png",
			DisplayOnPress: GCW_BUTTON_DOWN,
		},
		{
			OffsetX:        59,
			OffsetY:        92,
			ImageName:      "info_paddown.png",
			DisplayOnPress: GCW_BUTTON_DOWN,
		},
		//LEFT
		{
			OffsetX:        95,
			OffsetY:        77,
			ImageName:      "pg2_button_left.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        95,
			OffsetY:        77,
			ImageName:      "pg2_button_left_pressed.png",
			DisplayOnPress: GCW_BUTTON_LEFT,
		},
		{
			OffsetX:        63,
			OffsetY:        73,
			ImageName:      "info_padleft.png",
			DisplayOnPress: GCW_BUTTON_LEFT,
		},
		//RIGHT
		{
			OffsetX:        109,
			OffsetY:        77,
			ImageName:      "pg2_button_right.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        109,
			OffsetY:        77,
			ImageName:      "pg2_button_right_pressed.png",
			DisplayOnPress: GCW_BUTTON_RIGHT,
		},
		{
			OffsetX:        58,
			OffsetY:        81,
			ImageName:      "info_padright.png",
			DisplayOnPress: GCW_BUTTON_RIGHT,
		},
		//MENU
		{
			OffsetX:        200,
			OffsetY:        96,
			ImageName:      "pg2_button_s.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        200,
			OffsetY:        96,
			ImageName:      "pg2_button_s_pressed.png",
			DisplayOnPress: GCW_BUTTON_MENU,
		},
		{
			OffsetX:        206,
			OffsetY:        98,
			ImageName:      "info_menu.png",
			DisplayOnPress: GCW_BUTTON_MENU,
		},
		//SELECT
		{
			OffsetX:        200,
			OffsetY:        105,
			ImageName:      "pg2_button_s.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        200,
			OffsetY:        105,
			ImageName:      "pg2_button_s_pressed.png",
			DisplayOnPress: GCW_BUTTON_SELECT,
		},
		{
			OffsetX:        206,
			OffsetY:        107,
			ImageName:      "info_select.png",
			DisplayOnPress: GCW_BUTTON_SELECT,
		},
		//START
		{
			OffsetX:        200,
			OffsetY:        114,
			ImageName:      "pg2_button_s.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        200,
			OffsetY:        114,
			ImageName:      "pg2_button_s_pressed.png",
			DisplayOnPress: GCW_BUTTON_START,
		},
		{
			OffsetX:        206,
			OffsetY:        116,
			ImageName:      "info_start.png",
			DisplayOnPress: GCW_BUTTON_START,
		},
		//L1
		{
			OffsetX:        92,
			OffsetY:        55,
			ImageName:      "pg2_button_l1.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        92,
			OffsetY:        55,
			ImageName:      "pg2_button_l1_pressed.png",
			DisplayOnPress: GCW_BUTTON_L1,
		},
		{
			OffsetX:        86,
			OffsetY:        40,
			ImageName:      "info_btnl1.png",
			DisplayOnPress: GCW_BUTTON_L1,
		},
		//L2
		{
			OffsetX:        110,
			OffsetY:        53,
			ImageName:      "pg2_button_l2.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        110,
			OffsetY:        53,
			ImageName:      "pg2_button_l2_pressed.png",
			DisplayOnPress: GCW_BUTTON_L2,
		},
		{
			OffsetX:        109,
			OffsetY:        40,
			ImageName:      "info_btnl2.png",
			DisplayOnPress: GCW_BUTTON_L2,
		},
		//R1
		{
			OffsetX:        213,
			OffsetY:        55,
			ImageName:      "pg2_button_r1.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        213,
			OffsetY:        55,
			ImageName:      "pg2_button_r1_pressed.png",
			DisplayOnPress: GCW_BUTTON_R1,
		},
		{
			OffsetX:        213,
			OffsetY:        40,
			ImageName:      "info_btnr1.png",
			DisplayOnPress: GCW_BUTTON_R1,
		},
		//R2
		{
			OffsetX:        199,
			OffsetY:        53,
			ImageName:      "pg2_button_r2.png",
			DisplayOnPress: sdl.K_UNKNOWN,
		},
		{
			OffsetX:        199,
			OffsetY:        53,
			ImageName:      "pg2_button_r2_pressed.png",
			DisplayOnPress: GCW_BUTTON_R2,
		},
		{
			OffsetX:        200,
			OffsetY:        40,
			ImageName:      "info_btnr2.png",
			DisplayOnPress: GCW_BUTTON_R2,
		},
	}
)
