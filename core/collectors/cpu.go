package collectors

import (
	"runtime"
	"system-monitor/core/types"

	"github.com/shirou/gopsutil/cpu"
)

func GetCPUInfo(ch chan<- types.CPUInfo) {
	cpuPercent, err := cpu.Percent(0, false)
	if err != nil {
		ch <- types.CPUInfo{Usage: 0, Cores: 0, Err: err}
		return
	}
	usage := cpuPercent[0]
	cores := runtime.NumCPU()
	ch <- types.CPUInfo{Usage: usage, Cores: cores, Err: nil}
}
