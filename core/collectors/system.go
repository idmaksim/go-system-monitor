package collectors

import (
	"system-monitor/core/displays"

	"github.com/shirou/gopsutil/host"
)

func PrintSystemInfo() {
	hostInfo, err := host.Info()
	if err != nil {
		return
	}
	displays.Cyan.Printf("\n💻 System: ")
	displays.White.Printf("%s %s\n", hostInfo.Platform, hostInfo.PlatformVersion)
	displays.Cyan.Printf("🏠 Hostname: ")
	displays.White.Printf("%s\n", hostInfo.Hostname)
}
