package collectors

import (
	"system-monitor/core/types"

	"github.com/shirou/gopsutil/load"
)

func GetLoadAverage(ch chan<- types.LoadAverage) {
	loadStat, err := load.Avg()
	if err != nil {
		ch <- types.LoadAverage{Load1: 0, Load5: 0, Load15: 0, Err: err}
		return
	}
	ch <- types.LoadAverage{
		Load1:  loadStat.Load1,
		Load5:  loadStat.Load5,
		Load15: loadStat.Load15,
		Err:    nil,
	}
}
