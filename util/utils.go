package util

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : utils.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/8 23:37
* 修改历史 : 1. [2022/4/8 23:37] 创建文件 by NST
*/

//Substring
//获取source的子串,如果start小于0或者end大于source长度则返回""
//start:开始index，从0开始，包括0
//end:结束index，以end结束，但不包括end
func Substring(source string, start int, end int) string {
	var r = []rune(source)

	if end <= 0 {
		return string(r[start:])
	}

	return string(r[start:end])
}

func Str2Int64(source string) (any, error) {
	v, error := strconv.ParseInt(fmt.Sprintf("%v", source), 10, 64)
	if error == nil {
		return int64(v), nil
	}
	return nil, error
}

func Str2Int32(source string) (any, error) {
	v, error := strconv.ParseInt(fmt.Sprintf("%v", source), 10, 64)
	if error == nil {
		return int32(v), nil
	}
	return nil, error
}

func Str2Int16(source string) (any, error) {
	v, error := strconv.ParseInt(fmt.Sprintf("%v", source), 10, 64)
	if error == nil {
		return int16(v), nil
	}
	return nil, error
}

func Str2Int8(source string) (any, error) {
	v, error := strconv.ParseInt(fmt.Sprintf("%v", source), 10, 64)
	if error == nil {
		return int8(v), nil
	}
	return nil, error
}

func Str2Int(source string) (any, error) {
	v, error := strconv.Atoi(fmt.Sprintf("%v", source))
	if error == nil {
		return v, nil
	}
	return nil, error
}

func Str2UInt64(source string) (any, error) {
	v, error := strconv.ParseUint(fmt.Sprintf("%v", source), 10, 64)
	if error == nil {
		return uint64(v), nil
	}
	return nil, error
}

func Str2Uint32(source string) (any, error) {
	v, error := strconv.ParseUint(fmt.Sprintf("%v", source), 10, 64)
	if error == nil {
		return uint32(v), nil
	}
	return nil, error
}

func Str2Uint16(source string) (any, error) {
	v, error := strconv.ParseUint(fmt.Sprintf("%v", source), 10, 64)
	if error == nil {
		return uint16(v), nil
	}
	return nil, error
}

func Str2Uint8(source string) (any, error) {
	v, error := strconv.ParseUint(fmt.Sprintf("%v", source), 10, 64)
	if error == nil {
		return uint8(v), nil
	}
	return nil, error
}

func Str2Uint(source string) (any, error) {
	v, error := strconv.ParseUint(fmt.Sprintf("%v", source), 10, 64)
	if error == nil {
		return uint(v), nil
	}
	return nil, error
}

func Str2Bool(source string) (any, error) {
	v, error := strconv.ParseBool(fmt.Sprintf("%v", source))
	if error == nil {
		return v, nil
	}
	return nil, error
}

func Str2Float64(source string) (any, error) {
	v, error := strconv.ParseFloat(fmt.Sprintf("%v", source), 10)
	if error == nil {
		return v, nil
	}
	return nil, error
}

func Str2Float32(source string) (any, error) {
	v, error := strconv.ParseFloat(fmt.Sprintf("%v", source), 10)
	if error == nil {
		return float32(v), nil
	}
	return nil, error
}

func Str2Object(v string, k reflect.Kind) (any, error) {
	if len(v) == 0 {
		return nil, nil
	}

	switch k {
	case reflect.String:
		return v, nil
	case reflect.Int:
		return Str2Int(v)
	case reflect.Int16:
		return Str2Int16(v)
	case reflect.Int32:
		return Str2Int32(v)
	case reflect.Int64:
		return Str2Int64(v)
	case reflect.Int8:
		return Str2Int8(v)
	case reflect.Uint:
		return Str2Uint(v)
	case reflect.Uint8:
		return Str2Uint8(v)
	case reflect.Uint16:
		return Str2Uint16(v)
	case reflect.Uint32:
		return Str2Uint32(v)
	case reflect.Uint64:
		return Str2UInt64(v)
	case reflect.Bool:
		return Str2Bool(v)
	case reflect.Float32:
		return Str2Float32(v)
	case reflect.Float64:
		return Str2Float64(v)
	}

	return nil, nil
}

func Reduce[T any, R any](source []T, fn func(one T) (R, bool)) []R {
	if source == nil {
		return nil
	}

	rtn := make([]R, 0, len(source))

	for _, o := range source {
		if v, remain := fn(o); remain {
			rtn = append(rtn, v)
		}
	}

	return rtn
}

// utils
func Type2Str(t reflect.Type) (string, error) {
	if t.Kind() == reflect.Struct || t.Kind() == reflect.Interface {
		return t.String(), nil
	} else if t.Kind() == reflect.Ptr {
		return t.Elem().String(), nil
	}

	return t.String(), errors.New(t.String() + " is not struct or interface")
}

var _expression_reg = regexp.MustCompile(`\{(?s:(.*?))\}`)

func ParseParameterName(str string) []string {
	result := _expression_reg.FindAllStringSubmatch(str, -1)
	l := len(result)
	if l == 0 {
		return nil
	}

	if l == 1 {
		return []string{result[0][1]}
	}

	yet := make(map[string]bool)

	ks := Reduce(result, func(one []string) (string, bool) {
		if one != nil {
			if len(one[1]) > 0 {
				if _, ok := yet[one[1]]; !ok {
					return one[1], true
				}
			}
		}

		return "", false
	})

	return ks
}

const _KEY_1 = "{"
const _KEY_2 = "}"

func ReplaceParameterValue(str string, keyAndValue map[string]string) string {
	keys := ParseParameterName(str)

	if len(keys) == 0 {
		return str
	}

	for _, k := range keys {
		if v, ok := keyAndValue[k]; ok {
			str = strings.ReplaceAll(str, _KEY_1+k+_KEY_2, v)
		} else {
			//str = strings.ReplaceAll(str, k, v)
		}
	}

	return str
}

func ReplaceParameterWithKeyValue(str string, keyAndValue map[string]string) string {
	if len(keyAndValue) == 0 {
		return str
	}

	for k, v := range keyAndValue {
		str = strings.ReplaceAll(str, k, v)
	}

	return str
}
