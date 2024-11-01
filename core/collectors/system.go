package collectors

import (
	"system-monitor/core/types"

	"github.com/shirou/gopsutil/host"
)

func GetSystemInfo(ch chan<- types.SystemInfo) {
	hostInfo, err := host.Info()
	if err != nil {
		ch <- types.SystemInfo{Platform: "", Version: "", Hostname: "", Err: err}
		return
	}
	ch <- types.SystemInfo{
		Platform: hostInfo.Platform,
		Version:  hostInfo.PlatformVersion,
		Hostname: hostInfo.Hostname,
		Err:      nil,
	}
}
