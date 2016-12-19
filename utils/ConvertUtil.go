package utils

import (
  "math"
  "strconv"
  "reflect"
)

func StringToInt64(str string) int64 {
  i, err := strconv.ParseInt(str, 10, 64)
  if (err != nil) {
    i = 0
  }
  return i
}

func StringToInt(str string) int {
  i, err := strconv.Atoi(str)
  if (err != nil) {
    i = 0
  }
  return i
}

func Int64ToString(i int64) string {
  return strconv.FormatInt(i, 10)
}

func IntToString(i int) string {
  return strconv.Itoa(i)
}


func ConvertToInt(obj interface{}) int {
  return StringToInt(obj.(string))
}

func ConvertToInt64(obj interface{}) int64 {
  return StringToInt64(obj.(string))
}

func ConvertToFloat64(obj interface{}) float64 {
  return obj.(float64)
}

func Round(f float64, n int, roundHalfUp bool) float64 {
  pow10_n := math.Pow10(n)
  if roundHalfUp {
    f += 0.5
  }
  return math.Trunc((f/pow10_n)*pow10_n) / pow10_n
}

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}
