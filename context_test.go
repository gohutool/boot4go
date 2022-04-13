package boot4go

import (
	"fmt"
	"github.com/gohutool/log4go"
	"reflect"
	"testing"
	"time"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : context_test.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/7 10:01
* 修改历史 : 1. [2022/4/7 10:01] 创建文件 by NST
*/

var logger log4go.Logger
var bean any

func init() {
	log4go.LoggerManager.InitWithDefaultConfig()

	d, _ := Context.RegistryBean("aaa", Hello{})
	fmt.Println(d)

	//bean, _ = Context.GetBean(Test{})
	//Context.RegistryBeanInstance("boot4go.IHello", d)
	/*
		h := &Hello{}
		Test{hello2: h.(IHello)}*/

	logger = log4go.LoggerManager.GetLogger("boot4go.context.test")
}

type Test struct {
	age     int16          `bootable:"${metadata.major}"`
	name    string         `bootable:"${metadata.name}"`
	version string         `bootable:"${metadata.version}"`
	hello   IHello         `bootable:"aaa"`
	hello2  Hello2         `bootable`
	data    map[string]any `bootable:"${spec.runAsUser}"`
	list    []any          `bootable:"${spec.volumes}"`
}

func (t *Test) PostConstruct() {
	logger.Info("PostConstruct Test")
}

func (t *Test) sayHello2(w string) string {
	return t.hello2.sayHello(w)
}

func (t *Test) sayHello(w string) string {
	return t.hello.sayHello(w)
}

type Hello2 struct {
	tag string `bootable:"${tag.hello2}"`
}

func (h *Hello2) sayHello(t string) string {
	return "Hello2 " + h.tag + " : " + t
}

func (t *Hello2) PostConstruct() {
	logger.Info("PostConstruct Hello2")
}

type Hello struct {
	tag string `bootable:"${tag.hello}"`
}

func (h *Hello) sayHello(t string) string {
	return "Hello " + h.tag + " : " + t
}

func (t *Hello) PostConstruct() {
	logger.Info("PostConstruct Hello")
	panic("panic testing")
}

type IHello interface {
	sayHello(t string) string
}

func TestContext(t *testing.T) {
	fmt.Println(log4go.LoggerManager)

	test := &Test{}
	typeof := reflect.TypeOf(test)

	fmt.Println(typeof.String())
	fmt.Println(typeof.Kind().String())

	fmt.Println(type2Str(reflect.TypeOf(test)))
	fmt.Println(type2Str(reflect.TypeOf(*test)))

	h, ok := interface{}(test).(IHello)

	fmt.Println(ok)

	var ih IHello = &Hello{}

	fmt.Println(type2Str(reflect.TypeOf(h)))
	fmt.Println(type2Str(reflect.TypeOf(ih)))

	fmt.Println(IHello.sayHello(ih, "boot4go"))

	h.sayHello("boot4go")
}

func TestGetBean(t *testing.T) {
	bean, _ = Context.GetBean(Test{})
	t1 := bean.(*Test)
	logger.Info(t1.hello)
	logger.Info("Hello2=" + t1.sayHello2("AAA"))
	logger.Info("Hello=" + t1.sayHello("AAA"))
	logger.Info("%v", &t1.hello)

	time.Sleep(2 * time.Second)
}

func TestGetBeanByName(t *testing.T) {
	bean, _ := Context.getBeanByName("boot4go.Test")
	t1 := bean.(*Test)
	logger.Info(t1.hello)
	logger.Info("Hello2=" + t1.sayHello2("BBB"))
	logger.Info("Hello=" + t1.sayHello("BBB"))
	logger.Info("%v", &t1.hello)

	logger.Info("%v", &t1.data)
	logger.Info("%v", &t1.list)

	bean, _ = Context.getBeanByName("boot4go.Hello2")
	h2 := bean.(*Hello2)
	logger.Info("%s", h2.sayHello("CCC"))
}

func TestContextConfiguration(t *testing.T) {

	bean, ok := Context.GetBean(&Test{})
	fmt.Println(reflect.TypeOf(bean.(*Test)).String(), bean, ok)

	logger := log4go.LoggerManager.GetLogger("test")

	logger.Info("YAML %v", ConfigurationContext.ToMap())

	logger.Info("YAML %v", ConfigurationContext.GetValue("${metadata.name}"))
	logger.Info("YAML %v", ConfigurationContext.GetValue("${spec.volumes[0]}"))

	time.Sleep(10 * time.Second)
}
