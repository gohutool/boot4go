package boot4go

import (
	. "github.com/gohutool/boot4go/reflect"
	. "reflect"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : autowire_handler.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/13 16:29
* 修改历史 : 1. [2022/4/13 16:29] 创建文件 by LongYong
*/

type SupportType Kind

const (
	CLASS  = Struct
	METHOD = Func
	FIELD  = String
)

type AutoWiredMeta struct {
	Value Value
	Tag   StructTagPlus
	Type  Type
	bean  any
}

type AutowiredBeanHandler interface {
	ID() string
	BeforeAutowired(meta AutoWiredMeta) any
	PostAutowired(source any, meta AutoWiredMeta)
	SupportType() []Kind
}

const AUTOCONFIG_AUTOWIRED_TAG = "auto"

type AutoConfigurationAutowiredBeanHandler struct {
}

func (s *AutoConfigurationAutowiredBeanHandler) ID() string {
	return AUTOCONFIG_AUTOWIRED_TAG
}

func (s *AutoConfigurationAutowiredBeanHandler) BeforeAutowired(meta AutoWiredMeta) any {
	bn, _ := meta.Tag.Get(AUTOWIRED_FLAG + AUTOCONFIG_AUTOWIRED_TAG)

	k := meta.Type.Kind()
	if len(bn) == 0 {
		if k == Ptr {
			bn = meta.Value.Elem().Type().String()
		} else {
			bn = meta.Value.Type().String()
		}
	}

	var v any
	if b, _ := Context.pooled[bn]; b != nil {
		v = b
	} else {
		if k == Ptr {
			v, _ = Context.getBeanByType(meta.Value.Elem().Type())
		} else {
			v, _ = Context.getBeanByType(meta.Value.Type())
		}
	}

	//fmt.Printf("%v %v %v\n", meta.Type, meta.bean, meta.Tag)
	return v
}

func (s *AutoConfigurationAutowiredBeanHandler) PostAutowired(source any, meta AutoWiredMeta) {

}

func (s *AutoConfigurationAutowiredBeanHandler) SupportType() []Kind {
	return []Kind{CLASS, METHOD, FIELD}
}
