package plugin

import (
	"github.com/leemingeer/pipeline-factory/pkg/pipeline/plugin/filter"
	"github.com/leemingeer/pipeline-factory/pkg/pipeline/plugin/input"
	"github.com/leemingeer/pipeline-factory/pkg/pipeline/plugin/output"
	"reflect"
)

type PluginConfig struct {
	Name       string
	PluginType string
}

// 插件抽象接口定义
type Plugin interface{}

// 输入插件，用于接收消息
type Input interface {
	Plugin
	Receive() string
}

// 过滤插件，用于处理消息
type Filter interface {
	Plugin
	Process(msg string) string
}

// 输出插件，用于发送消息
type Output interface {
	Plugin
	Send(msg string)
}

// 插件名称与类型的映射关系，主要用于通过反射创建output对象
var (
	filterNames = make(map[string]reflect.Type)
	inputNames  = make(map[string]reflect.Type)
	outputNames = make(map[string]reflect.Type)
)

// 初始化插件映射关系表
func init() {
	inputNames["hello"] = reflect.TypeOf(input.HelloInput{})
	filterNames["upper"] = reflect.TypeOf(filter.UpperFilter{})
	outputNames["console"] = reflect.TypeOf(output.ConsoleOutput{})
}
