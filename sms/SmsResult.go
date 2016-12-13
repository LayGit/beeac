package sms

type SmsResult struct {
  SmsId   string `sms:"smsid"`
  SmsMsg  string `sms:"msg"`
  SmsCode int    `sms:"code"`
}
