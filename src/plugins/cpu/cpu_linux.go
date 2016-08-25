package cpu

import (
	"time"
	"github.com/shirou/gopsutil/cpu"
)

func CPUTimeStat_f() ([]CPUTimeStat_j,error){
	cpu,_:=cpu.CPUTimes(true)
	timestamp:=time.Now().Unix()

	var ret []CPUTimeStat_j
	var info CPUTimeStat_j
	for _,v :=range cpu {
		var total float64=v.User+v.System+v.Idle+v.Nice+v.Iowait+v.Irq+v.Softirq
		info = CPUTimeStat_j{
			CPU:v.CPU,
			User:v.User/total*100,
			System:v.System/total*100,
			Idle:v.Idle/total*100,
			Nice:v.Nice/total*100,
			Iowait:v.Iowait/total*100,
			Irq:v.Irq/total*100,
			Softirq:v.Softirq/total*100,
			Steal:v.Steal,
			Guest:v.Guest,
			GuestNice:v.GuestNice,
			Stolen:v.Stolen,
			Timestamp:timestamp,
		}
		ret=append(ret,info)
	}
	return ret,nil
}
