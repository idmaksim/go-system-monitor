package collectors

import (
	"system-monitor/core/types"
	"time"

	"github.com/shirou/gopsutil/net"
)

type NetworkCache struct {
	Interfaces []net.InterfaceStat
	LastUpdate time.Time
}

var networkCache = NetworkCache{
	LastUpdate: time.Time{},
}

const networkCacheTTL = 30 * time.Second

func GetNetworkInfo(ch chan<- types.NetworkInfo) {
	if time.Since(networkCache.LastUpdate) > networkCacheTTL {
		interfaces, err := net.Interfaces()
		if err != nil {
			ch <- types.NetworkInfo{Interfaces: nil, Err: err}
			return
		}
		networkCache.Interfaces = interfaces
		networkCache.LastUpdate = time.Now()
	}
	ch <- types.NetworkInfo{Interfaces: networkCache.Interfaces, Err: nil}
}
