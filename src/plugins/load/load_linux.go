package load

import(
	"github.com/shirou/gopsutil/load"
	"time"
)

func Load_f() (*LoadAvgStat_j,error){
	load,_:=load.LoadAvg()
	timestamp:=time.Now().Unix()

	//	fmt.Println(mem)
	ret:=&LoadAvgStat_j{
		Load1:load.Load1,
		Load5:load.Load5,
		Load15:load.Load15,
		Timestamp	:timestamp,
	}
	return ret,nil

}