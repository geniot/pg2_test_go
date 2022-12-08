package api

type IApplication interface {
	Start()
	PlaySound()
	Stop()
}

type IWindow interface {
	SaveWindowState()
}

type IConfig interface {
	Get(key string) uint32
	Set(key string, value string)
	Save()
}

type IDevice interface {
	GetWindowState() uint32
	GetWindowPosAndSize() (int32, int32, int32, int32)
	ProcessKeyActions()
	UpdateBatteryStatus()
	UpdateDiskStatus()
	UpdateVolume()
	Stop()
	GetJoystickAxis(axis int) int16
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
