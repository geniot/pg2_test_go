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

	PowerInformation PowerInfo = PowerInfo{100, false}
	DiskInfo1        DiskInfo  = DiskInfo{false, "", ""}
	DiskInfo2        DiskInfo  = DiskInfo{false, "", ""}
	CurrentVolume    int       = 0

	Renderer         *sdl.Renderer
	Font             *ttf.Font
	PressedKeysCodes mapset.Set[sdl.Keycode] = mapset.NewSet[sdl.Keycode]()
)
