package entities

type MutableResult struct {
  Code    int                     `json:"code"`
  Data    map[string]interface{}  `json:"data"`
  Message string                  `json:"message"`
}

func NewMutableResult() *MutableResult {
  return &MutableResult{ Code: 0, Data: make(map[string]interface{}), Message: "" }
}

func (this *MutableResult) Put(key string, val interface{}) {
  this.Data[key] = val
}

func (this *MutableResult) Remove(key string) {
  delete(this.Data, key)
}

func (this *MutableResult) Contains(key string) bool {
  if _, exist := this.Data[key]; exist {
    return true
  } else {
    return false
  }
}

func (this *MutableResult) ToResult() *Result {
  return NewResult(this.Code, this.Data, this.Message)
}
