package pipeline

import "github.com/leemingeer/pipeline-factory/pkg/pipeline/plugin"

type Config struct {
	Input  plugin.PluginConfig
	Filter plugin.PluginConfig
	Output plugin.PluginConfig
}

func DefaultConfig() Config {
	c := Config{}
	return c.WithOutput().WithInput().WithFilter()
}

func (c Config) WithInput() Config {
	c.Input = plugin.PluginConfig{
		Name:       "hello",
		PluginType: "input",
	}
	return c
}
func (c Config) WithFilter() Config {
	c.Filter = plugin.PluginConfig{
		Name:       "upper",
		PluginType: "filter",
	}
	return c
}
func (c Config) WithOutput() Config {
	c.Output = plugin.PluginConfig{
		Name:       "console",
		PluginType: "output",
	}
	return c
}
