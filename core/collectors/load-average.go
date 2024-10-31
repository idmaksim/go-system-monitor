package collectors

import (
	"system-monitor/core/displays"

	"github.com/shirou/gopsutil/load"
)

func PrintLoadAverage() {
	loadStat, err := load.Avg()
	if err == nil {
		displays.Cyan.Printf("\n📊 Load Average: ")
		displays.White.Printf("%.2f %.2f %.2f\n",
			loadStat.Load1,
			loadStat.Load5,
			loadStat.Load15)
	}
}
