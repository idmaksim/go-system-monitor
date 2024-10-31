package collectors

import (
	"system-monitor/core/displays"

	"github.com/shirou/gopsutil/disk"
)

func PrintDiskInfo() {
	displays.Cyan.Printf("\n💿 Disks:\n")
	partitions, err := disk.Partitions(false)
	if err == nil {
		for _, partition := range partitions {
			usage, err := disk.Usage(partition.Mountpoint)
			if err != nil || usage.Total == 0 {
				continue
			}
			displays.Cyan.Printf("   %s: ", partition.Mountpoint)
			switch {
			case usage.UsedPercent >= 80:
				displays.Red.Printf("%.1f%%\n", usage.UsedPercent)
			case usage.UsedPercent >= 60:
				displays.Yellow.Printf("%.1f%%\n", usage.UsedPercent)
			default:
				displays.Green.Printf("%.1f%%\n", usage.UsedPercent)
			}
		}
	}
}
