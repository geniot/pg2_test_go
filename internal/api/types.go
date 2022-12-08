package api

type PowerInfo struct {
	Pct        int
	IsCharging bool
}

type DiskInfo struct {
	IsDiskAvailable bool
	MaxDiskSpace    string
	FreeDiskSpace   string
}
