package examples

import (
	"fmt"
	. "github.com/gohutool/boot4go"
	. "github.com/gohutool/boot4go/configuration"
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

func TestContext(t *testing.T) {
	fmt.Println(Context)
}

func TestConfiguration(t *testing.T) {
	fmt.Println("Start")
	ConfigurationContext.Put("com.joinsunsoft.name", "liuyong")
	ConfigurationContext.Put("com.joinsunsoft.age", "10")
	ConfigurationContext.Put("com.joinsunsoft.name", "DavidLiu")

	fmt.Println(ConfigurationContext.GetValue("${test}"))
	fmt.Println(ConfigurationContext.GetValue("${tag.hello}"))

	fmt.Println("DATA :", ConfigurationContext.GetValue("${com.joinsunsoft}"))

	time.Sleep(2 * time.Second)
}
