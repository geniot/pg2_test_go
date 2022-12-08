package dev

import (
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/glb"
	"github.com/pydio/minio-srv/pkg/disk"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"runtime"
	"strings"
)

func NewDevice() api.IDevice {
	if strings.Index(runtime.GOARCH, "mips") == 0 {
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
}

func closeCommon() {
	ttf.Quit()
	sdl.Quit()
	mix.CloseAudio()
}

func updateDiskInfo(path string, diskInfo *api.DiskInfo) {
	di, err := disk.GetInfo(path)
	if err != nil {
		println(err.Error())
		diskInfo.IsDiskAvailable = false
		return
	}
	//we only do this once: when initializing or the disk has been inserted
	if !diskInfo.IsDiskAvailable {
		diskInfo.IsDiskAvailable = true
		diskInfo.MaxDiskSpace = glb.Bytes(di.Total)
		diskInfo.FreeDiskSpace = glb.Bytes(di.Free)
	}

}
