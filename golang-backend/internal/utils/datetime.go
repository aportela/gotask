package utils

import "time"

func CurrentMSTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
