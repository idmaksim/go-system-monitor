package collectors

import (
	"runtime"

	"system-monitor/core/displays"

	"github.com/shirou/gopsutil/cpu"
)

func PrintCPUInfo() {
	cpuPercent, err := cpu.Percent(0, false)
	if err == nil {
		usage := cpuPercent[0]

		displays.Cyan.Printf("\n🔄 CPU: ")
		switch {
		case usage >= 80:
			displays.Red.Printf("%.1f%%", usage)
		case usage >= 60:
			displays.Yellow.Printf("%.1f%%", usage)
		default:
			displays.Green.Printf("%.1f%%", usage)
		}

		displays.Cyan.Printf(" (Cores: ")
		displays.White.Printf("%d)\n", runtime.NumCPU())
	}
}
