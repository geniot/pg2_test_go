package gui

import (
	"geniot.com/geniot/pg2_test_go/internal/impl/mdl"
	"github.com/magiconair/properties"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"path/filepath"
	"strconv"
)

type ConfigImpl struct {
	props        *properties.Properties
	homeDir      string
	pathToConfig string
}

func NewConfig() *ConfigImpl {
	hD, _ := os.UserHomeDir()
	pToC := filepath.Join(hD, mdl.CONF_FILE_NAME)
	cfg := &ConfigImpl{nil,
		hD,
		pToC}
	cfg.load()
	return cfg
}

func (cfg ConfigImpl) Get(key string) uint32 {
	valStr, _ := cfg.props.Get(key)
	valInt, _ := strconv.ParseInt(valStr, 10, 0)
	return uint32(valInt)
}

func (cfg ConfigImpl) Set(key string, value string) {
	cfg.props.Set(key, value)
}

func (cfg *ConfigImpl) load() {
	loadedProps, _ := properties.LoadFile(cfg.pathToConfig, properties.UTF8)

	if loadedProps == nil {
		loadedProps = properties.NewProperties()
		loadedProps.Set(mdl.WINDOW_XPOS_KEY, strconv.FormatInt(int64(sdl.WINDOWPOS_UNDEFINED), 10))
		loadedProps.Set(mdl.WINDOW_YPOS_KEY, strconv.FormatInt(int64(sdl.WINDOWPOS_UNDEFINED), 10))
		displayMode, _ := sdl.GetCurrentDisplayMode(0)
		loadedProps.Set(mdl.WINDOW_WIDTH_KEY, strconv.FormatInt(int64(displayMode.W/2), 10))
		loadedProps.Set(mdl.WINDOW_HEIGHT_KEY, strconv.FormatInt(int64(displayMode.H/2), 10))
		loadedProps.Set(mdl.WINDOW_STATE_KEY, strconv.FormatInt(int64(sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE), 10))
	}

	//patching window state
	windowStateStr, _ := loadedProps.Get(mdl.WINDOW_STATE_KEY)
	windowState, _ := strconv.ParseInt(windowStateStr, 10, 0)
	windowState |= sdl.WINDOW_SHOWN
	windowState |= sdl.WINDOW_RESIZABLE
	loadedProps.Set(mdl.WINDOW_STATE_KEY, strconv.FormatInt(windowState, 10))

	cfg.props = loadedProps
}

func (cfg ConfigImpl) Save() {
	f, err := os.OpenFile(cfg.pathToConfig, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		println(err.Error())
	}
	defer f.Close()
	err = f.Truncate(0)
	_, err = f.Seek(0, 0)
	cfg.props.Write(f, properties.UTF8)
}
