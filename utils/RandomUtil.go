package utils

import (
  "time"
  "math/rand"
)

func RandomNum(length int) string {
  if length < 0 {
    length = 1
  }

  sbNum := NewStringBuilder()
  r := rand.New(rand.NewSource(time.Now().UnixNano()))
  for i := 0; i < length; i++ {
    sbNum.Append(r.Intn(9))
  }
  return sbNum.ToString()
}
