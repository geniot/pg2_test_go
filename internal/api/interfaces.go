package api

type IApplication interface {
	Start()
	PlaySound()
}

type IWindow interface {
	OnBeforeClose()
}

type IConfig interface {
	Get(key string) uint32
	Set(key string, value string)
	Save()
}

type IDevice interface {
	IsRumbleSupported() bool
	Rumble()
	PlaySound()
	GetWindowState() uint32
	GetWindowPosAndSize() (int32, int32, int32, int32)
}

type IRunnable interface {
	Run()
}

type IStartable interface {
	Start()
	Stop()
}

type IRenderable interface {
	Render()
}
