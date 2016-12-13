package sms

import (
  "net/http"
  "net/url"
  "strings"
  "time"
  "io/ioutil"
  "encoding/xml"

  "sesan.cc/app/utils"

  "github.com/astaxie/beego"
)

func Send(phone, code string) (*SmsResult, error) {
  content := strings.Replace(beego.AppConfig.String("SMS_TEMPLATE"), "%code%", code, -1)
  unixTime := time.Now().Unix()

  sbPwd := utils.NewStringBuilder()
  sbPwd.Append(beego.AppConfig.String("SMS_ACCOUNT"))
  sbPwd.Append(beego.AppConfig.String("SMS_API_KEY"))
  sbPwd.Append(phone)
  sbPwd.Append(content)
  sbPwd.Append(unixTime)

  resp, err := http.PostForm(beego.AppConfig.String("SMS_API_URL"), url.Values{
    "account": {beego.AppConfig.String("SMS_ACCOUNT")},
    "password": {utils.MD5Encrypt(sbPwd.ToString())},
    "mobile": {phone},
    "content": {content},
    "time": {utils.Int64ToString(unixTime)},
  })

  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  // 解析结果
  var smsResult SmsResult
  if xml.Unmarshal(body, &smsResult) != nil {
    return nil, err
  }

  return &smsResult, nil
}
