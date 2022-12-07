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
}

type ILoop interface {
	Stop()
}
