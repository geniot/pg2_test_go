package main

import (
	"fmt"
	"github.com/pydio/minio-srv/pkg/disk"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func updateDiskStatus() {
	if runtime.GOOS == "windows" {
		updateDiskInfo("C:\\", &diskInfos[0])
	} else {
		updateDiskInfo("/usr/local/home", &diskInfos[0])
		updateDiskInfo("/media/sdcard/", &diskInfos[1])
	}
}

func updateDiskInfo(path string, diskInfo *DiskInfo) {
	di, err := disk.GetInfo(path)
	if err != nil {
		println(err.Error())
		diskInfo.isDiskAvailable = false
		return
	}
	//we only do this once: when initializing or the disk has been inserted
	if !diskInfo.isDiskAvailable {
		diskInfo.isDiskAvailable = true
		diskInfo.maxDiskSpace = Bytes(di.Total)
		diskInfo.freeDiskSpace = Bytes(di.Free)
	}

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
	//we cannot go down when charging
	if isCharging && pct < powerInfo.pct {
		pct = powerInfo.pct
	}

	powerInfo.pct = pct
	powerInfo.isCharging = isCharging
}
