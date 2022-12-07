package ctx

import (
	"geniot.com/geniot/pg2_test_go/internal/api"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

var (
	Application api.IApplication
	Window      api.IWindow
	Device      api.IDevice
	Config      api.IConfig

	Loop        api.IStartable
	EventLoop   api.IRunnable
	PhysicsLoop api.IRunnable
	RenderLoop  api.IRunnable

	CurrentScene api.IRenderable

	Renderer         *sdl.Renderer
	Font             *ttf.Font
	PressedKeysCodes mapset.Set[sdl.Keycode] = mapset.NewSet[sdl.Keycode]()
)
