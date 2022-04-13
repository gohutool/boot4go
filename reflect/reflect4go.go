package reflect4go

import (
	"reflect"
	"strings"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : reflect4go.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/13 18:03
* 修改历史 : 1. [2022/4/13 18:03] 创建文件 by LongYong
*/

type StructTagPlus struct {
	reflect.StructTag
}

// GetKeys
// 先简单实现，不考虑解析的性能和表达式的多样性
func (tag StructTagPlus) GetKeys() []string {
	if len(strings.TrimSpace(string(tag.StructTag))) == 0 {
		return []string{}
	}

	segment := strings.Split(string(tag.StructTag), " ")

	rtn := make([]string, 0, 10)

	for _, s := range segment {
		s = strings.TrimSpace(s)
		ss := strings.Split(s, ":")
		if len(ss[0]) > 0 {
			rtn = append(rtn, ss[0])
		}
	}

	return rtn
}

// isExistKey
// 先简单实现，不考虑解析的性能和表达式的多样性
func (tag StructTagPlus) isExistKey(key string) bool {
	str := string(tag.StructTag)
	if len(strings.TrimSpace(str)) == 0 {
		return false
	}

	if str == key || str == key+" " || str == " "+key {
		return true
	}

	return strings.Index(str, key+":") >= 0
}

// Get
// 先简单实现，不考虑解析的性能和表达式的多样性
func (tag StructTagPlus) Get(key string) (string, bool) {
	if tag.isExistKey(key) {
		return tag.StructTag.Get(key), true
	}

	return "", false
}

// GetAutowirdValue
// 先简单实现，不考虑解析的性能和表达式的多样性
func (tag StructTagPlus) GetAutowirdValue(key string) (string, bool) {
	if tag.isExistKey(AUTOWIRED_FLAG + key) {
		return tag.StructTag.Get(AUTOWIRED_FLAG + key), true
	}

	return "", false
}

const AUTOWIRED_FLAG = "@"

// GetAutowiredTag
// 先简单实现，不考虑解析的性能和表达式的多样性
func (tag StructTagPlus) GetAutowiredTag(key string) (string, bool) {
	if strings.Index(key, AUTOWIRED_FLAG) == 0 {
		return strings.Replace(key, AUTOWIRED_FLAG, "", 1), true
	} else {
		return "", false
	}
}
