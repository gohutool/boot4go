package boot4go

import (
	. "github.com/gohutool/expression4go"
	. "github.com/gohutool/expression4go/spel"
	"github.com/gohutool/log4go"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

/**
* golang-sample源代码，版权归锦翰科技（深圳）有限公司所有。
* <p>
* 文件名称 : configuration.go
* 文件路径 :
* 作者 : DavidLiu
× Email: david.liu@ginghan.com
*
* 创建日期 : 2022/4/8 19:43
* 修改历史 : 1. [2022/4/8 19:43] 创建文件 by NST
*/

type configurationContext struct {
	data map[string]any
}

var ConfigurationContext = configurationContext{data: make(map[string]any)}

var configLogger = log4go.LoggerManager.GetLogger("boot4go.configuration")

func (c configurationContext) IsConfigFileExist(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func (c configurationContext) toMap() map[string]any {
	return c.data
}

func (c configurationContext) GetValue(expression string) any {
	context := StandardEvaluationContext{}
	context.AddPropertyAccessor(MapAccessor{})
	context.SetVariables(c.toMap())
	parser := SpelExpressionParser{}

	return parser.ParseExpression(expression).GetValueContext(&context)
}

func (c configurationContext) LoadYaml(filename string) {

	fd, err := os.Open(filename)
	if err != nil {
		e := configLogger.Error("LoadYaml: Error: Could not open %q for reading: %s", filename, err)
		panic(e.Error())
	}
	contents, err := ioutil.ReadAll(fd)
	if err != nil {
		e := configLogger.Error("LoadYaml: Error: Could not read %q: %s", filename, err)
		panic(e.Error())
	}

	err = yaml.Unmarshal([]byte(contents), ConfigurationContext.data)
	if err != nil {
		e := configLogger.Error("LoadYaml: Error: Could not read %q: %s", filename, err)
		panic(e.Error())
	}
}
