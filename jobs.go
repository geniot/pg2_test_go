package main

import (
	"fmt"
	humanize "github.com/dustin/go-humanize"
	"github.com/pydio/minio-srv/pkg/disk"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func updateDiskStatus() {
	printUsage("/usr/local/home")
	printUsage("/media/sdcard/")
}

func printUsage(path string) {
	di, _ := disk.GetInfo(path)
	percentage := (float64(di.Total-di.Free) / float64(di.Total)) * 100
	fmt.Printf("%s of %s disk space used (%0.2f%%)\n",
		humanize.Bytes(di.Total-di.Free),
		humanize.Bytes(di.Total),
		percentage,
	)
}

func updateBatteryStatus() {
	if runtime.GOOS == "windows" {
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

	//voltage jumps a little but percentage cannot go up if we are not charging
	if !isCharging && pct > powerInfo.pct {
		pct = powerInfo.pct
	}

	powerInfo.pct = pct
	powerInfo.isCharging = isCharging
}
