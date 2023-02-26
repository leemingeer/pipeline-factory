package test

import (
	"github.com/leemingeer/pipeline-factory/pkg/pipeline"
	"testing"
)

func TestPipeline(t *testing.T) {
	// 其中pipeline.DefaultConfig()的配置内容见【抽象工厂模式示例图】
	// 消息处理流程为 HelloInput -> UpperFilter -> ConsoleOutput
	p := pipeline.Of(pipeline.DefaultConfig())
	p.Exec()
}
