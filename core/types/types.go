package types

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/net"
)

type CPUInfo struct {
	Usage float64
	Cores int
	Err   error
}

type MemoryInfo struct {
	Total float64
	Used  int
	Err   error
}

type LoadAverage struct {
	Load1  float64
	Load5  float64
	Load15 float64
	Err    error
}

type DiskInfo struct {
	Partitions []disk.UsageStat
	Err        error
}

type NetworkInfo struct {
	Interfaces []net.InterfaceStat
	Err        error
}

type TemperatureInfo struct {
	Sensors []host.TemperatureStat
	Err     error
}

type SystemInfo struct {
	Platform string
	Version  string
	Hostname string
	Err      error
}
