package osexec

var debugModeOpen = false

// SetDebugMode sets the package-wide debug mode
// SetDebugMode 设置包级别的调试模式
func SetDebugMode(enable bool) {
	debugModeOpen = enable
}

// GetDebugMode returns the package-wide debug mode status
// GetDebugMode 返回包级别的调试模式状态
func GetDebugMode() bool {
	return debugModeOpen
}
