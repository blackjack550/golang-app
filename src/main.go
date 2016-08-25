package main

import (
	"plugins/memory"
//	"plugins/load"
//	"plugins/net"
//	"plugins/cpu"
//	"time"
//	"plugins/disk"
//	"log"
	"os"
	"io"
	"fmt"
	"time"
)

const (
	LOG_FILE_CPU="monitor_cpu.csv"
	LOG_FILE_MEM="monitor_mem.csv"
)

func checkFileIsNotExist(filename string) bool{
	var exist=false
	if _,err := os.Stat(filename); os.IsNotExist(err) {
		exist=true;
	}
	return exist
}
func csvTimeChange(unixtime int64) string{
	csvtime:=time.Unix(unixtime,0).Format("2006-01-02 15:04:05")
	return csvtime
}
func Memory_Info() {
	mem,_:=memory.Memory_f()
	header:="时间,内存总数,内存可用,内存已使用,内存使用率,内存活动,内存非活动,内存buffer,内存cache"
	if checkFileIsNotExist(LOG_FILE_MEM) {
		os.Create(LOG_FILE_MEM)
		f,_ :=os.OpenFile(LOG_FILE_MEM,os.O_CREATE,0660)
		_,err :=io.WriteString(f,header+"\n")
		if err != nil {
			fmt.Println("xierucuowu")
		}
	}
	timestamp:=csvTimeChange(mem.Timestamp)
	meminfo:=fmt.Sprintf("%s,%d,%d,%d,%f,%d,%d,%d,%d",timestamp,mem.Total,mem.Free,mem.Used,mem.UsedPercent,mem.Active,mem.Inactive,mem.Buffers,mem.Cached)
	f,_ :=os.OpenFile(LOG_FILE_MEM,os.O_APPEND|os.O_WRONLY,0666)
	_,err :=io.WriteString(f,meminfo+"\n")
	if err != nil {
		fmt.Println("xierucuowu")
	}
	defer f.Close()
}

func main(){
	Memory_Info()
//	Memory_Info()
//	fmt.Println(mem)
//	test1,_ := load.Load_f()
//	fmt.Println(test1)
//	adapter,_:=net.NetInterfaces()
//	fmt.Println(adapter)
//	traffic,_:=net.NetIOCounters(true)
//	fmt.Printf("网卡%s,接受流量:%d MB,发送流量:%d MB",traffic[3].Name,traffic[3].BytesRecv>>20,traffic[3].BytesSent>>20)
//	network,_:=net.NetIOCountersStat_f()
//	for _,v :=range network {
//		fmt.Println(v)
//	}
//	cpu,_ := cpu.CPUTimeStat_f()
//	for _,v :=range cpu {
//		fmt.Println(v)
//	}
//	diskuse,_:=disk.DiskUsage_f()
//	fmt.Println(diskuse)
//	count,_ :=disk.DiskIOCounters()
//	for k,v :=range count{
//		fmt.Println(k,v)
//	}



}