package boot4go

import (
	"fmt"
	"github.com/gohutool/boot4go/configuration"
	"github.com/gohutool/log4go"
	. "reflect"
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

	Context.RegistryAutowiredBeanHandler(new(SampleAutowiredHandler))

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
	age     int16          `@auto:"${metadata.major}"`
	name    string         `@auto:"${metadata.name}"`
	version string         `@auto:"${metadata.version}"`
	hello   IHello         `@auto:"aaa"`
	hello2  Hello2         `@auto`
	hello3  Hello2         `@sample`
	data    map[string]any `@auto:"${spec.runAsUser}"`
	list    []any          `@auto:"${spec.volumes}"`
}

func (t *Test) PostConstruct() {
	logger.Info("PostConstruct Test")
}

func (t *Test) sayHello2(w string) string {
	return t.hello2.sayHello(w)
}

func (t *Test) sayHello3(w string) string {
	return t.hello3.sayHello(w)
}

func (t *Test) sayHello(w string) string {
	return (t.hello).sayHello(w)
}

type Hello2 struct {
	tag string `@auto:"${tag.hello2}  ${kind}"`
}

func (h *Hello2) sayHello(t string) string {
	return "hello2 " + h.tag + " : " + t
}

func (t *Hello2) PostConstruct() {
	logger.Info("PostConstruct hello2")
}

type Hello struct {
	tag string `@auto:"${tag.hello} ${kind}"`
}

func (h *Hello) sayHello(t string) string {
	return "Hello " + h.tag + " : " + t
}

func (t *Hello) PostConstruct() {
	logger.Info("PostConstruct Hello")
	//panic("panic testing")
}

type IHello interface {
	sayHello(t string) string
}

type SampleAutowiredHandler struct {
}

func (s *SampleAutowiredHandler) BeforeAutowired(meta AutoWiredMeta) any {
	logger.Info("%v %v %v", meta.Type, meta.Bean, meta.Tag)
	return Hello2{tag: "SampleAutowiredHandler autowired"}
}

func (s *SampleAutowiredHandler) PostAutowired(source any, meta AutoWiredMeta) {

}

func (s *SampleAutowiredHandler) ID() string {
	return "sample"
}

func (s *SampleAutowiredHandler) SupportType() []Kind {
	return []Kind{CLASS, METHOD, FIELD}
}

func TestContext(t *testing.T) {
	fmt.Println(log4go.LoggerManager)

	test := &Test{}
	typeof := TypeOf(test)

	fmt.Println(typeof.String())
	fmt.Println(typeof.Kind().String())

	fmt.Println(type2Str(TypeOf(test)))
	fmt.Println(type2Str(TypeOf(*test)))

	h, ok := interface{}(test).(IHello)

	fmt.Println(ok)

	var ih IHello = &Hello{}

	fmt.Println(type2Str(TypeOf(h)))
	fmt.Println(type2Str(TypeOf(ih)))

	fmt.Println(IHello.sayHello(ih, "boot4go"))

	h.sayHello("boot4go")
}

func TestGetBean(t *testing.T) {
	bean, _ = Context.GetBean(Test{})
	t1 := bean.(*Test)
	logger.Info(t1.hello)
	logger.Info("hello2=" + t1.sayHello2("AAA"))
	logger.Info("Hello=" + t1.sayHello("AAA"))
	logger.Info("Hello3=" + t1.sayHello3("AAA"))
	logger.Info("%v", &t1.hello)

	time.Sleep(2 * time.Second)
}

func TestGetBeanByName(t *testing.T) {
	bean, _ := Context.getBeanByName("boot4go.Test")
	t1 := bean.(*Test)
	logger.Info(t1.hello)
	logger.Info("hello2=" + t1.sayHello2("BBB"))
	logger.Info("Hello=" + t1.sayHello("BBB"))
	logger.Info("%v", &t1.hello)

	logger.Info("%v", &t1.data)
	logger.Info("%v", &t1.list)

	bean, _ = Context.getBeanByName("boot4go.hello2")
	h2 := bean.(*Hello2)
	logger.Info("%s", h2.sayHello("CCC"))
}

func TestContextConfiguration(t *testing.T) {

	bean, ok := Context.GetBean(&Test{})
	fmt.Println(TypeOf(bean.(*Test)).String(), bean, ok)

	logger := log4go.LoggerManager.GetLogger("test")

	logger.Info("YAML %v", configuration.ConfigurationContext.ToMap())

	logger.Info("YAML %v", configuration.ConfigurationContext.GetValue("${metadata.name}"))
	logger.Info("YAML %v", configuration.ConfigurationContext.GetValue("${spec.volumes[0]}"))

	time.Sleep(10 * time.Second)
}
