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
	buf := make([]byte, stat.Size())
	file.Read(buf)
	rwOps, _ := sdl.RWFromMem(buf)
	return rwOps
}
