package memory
import(
	"github.com/shirou/gopsutil/mem"
	"time"
)

func Memory_f() (*Memory_j,error){
	mem,_:=mem.VirtualMemory()
	timestamp:=time.Now().Unix()

	//	fmt.Println(mem)
	ret:=&Memory_j{
		Total       :mem.Total,
		Available   :mem.Available,
		Used        :mem.Used,
		UsedPercent :mem.UsedPercent,
		Free        :mem.Free,
		Active      :mem.Active,
		Inactive    :mem.Inactive,
		Buffers     :mem.Buffers,
		Cached      :mem.Cached,
		Timestamp	:timestamp,
	}
	return ret,nil

}