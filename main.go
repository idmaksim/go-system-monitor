package main

import (
	"strings"
	"system-monitor/core/collectors"
	"system-monitor/core/displays"
	"system-monitor/core/utils"
	"time"
)

func main() {
	for {
		utils.ClearScreen()

		displays.Magenta.Println(strings.Repeat("━", 58))
		displays.Magenta.Printf("%s\n", utils.CenterText("SYSTEM MONITOR", 65))
		displays.Magenta.Println(strings.Repeat("━", 58))

		now := time.Now()
		displays.Cyan.Printf("\n📅 Date: ")
		displays.White.Printf("%s\n", now.Format("02.01.2006"))
		displays.Cyan.Printf("🕒 Time: ")
		displays.White.Printf("%s\n", now.Format("15:04:05"))

		collectors.PrintSystemInfo()
		collectors.PrintCPUInfo()
		collectors.PrintMemoryInfo()
		collectors.PrintDiskInfo()
		collectors.PrintLoadAverage()
		collectors.PrintNetworkInfo()

		displays.Magenta.Printf("\n%s\n", strings.Repeat("━", 58))
		displays.White.Printf("Press Ctrl+C to exit\n")

		time.Sleep(500 * time.Millisecond)
	}
}
