package timerhelp

import "time"

type TimeSlice []time.Time

func (s TimeSlice) Less(i, j int) bool { return s[i].Before(s[j]) }
func (s TimeSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s TimeSlice) Len() int           { return len(s) }
