package boot4go

import (
	"errors"
	"github.com/gohutool/log4go"
	"reflect"
	"sync"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : context.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 09:52
* 修改历史 : 1. [2022/4/7 09:52] 创建文件 by NST
*/

var contextLogger = log4go.LoggerManager.GetLogger("boot4go.context")

func init() {
	configFile := "boot4go.yaml"

	if ConfigurationContext.IsConfigFileExist(configFile) {
		ConfigurationContext.LoadYaml(configFile)
	}

	contextLogger.Debug("Yaml %v", ConfigurationContext.ToMap())
}

var Context = context{processing: make(map[string]any), pooled: make(map[string]any)}

type context struct {
	lock       sync.RWMutex
	processing map[string]any
	pooled     map[string]any
}

func (c *context) RegistryBean(t reflect.Type, name string) {
	c.lock.RLock()
	defer c.lock.RUnlock()
}

func (c *context) registryBeanInstance(i any, name string) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	c.pooled[name] = i
}

func (c *context) getBean(obj any) (any, error) {
	t := reflect.TypeOf(obj)
	beanName, _ := type2Str(t)
	//
	//c.lock.RLock()
	//defer c.lock.RUnlock()

	bean, ok := c.resolveBean(beanName)

	if ok == nil {
		return bean, nil
	}

	if t.Kind() == reflect.Struct {
		bean = reflect.New(t).Interface()
	} else {
		bean = reflect.New(t.Elem()).Interface()
	}

	/*
		count := t.NumField()
		for idx := 0; idx < count; idx++ {
			a, ok := type2Str(t.Field(idx).Type)
			fmt.Println(t.Field(idx).Name, " ", a, " ", ok)
		}*/

	c.registryBeanInstance(bean, beanName)

	if t.Kind() == reflect.Struct {
		return bean, nil
	} else {
		return bean, nil
	}

}

func (c *context) resolveBean(name string) (any, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	v, ok := c.pooled[name]
	if ok {
		return v, nil
	}

	return nil, errors.New("Not found \"" + name + "\" bean")
}

// utils

func type2Str(t reflect.Type) (string, error) {
	if t.Kind() == reflect.Struct || t.Kind() == reflect.Interface {
		return t.String(), nil
	} else if t.Kind() == reflect.Ptr {
		return t.Elem().String(), nil
	}

	return "", errors.New(t.String() + " is not struct or interface")
}
