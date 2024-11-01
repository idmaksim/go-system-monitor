package collectors

import (
	"math"

	"system-monitor/core/types"

	"github.com/shirou/gopsutil/mem"
)

func GetMemoryInfo(ch chan<- types.MemoryInfo) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		ch <- types.MemoryInfo{Total: 0, Used: 0, Err: err}
		return
	}
	total := float64(vmStat.Total) / math.Pow(1024, 3)
	used := int(vmStat.UsedPercent)
	ch <- types.MemoryInfo{Total: total, Used: used, Err: nil}
}
