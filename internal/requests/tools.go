package requests

import "time"

const (
	timeLayout      = "2006-01-02T15:04:00"
	timeLayoutShort = "2006-01"
)

func TimeNow() string {
	now := time.Now().Format(timeLayoutShort)
	return now
}
