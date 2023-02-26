package output

import (
	"fmt"
)

// Console output插件，将消息输出到控制台上
type ConsoleOutput struct{}

func (c *ConsoleOutput) Send(msg string) {
	fmt.Println(msg)
}
