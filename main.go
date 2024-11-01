package main

import (
	"strings"
	"sync"
	"system-monitor/core/collectors"
	"system-monitor/core/displays"
	"system-monitor/core/types"
	"system-monitor/core/utils"
	"time"
)

func main() {
	for {
		utils.ClearScreen()

		displays.Magenta.Println(strings.Repeat("━", 58))
		displays.Magenta.Printf("%s\n", utils.CenterText("SYSTEM MONITOR", 58))
		displays.Magenta.Println(strings.Repeat("━", 58))

		now := time.Now()
		displays.Cyan.Printf("\n📅 Date: ")
		displays.White.Printf("%s\n", now.Format("02.01.2006"))
		displays.Cyan.Printf("🕒 Time: ")
		displays.White.Printf("%s\n", now.Format("15:04:05"))

		cpuChan := make(chan types.CPUInfo, 1)
		memChan := make(chan types.MemoryInfo, 1)
		loadChan := make(chan types.LoadAverage, 1)
		diskChan := make(chan types.DiskInfo, 1)
		networkChan := make(chan types.NetworkInfo, 1)
		tempChan := make(chan types.TemperatureInfo, 1)
		systemChan := make(chan types.SystemInfo, 1)

		var wg sync.WaitGroup
		wg.Add(7)

		go func() {
			collectors.GetCPUInfo(cpuChan)
			wg.Done()
		}()
		go func() {
			collectors.GetMemoryInfo(memChan)
			wg.Done()
		}()
		go func() {
			collectors.GetLoadAverage(loadChan)
			wg.Done()
		}()
		go func() {
			collectors.GetDiskInfo(diskChan)
			wg.Done()
		}()
		go func() {
			collectors.GetNetworkInfo(networkChan)
			wg.Done()
		}()
		go func() {
			collectors.GetTemperatureInfo(tempChan)
			wg.Done()
		}()
		go func() {
			collectors.GetSystemInfo(systemChan)
			wg.Done()
		}()

		wg.Wait()

		cpuData := <-cpuChan
		memData := <-memChan
		loadData := <-loadChan
		diskData := <-diskChan
		networkData := <-networkChan
		tempData := <-tempChan
		systemData := <-systemChan

		if cpuData.Err == nil {
			displays.DisplayCPUInfo(cpuData.Usage, cpuData.Cores)
		}

		if memData.Err == nil {
			displays.DisplayMemoryInfo(memData.Total, memData.Used)
		}

		if loadData.Err == nil {
			displays.DisplayLoadAverage(loadData.Load1, loadData.Load5, loadData.Load15)
		}

		if diskData.Err == nil {
			displays.DisplayDiskInfo(diskData.Partitions)
		}

		if networkData.Err == nil {
			displays.DisplayNetworkInfo(networkData.Interfaces)
		}

		if tempData.Err == nil {
			displays.DisplayTemperatureInfo(tempData.Sensors)
		}

		if systemData.Err == nil {
			displays.DisplaySystemInfo(systemData.Platform, systemData.Version, systemData.Hostname)
		}

		displays.Magenta.Printf("\n%s\n", strings.Repeat("━", 58))
		displays.White.Printf("Press Ctrl+C to exit\n")

		time.Sleep(500 * time.Millisecond)
	}
}
