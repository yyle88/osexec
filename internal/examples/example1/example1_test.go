package example1_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
)

// TestBasicExec demonstrates the most basic approach to execute a command
// TestBasicExec 演示执行命令的最基本方式
func TestBasicExec(t *testing.T) {
	// Execute a simple command
	// 执行一个简单的命令
	output, err := osexec.Exec("go", "version")
	require.NoError(t, err)
	require.NotEmpty(t, output)
	t.Logf("Go version: %s", string(output))
}

// TestExecInPath demonstrates executing a command in a specific path
// TestExecInPath 演示在指定路径中执行命令
func TestExecInPath(t *testing.T) {
	// Create temp DIR using t.TempDir()
	// 使用 t.TempDir() 创建临时 DIR
	tempDIR := t.TempDir()

	// Execute ls command in temp DIR
	// 在临时 DIR 执行 ls 命令
	output, err := osexec.ExecInPath(tempDIR, "ls", "-la")
	require.NoError(t, err)
	require.NotEmpty(t, output)
	t.Logf("Temp DIR: %s", tempDIR)
	t.Logf("Output:\n%s", string(output))
}

// TestExecWithEnv demonstrates executing a command with custom environment variables
// TestExecWithEnv 演示使用自定义环境变量执行命令
func TestExecWithEnv(t *testing.T) {
	// Set custom environment variables
	// 设置自定义环境变量
	envs := []string{
		"MY_VAR=hello",
		"ANOTHER_VAR=world",
	}

	// Execute command with custom environment
	// 使用自定义环境执行命令
	output, err := osexec.ExecInEnvs(envs, "bash", "-c", "echo $MY_VAR $ANOTHER_VAR")
	require.NoError(t, err)
	require.Contains(t, string(output), "hello world")
	t.Logf("Output: %s", string(output))
}
