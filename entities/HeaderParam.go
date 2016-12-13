package entities

type HeaderParam struct {
  ClientType    int
  UserId        int
  Token         string
  Longitude     float64
  Latitude      float64
  SoftVersion   string
  SystemVersion string
  DeviceId      string
  Time          string
  SignVersion   string
  Sign          string
  PhoneVersion  string
}

func NewHeaderParam() *HeaderParam {
  return &HeaderParam{}
}
