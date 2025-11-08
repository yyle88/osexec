package osexec

// CMC is an alias representing CommandConfig
// CMC 是 CommandConfig 的别名
type CMC = CommandConfig

// NewCMC creates and returns a new CMC instance
// NewCMC 创建并返回新的 CMC 实例
func NewCMC() *CMC {
	return NewCommandConfig()
}

// OsCommand is an alias representing CommandConfig
// OsCommand 是 CommandConfig 的别名
type OsCommand = CommandConfig

// NewOsCommand creates and returns a new OsCommand instance
// NewOsCommand 创建并返回新的 OsCommand 实例
func NewOsCommand() *OsCommand {
	return NewCommandConfig()
}

// ExecConfig is an alias representing CommandConfig
// ExecConfig 是 CommandConfig 的别名
type ExecConfig = CommandConfig

// NewExecConfig creates and returns a new ExecConfig instance
// NewExecConfig 创建并返回新的 ExecConfig 实例
func NewExecConfig() *ExecConfig {
	return NewCommandConfig()
}

var debugModeOpen = false

// SetDebugMode sets the package-wide debug mode
// SetDebugMode 设置包级别的调试模式
func SetDebugMode(enable bool) {
	debugModeOpen = enable
}
