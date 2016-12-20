package utils

import qrcode "github.com/skip2/go-qrcode"

func QRCodeEncode(url string, size int) string {
  byt, err := qrcode.Encode(url, qrcode.Medium, size)
  if err != nil {
    return ""
  } else {
    return "data:image/png;base64," + Base64Encode(byt)
  }
}
