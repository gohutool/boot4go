package boot4go

import "reflect"

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

type AutoWiredMeta struct {
	Type reflect.Type
	bean any
}

type AutowiredBeanHandler interface {
	BeforeAutowired(meta AutoWiredMeta) any
	PostAutowired(source any, meta AutoWiredMeta)
}
