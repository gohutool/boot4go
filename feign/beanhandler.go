package feign

import (
	"fmt"
	. "github.com/gohutool/boot4go"
	. "github.com/gohutool/boot4go-proxy"
	. "github.com/gohutool/boot4go/configuration"
	reflect4go "github.com/gohutool/boot4go/reflect"
	"github.com/gohutool/log4go"
	. "reflect"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : beanhandler.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/14 11:37
* 修改历史 : 1. [2022/4/14 11:37] 创建文件 by LongYong
*/

func init() {
	Context.RegistryAutowiredBeanHandler(new(FeignAutowiredBeanHandler))
}

var FEIGN_AUTOWIRED_TAG = "feign"

type FeignClientMeta struct {
	ID      string
	Path    string
	Uri     string
	Method  string
	Name    string
	Address string
}

type FeignClientPool struct {
	pooled       map[string]FeignClientMeta
	pooledMethod map[string]FeignClientMeta
}

var _feignLogger = log4go.LoggerManager.GetLogger("gohutool.boot4go.feign")

var feignClientPool = FeignClientPool{pooled: make(map[string]FeignClientMeta), pooledMethod: make(map[string]FeignClientMeta)}

type FeignAutowiredBeanHandler struct {
}

func (s *FeignAutowiredBeanHandler) ID() string {
	return FEIGN_AUTOWIRED_TAG
}

func (s *FeignAutowiredBeanHandler) BeforeAutowired(meta AutoWiredMeta) any {
	//logger.Info("%v %v %v", meta.Type, meta.Interface, meta.Tag)

	v := meta.Value

	if v.Kind() == Ptr {
		// if nil new a point
		if !v.Elem().IsValid() {
			v = New(v.Type().Elem())
			meta.Value = v
			meta.Interface = v.Interface()
		}
	} else {
		// if struct new a point and set value
		//newv := New(v.Type())
		//nv := NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr()))
		//newv.Elem().Set(nv.Elem())
		//meta.Value = newv.Elem()
		//meta.Interface = meta.Value.Interface()
		//v = newv
		panic(meta.Object.Elem().Type().String() + "'s field named " + meta.Field.Name + "(" + v.Type().String() + ") is not a pointer of interface struct")
	}

	ptrV := v
	v = ptrV.Elem()
	anyP := ptrV.Interface()

	nf := v.NumField()
	nameF := v.FieldByName("Name")
	addressF := v.FieldByName("Address")
	var name, address string

	if !nameF.IsValid() && !addressF.IsValid() {
		panic("Not found Name or Address field at " + v.Type().String())
	}

	if nameF.IsValid() {
		if f, ok := v.Type().FieldByName("Name"); ok {
			if _v := ConfigurationContext.GetValue(string(f.Tag)); _v != nil {
				name = _v.(string)
			}
		}
	}

	if addressF.IsValid() {
		if f, ok := v.Type().FieldByName("Address"); ok {
			if _v := ConfigurationContext.GetValue(string(f.Tag)); _v != nil {
				address = _v.(string)
			}
		}
	}

	if len(address) == 0 && len(name) == 0 {
		panic("Not set tag value of Name or Address field at " + v.Type().String())
	}

	if len(name) == 0 {
		name = address
	}

	for idx := 0; idx < nf; idx++ {
		f := v.Type().Field(idx)
		vf := v.Field(idx)

		if f.Type.Kind() == Func {

			_feignLogger.Info("Input = %v, Output=%v", vf.Type().NumIn(), vf.Type().NumOut())

			if vf.Type().NumOut() > 1 {
				panic(v.Type().String() + "'s " + f.Name + " parameters is more than one")
			}

			fcm := FeignClientMeta{}
			fcm.Name = name
			fcm.Address = address

			tag := reflect4go.StructTagPlus{f.Tag}

			_feignLogger.Info(tag)

			rm, _ := tag.Get("method")
			path, _ := tag.Get("path")

			fcm.Path = path

			if len(rm) > 0 {
				rm = ConfigurationContext.GetValue(rm).(string)
			}

			if len(rm) == 0 {
				rm = "post"
			}

			fcm.Method = rm

			if len(path) > 0 {
				path = ConfigurationContext.GetValue(path).(string)
			}

			fcm.Uri = path
			fcm.ID = name + "@" + rm + "@" + fcm.Path

			if _, has := feignClientPool.pooled[fcm.ID]; has {
				panic(fcm.ID + " is exist yet, client registry error at " + v.Type().String() + "." + f.Name)
			}

			feignClientPool.pooledMethod[v.Type().String()+"@"+f.Name] = fcm

			_feignLogger.Info("%+v\n", fcm)

			ni := vf.Type().NumIn()
			vft := vf.Type()
			for j := 0; j < ni; j++ {
				pv := vft.In(j)
				_feignLogger.Info("%v=%v", (j + 1), pv.String())
			}

		}
	}

	InvocationProxy.NewProxyInstance(anyP, func(obj any, method InvocationMethod, args []Value) []Value {

		_feignLogger.Info("%v %v %v", v.Type().String(), method.Name, args)
		return []Value{ValueOf(v.Type().String() + " AAAA")}
	})

	fmt.Printf("%v %v %v %v\n", meta.Field.Name, meta.Interface, meta.Tag, InvocationProxy)
	return meta.Interface

	/*bn, _ := meta.Tag.Get(AUTOWIRED_FLAG + AUTOCONFIG_AUTOWIRED_TAG)

	k := meta.Type.Kind()
	if len(bn) == 0 {
		if k == Ptr {
			bn = meta.Value.Elem().Type().String()
		} else {
			bn = meta.Value.Type().String()
		}
	}

	var v any
	if b, _ := Context.GetPooledBean(bn); b != nil {
		v = b
	} else {
		if k == Ptr {
			v, _ = Context.GetBean(meta.Value.Elem().Type())
		} else {
			v, _ = Context.GetBean(meta.Value.Type())
		}
	}

	//fmt.Printf("%v %v %v\n", meta.Type, meta.Interface, meta.Tag)
	return v*/
}

func (s *FeignAutowiredBeanHandler) PostAutowired(source any, meta AutoWiredMeta) {

}

func (s *FeignAutowiredBeanHandler) SupportType() []Kind {
	return []Kind{CLASS, METHOD, FIELD}
}
