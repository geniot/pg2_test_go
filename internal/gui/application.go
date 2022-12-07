package gui

import (
	"geniot.com/geniot/pg2_test_go/internal/mdl"
	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	model  *mdl.Model
	loop   *Loop
	window *Window
	config *Config
	scene  *mdl.Scene
}

func NewApplication() *Application {
	return &Application{nil, nil, nil, nil, nil}
}

func (app Application) Start() {
	sdl.Init(sdl.INIT_EVERYTHING)

	app.config = NewConfig()
	app.model = mdl.NewModel()
	app.window = NewWindow(&app)
	app.loop = NewLoop(&app)
	app.scene = mdl.NewScene(app.window.sdlRenderer)

	app.loop.Start()
}
