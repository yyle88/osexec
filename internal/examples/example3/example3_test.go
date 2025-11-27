package example3_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexec/osexectest"
)

// TestBashExecution demonstrates bash shell command execution
// TestBashExecution 演示 bash shell 命令执行
func TestBashExecution(t *testing.T) {
	// Execute command using bash shell
	// 使用 bash shell 执行命令
	config := osexec.NewCommandConfig().WithBash()

	output, err := config.Exec("echo", "Hello from Bash!")
	require.NoError(t, err)
	require.Contains(t, string(output), "Hello from Bash!")
	t.Logf("Bash output: %s", string(output))
}

// TestShExecution demonstrates sh shell command execution
// TestShExecution 演示 sh shell 命令执行
func TestShExecution(t *testing.T) {
	// Execute command using sh shell
	// 使用 sh shell 执行命令
	config := osexec.NewCommandConfig().WithSh()

	output, err := config.Exec("echo", "Hello from Sh!")
	require.NoError(t, err)
	require.Contains(t, string(output), "Hello from Sh!")
	t.Logf("Sh output: %s", string(output))
}

// TestZshExecution demonstrates zsh shell command execution
// TestZshExecution 演示 zsh shell 命令执行
func TestZshExecution(t *testing.T) {
	osexectest.SkipIfCommandNotFound(t, "zsh")

	// Execute command using zsh shell
	// 使用 zsh shell 执行命令
	config := osexec.NewCommandConfig().WithZsh()

	output, err := config.Exec("echo", "Hello from Zsh!")
	require.NoError(t, err)
	require.Contains(t, string(output), "Hello from Zsh!")
	t.Logf("Zsh output: %s", string(output))
}

// TestComplexShellCommand demonstrates complex shell command with pipes and variables
// TestComplexShellCommand 演示包含管道和变量的复杂 shell 命令
func TestComplexShellCommand(t *testing.T) {
	// Execute complex shell command
	// 执行复杂的 shell 命令
	config := osexec.NewCommandConfig().
		WithBash().
		WithDebugMode(osexec.SHOW_COMMAND)

	output, err := config.Exec("echo 'apple\nbanana\ncherry' | grep 'banana'")
	require.NoError(t, err)
	require.Contains(t, string(output), "banana")
	t.Logf("Grep result: %s", string(output))
}

// TestShellWithEnvVars demonstrates shell command with environment variables
// TestShellWithEnvVars 演示带环境变量的 shell 命令
func TestShellWithEnvVars(t *testing.T) {
	// Use shell with environment variables
	// 在 shell 中使用环境变量
	config := osexec.NewCommandConfig().
		WithBash().
		WithEnvs([]string{"GREETING=Hello", "NAME=World"})

	output, err := config.Exec("echo", "$GREETING, $NAME!")
	require.NoError(t, err)
	require.Contains(t, string(output), "Hello, World!")
	t.Logf("Output: %s", string(output))
}
