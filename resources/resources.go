package resources

import (
	"embed"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	//go:embed media/*
	mediaList embed.FS
)

func GetResource(fileName string) *sdl.RWops {
	file, _ := mediaList.Open("media/" + fileName)
	stat, _ := file.Stat()
	size := stat.Size()
	buf := make([]byte, size)
	r, err := file.Read(buf)
	if err != nil {
		println(err.Error())
	}
	if r < int(size) {
		println("Couldn't read the whole file.")
	}
	rwOps, _ := sdl.RWFromMem(buf)
	return rwOps
}
