package rnd

import (
	"geniot.com/geniot/pg2_test_go/internal/api"
	"geniot.com/geniot/pg2_test_go/internal/ctx"
)

type DiskInfos struct {
	imageElements []ImageElement
}

func NewDiskInfos() api.IRenderable {
	return DiskInfos{initImageElements(api.DiskImages)}
}

func (diskInfos DiskInfos) Render() {
	renderImageElement(diskInfos.imageElements, 2)
	renderImageElement(diskInfos.imageElements, 3)

	if ctx.DiskInfo1.IsDiskAvailable {
		renderImageElement(diskInfos.imageElements, 0)
		text1 := ctx.DiskInfo1.FreeDiskSpace
		text2 := " / "
		text3 := ctx.DiskInfo1.MaxDiskSpace
		width1 := getTextWidth(text1)
		width2 := getTextWidth(text2)
		width3 := getTextWidth(text3)
		drawText(text1, 120-width1-width2-width3, 20, api.COLOR_RED)
		drawText(text2, 120-width2-width3, 20, api.COLOR_GREEN)
		drawText(text3, 120-width3, 20, api.COLOR_GRAY)
	}
	if ctx.DiskInfo2.IsDiskAvailable {
		renderImageElement(diskInfos.imageElements, 1)
		width1 := drawText(ctx.DiskInfo2.FreeDiskSpace, 197, 20, api.COLOR_RED)
		width2 := drawText(" / ", 197+width1, 20, api.COLOR_GREEN)
		drawText(ctx.DiskInfo2.MaxDiskSpace, 197+width1+width2, 20, api.COLOR_GRAY)
	}
}
