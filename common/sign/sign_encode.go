package sign

import (
	"errors"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type Dec []string

func (p Dec) Len() int           { return len(p) }
func (p Dec) Less(i, j int) bool { return p[i] < p[j] }
func (p Dec) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func Encode(o interface{}) string {
	dec := encodeField("", o)
	sort.Stable(dec)
	return strings.Join(dec, "&")
}
func EncodeAndAppend(o interface{}, key, value string) string {
	dec := encodeField("", o)
	dec = append(dec, key+"="+value)
	sort.Stable(dec)
	return strings.Join(dec, "&")
}
func EncodeAndAppendMap(o interface{}, dataMap map[string]string) string {
	dec := encodeField("", o)
	for k, v := range dataMap {
		dec = append(dec, k+"="+v)
	}
	sort.Stable(dec)
	return strings.Join(dec, "&")
}
func encodeField(prefix string, o interface{}) Dec {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	valStr, err := getValueString(v)
	var dec Dec
	if err == nil {
		values := url.Values{}
		values.Set("url", valStr)
		return append(dec, prefix+"="+strings.Split(values.Encode(), "=")[1])
	}
	if prefix != "" {
		prefix += "."
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		prefix2 := t.Field(i).Tag.Get("sign")
		value := v.Field(i)
		if prefix2 == "" {
			continue
		}
		arrs := strings.Split(prefix2, ",")
		if len(arrs) >= 1 {
			prefix2 = arrs[0]
		}
		if isArray(value) {
			for j := 0; j < value.Len(); j++ {
				dec = append(dec, encodeField(prefix+prefix2+"[]", value.Index(j).Interface())...)
			}
		} else if value.Kind() == reflect.Map {
			dataMap := value.Interface().(map[string]interface{})
			for k, v := range dataMap {
				dec = append(dec, encodeField(prefix+prefix2+"."+k, v)...)
			}

		} else {
			dec = append(dec, encodeField(prefix+prefix2, value.Interface())...)
		}

	}
	return dec
}

func getValueString(value reflect.Value) (string, error) {
	if value.Type().String() == "string" {
		return value.Interface().(string), nil
	}
	if value.Type().String() == "bool" {
		val := value.Interface().(bool)
		if val {
			return "true", nil
		} else {
			return "false", nil
		}
	}
	if value.Type().String() == "int8" {
		val := value.Interface().(int8)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "int16" {
		val := value.Interface().(int16)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "int32" {
		val := value.Interface().(int32)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "int64" {
		val := value.Interface().(int64)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "uint8" {
		val := value.Interface().(uint8)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "uint16" {
		val := value.Interface().(uint16)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "uint32" {
		val := value.Interface().(uint32)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "int" {
		val := value.Interface().(int)
		return strconv.FormatInt(int64(val), 10), nil
	}
	if value.Type().String() == "uint" {
		val := value.Interface().(uint)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "uint64" {
		val := value.Interface().(uint64)
		return strconv.FormatUint(uint64(val), 10), nil
	}
	if value.Type().String() == "float32" {
		val := value.Interface().(float32)
		return strconv.FormatFloat(float64(val), 'E', -1, 32), nil
	}
	if value.Type().String() == "float64" {
		val := value.Interface().(float64)
		return strconv.FormatFloat(float64(val), 'E', -1, 64), nil
	}
	if value.Type().String() == "uint32" {
		val := value.Interface().(uint32)
		return strconv.FormatUint(uint64(val), 10), nil
	}

	return "", errors.New("invalid field")

}
func isArray(value reflect.Value) bool {
	if strings.HasPrefix(value.Type().String(), "[]") {
		return true
	}
	return false
}
