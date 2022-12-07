package resources

import (
	"embed"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	//go:embed media/*
	MEDIA_LIST embed.FS
)

func GetResource(fileName string) *sdl.RWops {
	file, _ := MEDIA_LIST.Open("media/" + fileName)
	stat, _ := file.Stat()
	buf := make([]byte, stat.Size())
	file.Read(buf)
	rwOps, _ := sdl.RWFromMem(buf)
	return rwOps
}
