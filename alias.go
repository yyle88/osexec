package osexec

// ExecConfig is an alias representing CommandConfig
// ExecConfig 是 CommandConfig 的别名
type ExecConfig = CommandConfig

// NewExecConfig creates and returns a new ExecConfig instance
// NewExecConfig 创建并返回新的 ExecConfig 实例
func NewExecConfig() *ExecConfig {
	return NewCommandConfig()
}
