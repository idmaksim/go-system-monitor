package collectors

import (
	"system-monitor/core/types"

	"github.com/shirou/gopsutil/disk"
)

func GetDiskInfo(ch chan<- types.DiskInfo) {
	partitions, err := disk.Partitions(false)
	if err != nil {
		ch <- types.DiskInfo{Partitions: nil, Err: err}
		return
	}

	var usageStats []disk.UsageStat
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil || usage.Total == 0 {
			continue
		}
		usageStats = append(usageStats, *usage)
	}

	ch <- types.DiskInfo{Partitions: usageStats, Err: nil}
}
