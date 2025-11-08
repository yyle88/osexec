package osexec_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/runpath"
)

// TestExec tests basic command execution
// TestExec 测试基本命令执行
func TestExec(t *testing.T) {
	data, err := osexec.Exec("go", "version")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestExecInPath tests command execution in specific path
// TestExecInPath 测试在指定路径中执行命令
func TestExecInPath(t *testing.T) {
	root := runpath.PARENT.Path()
	data, err := osexec.ExecInPath(root, "ls", "-a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Contains(t, string(data), runpath.Name())
}

// TestExecInEnvs tests command execution with custom environment variables
// TestExecInEnvs 测试使用自定义环境变量执行命令
func TestExecInEnvs(t *testing.T) {
	data, err := osexec.ExecInEnvs([]string{"a=100", "b=200"}, "bash", "-c", "echo $a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Equal(t, "100", strings.TrimSpace(string(data)))
}

// TestExecXshRun tests shell command execution with multiple arguments
// TestExecXshRun 测试使用多个参数执行 shell 命令
func TestExecXshRun(t *testing.T) {
	data, err := osexec.ExecXshRun("bash", "-c", "echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestExecXshRun1 tests shell command execution with single command string
// TestExecXshRun1 测试使用单个命令字符串执行 shell 命令
func TestExecXshRun1(t *testing.T) {
	data, err := osexec.ExecXshRun("bash", "-c", "echo $HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}
