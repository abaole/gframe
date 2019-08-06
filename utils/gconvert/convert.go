package gconvert

import (
	"encoding/binary"
	"encoding/json"
	"reflect"
)

// int64 转 byte
func Int64ToBytes(i int64) []byte {
	var buf = make([]byte, 8)
	binary.BigEndian.PutUint64(buf, uint64(i))
	return buf
}

// byte 转 int64
func BytesToInt64(buf []byte) int64 {
	return int64(binary.BigEndian.Uint64(buf))
}

func Struct2MapString(obj interface{}) map[string]string {
	v := reflect.ValueOf(obj)
	k := v.Kind()
	if k == reflect.Ptr || k == reflect.Interface {
		v = v.Elem()
	}
	t := v.Type()
	var data = make(map[string]string)
	if v.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			data[SnakeName(t.Field(i).Name)] = v.Field(i).String()
		}
	}

	return data
}

func SnakeName(base string) string {
	var r = make([]rune, 0, len(base))
	var b = []rune(base)
	for i := 0; i < len(b); i++ {
		//if i > 0 && b[i] >= 'A' && b[i] <= 'Z' {
		//	r = append(r, '_', b[i]+32)
		//	continue
		//}
		if i == 0 && b[i] >= 'A' && b[i] <= 'Z' {
			r = append(r, b[i]+32)
			continue
		}
		r = append(r, b[i])
	}
	return string(r)
}

func Struct2String(src interface{}) string {
	b, err := json.Marshal(src)
	if err != nil {
		return ""
	}
	return string(b)
}

func FenToYuan(str string) string {
	amount := ToFloat64(str)
	amount = amount / 100
	return ToString(amount)
}
