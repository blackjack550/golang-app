package disk

import(
	"github.com/shirou/gopsutil/disk"
	"time"
)

func DiskUsage_f() ([]DiskUsageStat_j,error) {
	partition,_ := disk.DiskPartitions(true)
	timestamp:=time.Now().Unix()
	var ret []DiskUsageStat_j
	var info DiskUsageStat_j
	for _,v :=range partition {
		usage,_ := disk.DiskUsage(v.Device)
		info = DiskUsageStat_j{
			Path:usage.Path,
			Fstype:usage.Fstype,
			Total:usage.Total,
			Free:usage.Free,
			Used:usage.Used,
			UsedPercent:usage.UsedPercent,
			InodesTotal:usage.InodesTotal,
			InodesUsed:usage.InodesUsed,
			InodesFree:usage.InodesFree,
			InodesUsedPercent:usage.InodesUsedPercent,
			TimeStamp:timestamp,
		}
		ret = append(ret,info)
	}
	return ret,nil
}
