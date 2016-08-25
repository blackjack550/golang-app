package load

import (
	"encoding/json"
)

type LoadAvgStat_j struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
	Timestamp int64 `json:"timestamp"`
}

func (l LoadAvgStat_j) String() string {
	s, _ := json.Marshal(l)
	return string(s)
}