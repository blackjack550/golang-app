package memory

import(
	"encoding/json"
)

type Memory_j struct  {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
	Free        uint64  `json:"free"`
	Active      uint64  `json:"active"`
	Inactive    uint64  `json:"inactive"`
	Buffers     uint64  `json:"buffers"`
	Cached      uint64  `json:"cached"`
	Timestamp	int64 	`json:"timestamp"`
}
func (m Memory_j) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}
