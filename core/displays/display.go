package displays

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
)

func DisplayCPUInfo(usage float64, cores int) {
	fmt.Printf("\n🔄 CPU: ")
	switch {
	case usage >= 80:
		Red.Printf("%.1f%%", usage)
	case usage >= 60:
		Yellow.Printf("%.1f%%", usage)
	default:
		Green.Printf("%.1f%%", usage)
	}
	Cyan.Printf(" (Cores: ")
	White.Printf("%d)\n", cores)
}

func DisplayMemoryInfo(total float64, used int) {
	Cyan.Printf("\n💾 Memory:\n")
	Cyan.Printf("   Total: ")
	White.Printf("%.2f GB\n", total)
	Cyan.Printf("   Used: ")

	switch {
	case used >= 80:
		Red.Printf("%d%%\n", used)
	case used >= 60:
		Yellow.Printf("%d%%\n", used)
	default:
		Green.Printf("%d%%\n", used)
	}
}

func DisplayLoadAverage(load1, load5, load15 float64) {
	Cyan.Printf("\n📊 Load Average: ")
	White.Printf("%.2f %.2f %.2f\n", load1, load5, load15)
}

func DisplayDiskInfo(partitions []disk.UsageStat) {
	Cyan.Printf("\n💿 Disks:\n")
	for _, usage := range partitions {
		Cyan.Printf("   %s: ", usage.Path)
		switch {
		case usage.UsedPercent >= 80:
			Red.Printf("%.1f%%\n", usage.UsedPercent)
		case usage.UsedPercent >= 60:
			Yellow.Printf("%.1f%%\n", usage.UsedPercent)
		default:
			Green.Printf("%.1f%%\n", usage.UsedPercent)
		}
	}
}

func DisplayNetworkInfo(interfaces []net.InterfaceStat) {
	Cyan.Printf("\n🌐 Network Interfaces:\n")
	for _, iface := range interfaces {
		if len(iface.Addrs) > 0 {
			Cyan.Printf("   %s: ", iface.Name)
			White.Printf("%v\n", iface.Addrs[0].Addr)
		}
	}
}

func DisplayTemperatureInfo(sensors []host.TemperatureStat) {
	Cyan.Printf("\n🌡️ Temperature Sensors:\n")
	for _, sensor := range sensors {
		Cyan.Printf("   %s: ", sensor.SensorKey)
		White.Printf("%.1f°C\n", sensor.Temperature)
	}
}

func DisplaySystemInfo(platform, version, hostname string) {
	Cyan.Printf("\n💻 System: ")
	White.Printf("%s %s\n", platform, version)
	Cyan.Printf("🏠 Hostname: ")
	White.Printf("%s\n", hostname)
}
