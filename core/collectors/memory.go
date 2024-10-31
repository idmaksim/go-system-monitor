package collectors

import (
	"math"
	"system-monitor/core/displays"

	"github.com/shirou/gopsutil/mem"
)

func PrintMemoryInfo() {
	vmStat, err := mem.VirtualMemory()
	if err == nil {
		displays.Cyan.Printf("\n💾 Memory:\n")
		displays.Cyan.Printf("   Total: ")
		displays.White.Printf("%.2f GB\n", float64(vmStat.Total)/math.Pow(1024, 3))
		displays.Cyan.Printf("   Used: ")

		switch {
		case vmStat.UsedPercent >= 80:
			displays.Red.Printf("%d%%\n", int(vmStat.UsedPercent))
		case vmStat.UsedPercent >= 60:
			displays.Yellow.Printf("%d%%\n", int(vmStat.UsedPercent))
		default:
			displays.Green.Printf("%d%%\n", int(vmStat.UsedPercent))
		}
	}
}
