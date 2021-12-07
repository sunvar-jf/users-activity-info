package model

import (
	"fileparsemod/src/helpers/timerhelp"
)

type UserInfo struct {
	LoginDate     timerhelp.TimeSlice //[]time.Time
	LoginDateTime []string
}
