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

	PowerInformation api.PowerInfo = api.PowerInfo{100, false}
	DiskInfo1        api.DiskInfo  = api.DiskInfo{false, "", ""}
	DiskInfo2        api.DiskInfo  = api.DiskInfo{false, "", ""}
	CurrentVolume    int           = -1

	Renderer         *sdl.Renderer
	Font             *ttf.Font
	PressedKeysCodes mapset.Set[sdl.Keycode] = mapset.NewSet[sdl.Keycode]()
	LastPressedKey   sdl.Keycode
)
