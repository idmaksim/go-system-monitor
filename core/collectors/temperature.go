package collectors

import (
	"system-monitor/core/types"

	"github.com/shirou/gopsutil/host"
)

func GetTemperatureInfo(ch chan<- types.TemperatureInfo) {
	sensors, err := host.SensorsTemperatures()
	if err != nil {
		ch <- types.TemperatureInfo{Sensors: nil, Err: err}
		return
	}
	ch <- types.TemperatureInfo{Sensors: sensors, Err: nil}
}
