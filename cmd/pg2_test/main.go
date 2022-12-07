package main

import (
	"geniot.com/geniot/pg2_test_go/internal/ctx"
	"geniot.com/geniot/pg2_test_go/internal/impl/gui"
)

func main() {
	ctx.Application = gui.NewApplication()
	ctx.Application.Start()
}
