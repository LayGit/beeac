package entities

type Result struct {
  Code    int         `json:"code"`
  Data    interface{} `json:"data"`
  Message string      `json:"message"`
}

func NewResult(code int, data interface{}, message string) *Result {
  return &Result{ Code: code, Data: data, Message: message }
}

func NewResultWithCodeMessage(code int, message string) *Result {
  return NewResult(code, nil, message)
}

func NewResultWithData(data interface{}) *Result {
  return NewResult(0, data, "")
}

func NewResultOfParamError() *Result {
  return NewResultWithCodeMessage(-400, "请求参数不合法")
}

func NewResultOfNotFound() *Result {
  return NewResultWithCodeMessage(-404, "无效的接口请求")
}

func NewResultOfServerError() *Result {
  return NewResultWithCodeMessage(-500, "操作失败，请稍候再试")
}

func NewResultOfAuthFailed() *Result {
  return NewResultWithCodeMessage(-2, "无效的用户令牌")
}

func NewResultOfWechatAuthFailed() *Result {
  return NewResultWithCodeMessage(-3, "微信授权失败")
}

func NewResultOfError(message string) *Result {
  return NewResultWithCodeMessage(-4, message)
}
