package utils

import (
  "time"
)

func GetNowTimestamp() int64 {
  return time.Now().UnixNano() / int64(time.Millisecond)
}
