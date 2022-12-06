package gui

import (
	"geniot.com/geniot/pg2_test_go/internal/domain"
	"geniot.com/geniot/pg2_test_go/internal/model"
	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	model  *domain.Model
	loop   *Loop
	window *Window
	config *Config
	scene  *model.Scene
}

func NewApplication() *Application {
	return &Application{nil, nil, nil, nil, nil}
}

func (app Application) Start() {
	sdl.Init(sdl.INIT_EVERYTHING)

	cnf := NewConfig(&app)
	app.config = cnf
	mdl := domain.NewModel()
	app.model = mdl
	wnd := NewWindow(&app)
	app.window = wnd
	lp := NewLoop(&app)
	app.loop = lp

	scn := model.NewScene(wnd.sdlRenderer)
	app.scene = scn

	app.loop.Start()
}
