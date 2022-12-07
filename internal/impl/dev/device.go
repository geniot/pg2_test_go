package dev

import (
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/mdl"
	"geniot.com/geniot/pg2_test_go/resources"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"runtime"
)

func NewDevice() api.IDevice {
	if runtime.GOARCH == "mips" {
		return NewHandheldDevice()
	} else {
		return NewDesktopDevice()
	}
}

func initCommon() {
	err := ttf.Init()
	if err != nil {
		panic(err)
	}
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		println(err.Error())
	}
	ctx.Font, err = ttf.OpenFontRW(resources.GetResource(mdl.FONT_FILE_NAME), 1, mdl.FONT_SIZE)
	if err != nil {
		println(err.Error())
	}
}
