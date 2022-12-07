package mdl

import (
	"fmt"
	"math"
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

func initImageElements(imageDescriptors []ImageDescriptor) []ImageElement {
	imgElements := make([]ImageElement, len(imageDescriptors))
	for i := range imageDescriptors {
		iEl := NewImageElement(
			imageDescriptors[i].ImageName,
			imageDescriptors[i].OffsetX,
			imageDescriptors[i].OffsetY,
			imageDescriptors[i].DisplayOnPress)
		imgElements[i] = iEl
	}
	return imgElements
}
