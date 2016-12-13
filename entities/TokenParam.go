package entities

type TokenParam struct {
  UserId int
  Token string
}

func NewTokenParam() *TokenParam {
  return &TokenParam{ UserId: 0, Token: "" }
}
