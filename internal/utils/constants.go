package utils

import (
	"os"
	"path/filepath"
)

const (
	APP_NAME       = "PG2 Test"
	APP_VERSION    = "0.1"
	CONF_FILE_NAME = ".pg2_test.properties"

	WINDOW_XPOS_KEY   = "WINDOW_XPOS_KEY"
	WINDOW_YPOS_KEY   = "WINDOW_YPOS_KEY"
	WINDOW_WIDTH_KEY  = "WINDOW_WIDTH_KEY"
	WINDOW_HEIGHT_KEY = "WINDOW_HEIGHT_KEY"
	WINDOW_STATE_KEY  = "WINDOW_STATE_KEY"
)

var (
	HOME_DIR, _    = os.UserHomeDir()
	PATH_TO_CONFIG = filepath.Join(HOME_DIR, CONF_FILE_NAME)
	WINDOW_TITLE   = APP_NAME + " " + APP_VERSION
)
