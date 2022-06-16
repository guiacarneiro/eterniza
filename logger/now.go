package logger

import "time"

func NowSiga() time.Time {
	return time.Now().Round(time.Second)
}
