package osexec

import (
	"slices"

	"github.com/yyle88/tern"
)

// DebugMode represents the debug output setting when executing commands
// DebugMode 表示执行命令时的调试输出设置
type DebugMode string

//goland:noinspection GoSnakeCaseUsage
const (
	QUIET        DebugMode = "QUIET"        // Quiet mode with no debug output // 静默模式，无任何调试输出
	DEBUG        DebugMode = "DEBUG"        // Complete debug mode with detailed information // 完整调试模式，包含所有调试信息
	SHOW_COMMAND DebugMode = "SHOW_COMMAND" // Show command, not outputs // 显示命令，不显示输出
	SHOW_OUTPUTS DebugMode = "SHOW_OUTPUTS" // Show outputs, not command // 显示输出，不显示命令
)

// NewDebugMode creates a DebugMode based on the debug mode flag. When debugModeOpen is true, returns DEBUG; otherwise, returns QUIET.
// NewDebugMode 根据调试模式标志创建 DebugMode。当 debugModeOpen 为 true 时，返回 DEBUG；否则返回 QUIET。
func NewDebugMode(debugModeOpen bool) DebugMode {
	return tern.BVV(debugModeOpen, DEBUG, QUIET)
}

// isShowCommand checks if the command should be displayed based on the debug mode.
// isShowCommand 检查是否应根据调试模式显示命令。
func isShowCommand(debugMode DebugMode) bool {
	return slices.Contains([]DebugMode{DEBUG, SHOW_COMMAND}, debugMode)
}

// isShowOutputs checks if the command results should be displayed based on the debug mode.
// isShowOutputs 检查是否应根据调试模式显示命令结果。
func isShowOutputs(debugMode DebugMode) bool {
	return slices.Contains([]DebugMode{DEBUG, SHOW_OUTPUTS}, debugMode)
}
