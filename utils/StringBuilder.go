package utils

import (
  "bytes"
)

type StringBuilder struct {
    buf bytes.Buffer
}

func NewStringBuilder() *StringBuilder {
  return &StringBuilder{ buf: bytes.Buffer{} }
}

func (this *StringBuilder) Append(obj interface{}) {
  this.buf.WriteString(obj.(string))
}

func (this *StringBuilder) ToString() string {
  return this.buf.String()
}
