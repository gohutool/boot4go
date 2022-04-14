package configuration

import (
	"errors"
	"fmt"
	"github.com/gohutool/boot4go/util"
	. "github.com/gohutool/expression4go"
	. "github.com/gohutool/expression4go/spel"
	"github.com/gohutool/log4go"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
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

func init() {
	ConfigurationContext.LoadEnv()
}

type configurationContext struct {
	data map[string]any
}

var ConfigurationContext = configurationContext{data: make(map[string]any)}

var configLogger = log4go.LoggerManager.GetLogger("boot4go.configuration")

func (c configurationContext) LoadEnv() {
	envs := os.Environ()
	for _, env := range envs {
		idx := strings.Index(env, "=")
		if idx >= 1 && idx < len(env)-1 {
			c.Put(util.Substring(env, 0, idx), util.Substring(env, idx+1, -1))
		} else {
			c.Put(env, "")
		}

	}
}

func (c configurationContext) Put(key string, value any) {
	ks := strings.Split(key, ".")
	pm := c.data
	l := len(ks)
	ok := false

	for idx, k := range ks {

		if l == idx+1 {
			pm[k] = value
		} else {
			m, exist := pm[k]

			if exist {
				m, ok = interface{}(m).(map[string]any)
				if !ok {
					pm[k] = make(map[string]any)
					/*configLogger.Debug("%v is exist with %v, will be override with list", strings.Join(ks[:idx], "."), m)
					pm[k] = make([]any, 10)
					//pm[k]
					pm[k] = append(pm[k].([]any), m)
					pm[k] = append(pm[k].([]any), make(map[string]any))*/
				}
			} else {
				pm[k] = make(map[string]any)
			}

			pm = pm[k].(map[string]any)
		}
	}
}

func (c configurationContext) IsConfigFileExist(filename string) bool {
	_, err := os.Stat(filename)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

func (c configurationContext) ToMap() map[string]any {
	return c.data
}

func (c configurationContext) GetValue2(expression string) (rtn any, err error) {
	context := StandardEvaluationContext{}
	context.AddPropertyAccessor(MapAccessor{})
	context.SetVariables(c.ToMap())
	parser := SpelExpressionParser{}

	defer func() {
		if e := recover(); e != nil {
			err = errors.New(e.(string))
			rtn = nil
		}
	}()

	rtn = parser.ParseExpression(expression).GetValueContext(&context)
	err = nil
	return
}

var _expression_reg = regexp.MustCompile(`\$\{(?s:(.*?))\}`)

func (c configurationContext) GetValue(expression string) any {

	//strings.Split(expression)

	result := _expression_reg.FindAllStringSubmatch(expression, -1)
	l := len(result)
	if l == 0 {
		return expression
	}

	if l == 1 && strings.Index(expression, "$") == 0 && strings.Index(expression, "}") == len(expression)-1 {
		if v, err := c.GetValue2(expression); err == nil {
			return v
		} else {
			return nil
		}
	}

	yet := make(map[string]bool)

	ks := util.Reduce(result, func(one []string) (string, bool) {
		if one != nil {
			if len(one[1]) > 0 {
				if _, ok := yet[one[1]]; !ok {
					return one[1], true
				}
			}
		}

		return "", false
	})

	//过滤<></>
	for _, text := range ks {
		var value any
		olds := "${" + text + "}"
		if v, err := c.GetValue2(olds); err == nil {
			value = v
		} else {
			value = ""
		}

		expression = strings.ReplaceAll(expression, olds, fmt.Sprintf("%v", value))
	}

	return expression

	/*context := StandardEvaluationContext{}
	context.AddPropertyAccessor(MapAccessor{})
	context.SetVariables(c.ToMap())
	parser := SpelExpressionParser{}

	return parser.ParseExpression(expression).GetValueContext(&context)*/
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
