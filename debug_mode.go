package osexec

import "slices"

type DebugMode string

//goland:noinspection GoSnakeCaseUsage
const (
	QUIET        DebugMode = "QUIET"
	DEBUG        DebugMode = "DEBUG"
	SHOW_COMMAND DebugMode = "SHOW_COMMAND"
	SHOW_OUTPUTS DebugMode = "SHOW_OUTPUTS"
)

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
