package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func updateBatteryStatus() {
	if runtime.GOOS == "windows" {
		powerInfo = PowerInfo{100, false}
		return
	}

	dat, err := os.ReadFile("/sys/class/power_supply/usb/online")
	if err != nil {
		fmt.Println(err.Error())
	}
	isCharging := strings.TrimSpace(string(dat)) == "1"
	voltage, err := os.ReadFile("/sys/class/power_supply/battery/voltage_now")
	if err != nil {
		fmt.Println(err.Error())
	}
	pct, err := strconv.Atoi(strings.TrimSpace(string(voltage)))

	if isCharging {
		pct = ((pct - MIN_VOLTAGE) - USB_VOLTAGE) * 100 / (MAX_VOLTAGE - MIN_VOLTAGE)
	} else {
		pct = (pct - MIN_VOLTAGE) * 100 / (MAX_VOLTAGE - MIN_VOLTAGE)
	}

	if pct > 100 {
		pct = 100
	}
	if pct < 0 {
		pct = 0
	}

	//voltage jumps a little bit but percentage cannot go up if we are not charging
	if !isCharging && pct < powerInfo.pct {
		pct = powerInfo.pct
	}

	powerInfo = PowerInfo{pct, isCharging}
}
