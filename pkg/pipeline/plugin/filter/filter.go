package filter

import (
	"strings"
)

// Upper filter插件，将消息全部字母转成大写
type UpperFilter struct{}

func (u *UpperFilter) Process(msg string) string {
	return strings.ToUpper(msg)
}
