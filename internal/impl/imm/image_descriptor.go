package imm

import (
	"github.com/veandco/go-sdl2/sdl"
)

func GetImageDescriptorsLength() int {
	return len(imageDescriptors)
}

func GetImageDescriptorPropsByIndex(index int) (int32, int32, string, sdl.Keycode) {
	return imageDescriptors[index].offsetX,
		imageDescriptors[index].offsetY,
		imageDescriptors[index].imageName,
		imageDescriptors[index].displayOnPress

}

type imageDescriptor struct {
	offsetX        int32
	offsetY        int32
	imageName      string
	displayOnPress sdl.Keycode
}

var (
	imageDescriptors = []imageDescriptor{
		//background
		{
			offsetX:        90,
			offsetY:        50,
			imageName:      "pg2_back.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		//VOLUME DOWN
		{
			offsetX:        169,
			offsetY:        53,
			imageName:      "pg2_button_vol1.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        169,
			offsetY:        53,
			imageName:      "pg2_button_vol1_pressed.png",
			displayOnPress: GCW_VOLUMEDOWN,
		},
		{
			offsetX:        156,
			offsetY:        40,
			imageName:      "info_voldw.png",
			displayOnPress: GCW_VOLUMEDOWN,
		},
		//VOLUME UP
		{
			offsetX:        169,
			offsetY:        53,
			imageName:      "pg2_button_vol2.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        169,
			offsetY:        53,
			imageName:      "pg2_button_vol2_pressed.png",
			displayOnPress: GCW_VOLUMEUP,
		},
		{
			offsetX:        177,
			offsetY:        40,
			imageName:      "info_volup.png",
			displayOnPress: GCW_VOLUMEUP,
		},
		//A
		{
			offsetX:        216,
			offsetY:        77,
			imageName:      "pg2_button_a.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        216,
			offsetY:        77,
			imageName:      "pg2_button_a_pressed.png",
			displayOnPress: GCW_BUTTON_A,
		},
		{
			offsetX:        226,
			offsetY:        73,
			imageName:      "info_btna.png",
			displayOnPress: GCW_BUTTON_A,
		},
		//B
		{
			offsetX:        209,
			offsetY:        86,
			imageName:      "pg2_button_b.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        209,
			offsetY:        86,
			imageName:      "pg2_button_b_pressed.png",
			displayOnPress: GCW_BUTTON_B,
		},
		{
			offsetX:        218,
			offsetY:        90,
			imageName:      "info_btnb.png",
			displayOnPress: GCW_BUTTON_B,
		},
		//X
		{
			offsetX:        209,
			offsetY:        70,
			imageName:      "pg2_button_x.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        209,
			offsetY:        70,
			imageName:      "pg2_button_x_pressed.png",
			displayOnPress: GCW_BUTTON_X,
		},
		{
			offsetX:        218,
			offsetY:        64,
			imageName:      "info_btnx.png",
			displayOnPress: GCW_BUTTON_X,
		},
		//Y
		{
			offsetX:        201,
			offsetY:        78,
			imageName:      "pg2_button_y.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        201,
			offsetY:        78,
			imageName:      "pg2_button_y_pressed.png",
			displayOnPress: GCW_BUTTON_Y,
		},
		{
			offsetX:        211,
			offsetY:        82,
			imageName:      "info_btny.png",
			displayOnPress: GCW_BUTTON_Y,
		},
		//UP
		{
			offsetX:        102,
			offsetY:        70,
			imageName:      "pg2_button_up.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        102,
			offsetY:        70,
			imageName:      "pg2_button_up_pressed.png",
			displayOnPress: GCW_BUTTON_UP,
		},
		{
			offsetX:        71,
			offsetY:        64,
			imageName:      "info_padup.png",
			displayOnPress: GCW_BUTTON_UP,
		},
		//DOWN
		{
			offsetX:        102,
			offsetY:        84,
			imageName:      "pg2_button_down.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        102,
			offsetY:        84,
			imageName:      "pg2_button_down_pressed.png",
			displayOnPress: GCW_BUTTON_DOWN,
		},
		{
			offsetX:        59,
			offsetY:        92,
			imageName:      "info_paddown.png",
			displayOnPress: GCW_BUTTON_DOWN,
		},
		//LEFT
		{
			offsetX:        95,
			offsetY:        77,
			imageName:      "pg2_button_left.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        95,
			offsetY:        77,
			imageName:      "pg2_button_left_pressed.png",
			displayOnPress: GCW_BUTTON_LEFT,
		},
		{
			offsetX:        63,
			offsetY:        73,
			imageName:      "info_padleft.png",
			displayOnPress: GCW_BUTTON_LEFT,
		},
		//RIGHT
		{
			offsetX:        109,
			offsetY:        77,
			imageName:      "pg2_button_right.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        109,
			offsetY:        77,
			imageName:      "pg2_button_right_pressed.png",
			displayOnPress: GCW_BUTTON_RIGHT,
		},
		{
			offsetX:        58,
			offsetY:        81,
			imageName:      "info_padright.png",
			displayOnPress: GCW_BUTTON_RIGHT,
		},
		//MENU
		{
			offsetX:        200,
			offsetY:        96,
			imageName:      "pg2_button_s.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        200,
			offsetY:        96,
			imageName:      "pg2_button_s_pressed.png",
			displayOnPress: GCW_BUTTON_MENU,
		},
		{
			offsetX:        206,
			offsetY:        98,
			imageName:      "info_menu.png",
			displayOnPress: GCW_BUTTON_MENU,
		},
		//SELECT
		{
			offsetX:        200,
			offsetY:        105,
			imageName:      "pg2_button_s.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        200,
			offsetY:        105,
			imageName:      "pg2_button_s_pressed.png",
			displayOnPress: GCW_BUTTON_SELECT,
		},
		{
			offsetX:        206,
			offsetY:        107,
			imageName:      "info_select.png",
			displayOnPress: GCW_BUTTON_SELECT,
		},
		//START
		{
			offsetX:        200,
			offsetY:        114,
			imageName:      "pg2_button_s.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        200,
			offsetY:        114,
			imageName:      "pg2_button_s_pressed.png",
			displayOnPress: GCW_BUTTON_START,
		},
		{
			offsetX:        206,
			offsetY:        116,
			imageName:      "info_start.png",
			displayOnPress: GCW_BUTTON_START,
		},
		//L1
		{
			offsetX:        92,
			offsetY:        55,
			imageName:      "pg2_button_l1.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        92,
			offsetY:        55,
			imageName:      "pg2_button_l1_pressed.png",
			displayOnPress: GCW_BUTTON_L1,
		},
		{
			offsetX:        86,
			offsetY:        40,
			imageName:      "info_btnl1.png",
			displayOnPress: GCW_BUTTON_L1,
		},
		//L2
		{
			offsetX:        110,
			offsetY:        53,
			imageName:      "pg2_button_l2.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        110,
			offsetY:        53,
			imageName:      "pg2_button_l2_pressed.png",
			displayOnPress: GCW_BUTTON_L2,
		},
		{
			offsetX:        109,
			offsetY:        40,
			imageName:      "info_btnl2.png",
			displayOnPress: GCW_BUTTON_L2,
		},
		//R1
		{
			offsetX:        213,
			offsetY:        55,
			imageName:      "pg2_button_r1.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        213,
			offsetY:        55,
			imageName:      "pg2_button_r1_pressed.png",
			displayOnPress: GCW_BUTTON_R1,
		},
		{
			offsetX:        213,
			offsetY:        40,
			imageName:      "info_btnr1.png",
			displayOnPress: GCW_BUTTON_R1,
		},
		//R2
		{
			offsetX:        199,
			offsetY:        53,
			imageName:      "pg2_button_r2.png",
			displayOnPress: sdl.K_UNKNOWN,
		},
		{
			offsetX:        199,
			offsetY:        53,
			imageName:      "pg2_button_r2_pressed.png",
			displayOnPress: GCW_BUTTON_R2,
		},
		{
			offsetX:        200,
			offsetY:        40,
			imageName:      "info_btnr2.png",
			displayOnPress: GCW_BUTTON_R2,
		},
	}
)
