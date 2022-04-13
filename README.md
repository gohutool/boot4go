# Boot4go

*a framework to start a web application quickly like as spring-boot*

![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)

## INTRO

Boot4go helps you to create production-grade applications and services with absolute minimum fuss. It takes an opinionated view of the Spring platform so that new and existing users can quickly get to the bits they need.

You can use Boot4go to create stand-alone applications that can be started.

Our primary goals are:

- Provide a radically faster and widely accessible getting started experience for all Spring development.

- Be opinionated, but get out of the way quickly as requirements start to diverge from the defaults.

- Provide a range of non-functional features common to large classes of projects (for example, embedded servers, security, metrics, health checks, externalized configuration).

- Absolutely no code generation and no requirement for XML configuration.

## Installation and Getting Started

- Import
```
import (
	"fmt"
	"github.com/gohutool/log4go"
	"github.com/gohutool/boot4go"
)
```

The reference documentation includes detailed installation instructions as well as a comprehensive getting started guide.

Code example:

- Autoconfiguration

```
apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: mysql-snapshot
  version: v1.0.1
  major: 1
tag:
  hello: TAG-HELLO-A
  hello2: TAG-HELLO-B
spec:
  privileged: false
  allowPrivilegeEscalation: false
  volumes:
    - "*"
    - "*.json"
  hostNetwork: false
  hostIPC: false
  hostPID: false
  runAsUser:
    rule1: RunAsAny1
    rule2: RunAsAny2
    rule3: RunAsAny3
  seLinux:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  fsGroup:
    rule: RunAsAny
```

- Test struct 

Sample code will autowire test struct, it will autofill the properties using the yaml's configuration
```
type Test struct {
	age     int16          `bootable:"${metadata.major}"`
	name    string         `bootable:"${metadata.name}"`
	version string         `bootable:"${metadata.version}"`
	data    map[string]any `bootable:"${spec.runAsUser}"`
	list    []any          `bootable:"${spec.volumes}"`
}
```

- Test autowired properties with expression string

```
func TestContextConfiguration(t *testing.T) {
	logger := log4go.LoggerManager.GetLogger("test")
	
	logger.Info("YAML %v", ConfigurationContext.ToMap())

	logger.Info("YAML %v", ConfigurationContext.GetValue("${metadata.name}"))
	logger.Info("YAML %v", ConfigurationContext.GetValue("${spec.volumes[0]}"))

	time.Sleep(10 * time.Second)
}
```

- Output
```
[19:11:28 CST 2022/04/09 677] [INFO][test] (github.com/gohutool/boot4go.TestContextConfiguration:118) YAML map[=::=::\: ALLUSERSPROFILE:C:\ProgramData APPCODE_VM_OPTIONS:E:\BaiduNetdiskDownload\JetBrains 2022 最 新版本全家桶激活\方式2：激活到2099年补丁\ja-netfilter-all\vmoptions\appcode.                 vmoptions APPDATA:C:\Users\NST\AppData\Roaming AR:ar CC:gcc CGO_CFLAGS:-O0 -g CGO_ENABLED:1 CLION_VM_OPTIONS:E:\BaiduNetdiskDownload\JetBrains 2022 最新版   本全家桶激活\方式2：激活到2099年补丁\ja-netfilter-all\vmoptions\clion.vmopti               ons COMPUTERNAME:NST-PC CXX:g++ ComSpec:C:\WINDOWS\system32\cmd.exe CommonProgramFiles:C:\Program Files\Common Files CommonProgramFiles(x86):C:\Program Files (x86)\Common Files CommonProgramW6432:C:\Program Files\Common Files DATAGRIP_VM_OPTIONS:E:\BaiduNetdiskDownload\JetBrains 2022 最新版本全家桶激活\方式           2：激活到2099年补丁\ja-netfilter-all\vmoptions\datagrip.vmoptions DATASPELL_V       M_OPTIONS:E:\BaiduNetdiskDownload\JetBrains 2022 最新版本全家桶激活\方式2：激活到2099年补丁\ja-netfilter-all\vmo            ptions\dataspell.vmoptions DriverData:C:\Windows\System32\Drivers\DriverData FPS_BROWSER_APP_PROFILE_STRING:Internet Explorer FPS_BROWSER_USER_PROFILE_STRING:Default GATEWAY_VM_OPTIONS:E:\BaiduNetdiskDownload\JetBrains 2022 最新版本     全家桶激活\方式2：激活到2099年补丁\ja-netfilter-all\vmoptions\gateway.vmopti              ons GCCGO:gccgo GO111MODULE:on GOAMD64:v1 GOARCH:amd64 GOCACHE:C:\Users\NST\AppData\Local\go-build GOENV:C:\Users\NST\AppData\Roaming\go\env GOEXE:.exe GOHOSTARCH:amd64 GOHOSTOS:windows GOLAND_VM_OPTIONS:E:\BaiduNetdiskDownload\JetBrains 2022 最新版本全家桶激活\方式2：激活到2099年补丁\ja-netfilter-all\v                  moptions\goland.vmoptions GOMODCACHE:E:\WORK\SOFT\go1.18.windows-amd64\go\pkg\mod GOOS:windows GOPATH:E:\WORK\SOFT\go1.18.windows-amd64\go GOPROXY:https://goproxy.cn,direct GOROOT:E:\WORK\SOFT\go1.18.windows-amd64\go GOSUMDB:sum.golang.org GOTOOLDIR:E:\WORK\SOFT\go1.18.windows-amd64\go\pkg\tool\windows_amd64 GOVERSION:go1.18 GoLand:E:\WORK\SOFT\JetBrains\GoLand 2021.3.3\bin; HOMEDRIVE:C: HOMEPATH:\Users\NST IDEA_INITIAL_DIRECTORY:C:\WINDOWS\System32 IDEA_VM_OPTIONS:E:\BaiduNetdiskDownload\JetBrains 2022 最新版    本全家桶激活\方式2：激活到2099年补丁\ja-netfilter-all\vmoptions\idea.vmoptio               ns JAVA_HOME:E:\WORK\SOFT\JDK8-64 JETBRAINSCLIENT_VM_OPTIONS:E:\BaiduNetdiskDownload\JetBrains 2022 最新版本全家桶激活\方式2：激活到2099年补丁\ja-netfilt                  er-all\vmoptions\jetbrainsclient.vmoptions JETBRAINS_CLIENT_VM_OPTIONS:E:\BaiduNetdiskDownload\JetBrains 2022 最新版本全家桶激活\方式2：激活到2099年补丁\j                  a-netfilter-all\vmoptions\jetbrains_client.vmoptions LOCALAPPDATA:C:\Users\NST\AppData\Local LOGONSERVER:\\NST-PC MAVEN_HOME:E:\WORK\SOFT\maven-3.6.1 MONGO_HOME:E:\WORK\SOFT\mongodb-win32-x86_64-windows-5.0.3\ MOZ_PLUGIN_PATH:D:\Program Files (x86)\Foxit Software\Foxit PhantomPDF\plugins\ NPM_PREFIX:E:\WORK\SOFT\nodejs\node_global NUMBER_OF_PROCESSORS:8 OS:Windows_NT OneDrive:C:\Use
[21:40:01 CST 2022/04/09 423] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:117) mysql-snapshot
[21:40:01 CST 2022/04/09 423] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:120) [* *.json]
```

