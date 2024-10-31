package collectors

import (
	"system-monitor/core/displays"
	"time"

	"github.com/shirou/gopsutil/net"
)

type NetworkCache struct {
	interfaces []net.InterfaceStat
	lastUpdate time.Time
}

var networkCache = NetworkCache{
	lastUpdate: time.Time{},
}

const networkCacheTTL = 30 * time.Second

func PrintNetworkInfo() {
	if time.Since(networkCache.lastUpdate) > networkCacheTTL {
		interfaces, err := net.Interfaces()
		if err == nil {
			networkCache.interfaces = interfaces
			networkCache.lastUpdate = time.Now()
		}
	}

	displays.Cyan.Printf("\n🌐 Network Interfaces:\n")

	for _, iface := range networkCache.interfaces {
		if len(iface.Addrs) > 0 {
			displays.Cyan.Printf("   %s: ", iface.Name)
			displays.White.Printf("%v\n", iface.Addrs[1].Addr)
		}
	}
}
