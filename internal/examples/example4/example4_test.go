package example4_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
)

// TestQuietMode demonstrates execution without debug output
// TestQuietMode 演示无调试输出的执行
func TestQuietMode(t *testing.T) {
	// Execute in quiet mode (no debug output)
	// 以静默模式执行（无调试输出）
	config := osexec.NewCommandConfig().
		WithDebugMode(osexec.QUIET)

	output, err := config.Exec("echo", "Quiet mode")
	require.NoError(t, err)
	require.Contains(t, string(output), "Quiet mode")
	t.Logf("Result: %s", string(output))
}

// TestDebugMode demonstrates complete debug output
// TestDebugMode 演示完整的调试输出
func TestDebugMode(t *testing.T) {
	// Execute with complete debug information
	// 执行并显示完整的调试信息
	config := osexec.NewCommandConfig().
		WithDebugMode(osexec.DEBUG)

	output, err := config.Exec("echo", "Debug mode active")
	require.NoError(t, err)
	require.Contains(t, string(output), "Debug mode active")
	t.Logf("Result: %s", string(output))
}

// TestShowCommandMode demonstrates showing command without output
// TestShowCommandMode 演示仅显示命令不显示输出
func TestShowCommandMode(t *testing.T) {
	// Show command being executed but not the output
	// 显示正在执行的命令但不显示输出
	config := osexec.NewCommandConfig().
		WithDebugMode(osexec.SHOW_COMMAND)

	output, err := config.Exec("echo", "Command shown, output hidden")
	require.NoError(t, err)
	require.Contains(t, string(output), "Command shown, output hidden")
	t.Logf("Result: %s", string(output))
}

// TestShowOutputsMode demonstrates showing output without command
// TestShowOutputsMode 演示仅显示输出不显示命令
func TestShowOutputsMode(t *testing.T) {
	// Show command output but not the command itself
	// 显示命令输出但不显示命令本身
	config := osexec.NewCommandConfig().
		WithDebugMode(osexec.SHOW_OUTPUTS)

	output, err := config.Exec("echo", "Output shown, command hidden")
	require.NoError(t, err)
	require.Contains(t, string(output), "Output shown, command hidden")
	t.Logf("Result: %s", string(output))
}