- Test autofilled properties with Test struct autowired

```
func TestGetBean(t *testing.T) {
	bean, _ := Context.GetBean(Test{})

	logger.Info("%v", &t1.data)
	logger.Info("%v", &t1.list)
	logger.Info("%v", &t1.age)
	logger.Info("%v", &t1.name)
	logger.Info("%v", &t1.version)		
}
```

- Output
```
[21:40:01 CST 2022/04/09 424] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:129) &map[rule1:RunAsAny1 rule2:RunAsAny2 rule3:RunAsAny3]
[21:40:01 CST 2022/04/09 424] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:130) &[* *.json]
[21:40:01 CST 2022/04/09 424] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:130) 1
[21:40:01 CST 2022/04/09 424] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:130) mysql-snapshot
[21:40:01 CST 2022/04/09 424] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:130) v1.0.1
```

- Autowire with nested struct

```
type Test struct {
	hello   IHello         `bootable:"aaa"`
	hello2  Hello2         `bootable`
	data    map[string]any `bootable:"${spec.runAsUser}"`
	list    []any          `bootable:"${spec.volumes}"`
}

func (t *Test) sayHello2(w string) string {
	return t.hello2.sayHello(w)
}

func (t *Test) sayHello(w string) string {
	return t.hello.sayHello(w)
}

type IHello interface {
	sayHello(t string) string
}

type Hello struct {
	tag string `bootable:"${tag.hello}"`
}

type Hello2 struct {
	tag string `bootable:"${tag.hello2}"`
}

func (h *Hello) sayHello(t string) string {
	return "Hello " + h.tag + " : " + t
}

func (h *Hello2) sayHello(t string) string {
	return "Hello2 " + h.tag + " : " + t
}

```

- Code

```
func init() {
	log4go.LoggerManager.InitWithDefaultConfig()
	d, _ := Context.RegistryBean("aaa", Hello{})
	logger = log4go.LoggerManager.GetLogger("boot4go.context.test")
}

func TestGetBean(t *testing.T) {
    // Autowired Bean
	bean, _ := Context.GetBean(Test{})

	t1 := bean.(*Test)
	logger.Info(t1.hello)
	logger.Info("Hello2=" + t1.sayHello2("AAA"))
	logger.Info("Hello=" + t1.sayHello("AAA"))
}
```

- Output
```
=== RUN   TestGetBean
[21:40:01 CST 2022/04/09 423] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:117) 0xc00013bbb8
[21:40:01 CST 2022/04/09 423] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:118) Hello2=Hello2 TAG-HELLO-B : AAA
[21:40:01 CST 2022/04/09 424] [INFO][boot4go.context.test] (github.com/gohutool/boot4go.TestGetBean:126) Hello=Hello TAG-HELLO-A : AAA
--- PASS: TestGetBean (10.00s)
PASS


Debugger finished with the exit code 0
```


## Getting Help


## Modules


There are several modules in Spring Boot. Here is a quick overview:

### boot4go

### boot4go-proxy

### boot4go-feign

### boot4go-orm

## LICENCE

Apache License 2.0

