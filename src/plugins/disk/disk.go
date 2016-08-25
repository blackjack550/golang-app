package disk

import (
	"encoding/json"
)

type DiskUsageStat_j struct {
	Path              string  `json:"path"`
	Fstype            string  `json:"fstype"`
	Total             uint64  `json:"total"`
	Free              uint64  `json:"free"`
	Used              uint64  `json:"used"`
	UsedPercent       float64 `json:"used_percent"`
	InodesTotal       uint64  `json:"inodes_total"`
	InodesUsed        uint64  `json:"inodes_used"`
	InodesFree        uint64  `json:"inodes_free"`
	InodesUsedPercent float64 `json:"inodes_used_percent"`
	TimeStamp		  int64		`json:"timestamp"`
}

func (d DiskUsageStat_j) String() string {
	s, _ := json.Marshal(d)
	return string(s)
}