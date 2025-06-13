package osexec

import (
	"slices"

	"github.com/yyle88/tern"
)

type DebugMode string

//goland:noinspection GoSnakeCaseUsage
const (
	QUIET        DebugMode = "QUIET"        // 非调试模式，无任何调试输出
	DEBUG        DebugMode = "DEBUG"        // 完整调试模式，包含所有调试信息
	SHOW_COMMAND DebugMode = "SHOW_COMMAND" // 仅显示命令
	SHOW_OUTPUTS DebugMode = "SHOW_OUTPUTS" // 仅显示输出
)

// NewDebugMode creates a DebugMode based on whether debug mode is enabled. If debugModeOpen is true, returns DEBUG; otherwise, returns QUIET.
// NewDebugMode 根据是否启用调试模式创建 DebugMode。 如果 debugModeOpen 为 true，返回 DEBUG；否则返回 QUIET。
func NewDebugMode(debugModeOpen bool) DebugMode {
	return tern.BVV(debugModeOpen, DEBUG, QUIET)
}

// checks whether the command should be displayed based on the debug mode or DebugShowCmd flag.
// 检查是否应根据调试模式或 DebugShowCmd 标志显示命令。
func isShowCommand(debugMode DebugMode) bool {
	return slices.Contains([]DebugMode{DEBUG, SHOW_COMMAND}, debugMode)
}

// checks whether the command results should be displayed based on the debug mode or DebugShowRes flag.
// 检查是否应根据调试模式或 DebugShowRes 标志显示命令结果。
func isShowOutputs(debugMode DebugMode) bool {
	return slices.Contains([]DebugMode{DEBUG, SHOW_OUTPUTS}, debugMode)
}
