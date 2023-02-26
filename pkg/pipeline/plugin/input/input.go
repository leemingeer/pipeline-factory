package input

// Hello input插件，接收“Hello World”消息
type HelloInput struct{}

func (h *HelloInput) Receive() string {
	return "Hello World"
}
