package osexec_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/osexec/osexectest"
	"github.com/yyle88/runpath"
)

// TestCommandConfig_ExecInPath tests command execution in specific path
// TestCommandConfig_ExecInPath 测试在指定路径中执行命令
func TestCommandConfig_ExecInPath(t *testing.T) {
	root := runpath.PARENT.Path()

	t.Run("case1", func(t *testing.T) {
		data, err := osexec.NewCommandConfig().WithPath(root).Exec("ls", "-a")
		require.NoError(t, err)
		t.Log(string(data))
		require.Contains(t, string(data), runpath.Name())
	})

	t.Run("case2", func(t *testing.T) {
		data, err := osexec.NewCommandConfig().WithPath(root).WithDebug().Exec("ls", "-a")
		require.NoError(t, err)
		t.Log(string(data))
		require.Contains(t, string(data), runpath.Name())
	})

	t.Run("case3", func(t *testing.T) {
		data, err := osexec.NewCommandConfig().WithPath(root).WithDebugMode(osexec.SHOW_COMMAND).Exec("ls", "-a")
		require.NoError(t, err)
		t.Log(string(data))
		require.Contains(t, string(data), runpath.Name())
	})

	t.Run("case4", func(t *testing.T) {
		data, err := osexec.NewCommandConfig().WithPath(root).WithDebugMode(osexec.SHOW_OUTPUTS).Exec("ls", "-a")
		require.NoError(t, err)
		t.Log(string(data))
		require.Contains(t, string(data), runpath.Name())
	})
}

// TestCommandConfig_ExecInEnvs tests command execution with environment variables
// TestCommandConfig_ExecInEnvs 测试使用环境变量执行命令
func TestCommandConfig_ExecInEnvs(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithEnvs([]string{"a=100", "b=200"}).Exec("bash", "-c", "echo $a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Equal(t, "100", strings.TrimSpace(string(data)))
}

// TestCommandConfig_ExecXshRun tests shell command execution with custom shell type and flag
// TestCommandConfig_ExecXshRun 测试使用自定义 shell 类型和标志执行 shell 命令
func TestCommandConfig_ExecXshRun(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithShellType("bash").WithShellFlag("-c").Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestCommandConfig_ExecXshRun_WithBash tests bash shell execution with multiple arguments
// TestCommandConfig_ExecXshRun_WithBash 测试使用多个参数执行 bash shell 命令
func TestCommandConfig_ExecXshRun_WithBash(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithBash().Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestCommandConfig_ExecXshRun_WithBash2 tests bash shell execution with single command string
// TestCommandConfig_ExecXshRun_WithBash2 测试使用单个命令字符串执行 bash shell 命令
func TestCommandConfig_ExecXshRun_WithBash2(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithBash().Exec("echo $HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestCommandConfig_ExecXshRun_WithZsh tests zsh shell execution with environment check
// TestCommandConfig_ExecXshRun_WithZsh 测试带环境检查的 zsh shell 命令执行
func TestCommandConfig_ExecXshRun_WithZsh(t *testing.T) {
	osexectest.SkipIfCommandNotFound(t, "zsh")

	data, err := osexec.NewCommandConfig().WithZsh().Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestCommandConfig_ExecXshRun_WithSh tests sh shell execution in quiet mode
// TestCommandConfig_ExecXshRun_WithSh 测试在静默模式下执行 sh shell 命令
func TestCommandConfig_ExecXshRun_WithSh(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithDebugMode(osexec.QUIET).WithSh().Exec("echo $HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestCommandConfig_WithTakeExits tests accepting specific exit codes as success
// TestCommandConfig_WithTakeExits 测试将特定退出码视为成功
func TestCommandConfig_WithTakeExits(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithDebugMode(osexec.DEBUG).
		WithSh().
		WithTakeExits(map[int]string{1: "DIFFERENCES FOUND"}).
		Exec("diff", "-u", "go.mod", "go.sum")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestCommandConfig_WithExpectExit tests expecting a specific exit code with reason
// TestCommandConfig_WithExpectExit 测试期待特定退出码并附带原因
func TestCommandConfig_WithExpectExit(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithDebugMode(osexec.DEBUG).
		WithSh().
		WithExpectExit(1, "DIFFERENCES FOUND").
		Exec("diff", "-u", "go.mod", "go.sum")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestCommandConfig_WithExpectCode tests expecting a specific exit code without reason
// TestCommandConfig_WithExpectCode 测试期待特定退出码但不附带原因
func TestCommandConfig_WithExpectCode(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithDebugMode(osexec.DEBUG).
		WithSh().
		WithExpectCode(1). // DIFFERENCES FOUND // 当发现区别时就不算是有错误
		Exec("diff", "-u", "go.mod", "go.sum")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

// TestCommandConfig_ExecWith tests command execution with custom stdin setup
// TestCommandConfig_ExecWith 测试使用自定义 stdin 设置执行命令
func TestCommandConfig_ExecWith(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithDebug().
		ExecWith("grep", []string{"abc"},
			func(command *exec.Cmd) {
				command.Stdin = strings.NewReader("123abc\nabc456\n123xyz\nxyz456")
			},
		)
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
	require.Equal(t, "123abc\nabc456", strings.TrimSpace(string(data)))
}

// TestCommandConfig_ExecTake tests command execution with exit code capture
// TestCommandConfig_ExecTake 测试带退出码捕获的命令执行
func TestCommandConfig_ExecTake(t *testing.T) {
	tempDIR := t.TempDir()

	fileA := filepath.Join(tempDIR, "a.txt")
	fileB := filepath.Join(tempDIR, "b.txt")

	// Create two files with different content
	// 创建两个内容不同的文件
	require.NoError(t, os.WriteFile(fileA, []byte("hello\n"), 0644))
	require.NoError(t, os.WriteFile(fileB, []byte("world\n"), 0644))

	// diff returns exit code 1 when files have differences
	// diff 在文件有差异时返回退出码 1
	output, exitCode, err := osexec.NewCommandConfig().WithExpectCode(1).ExecTake("diff", fileA, fileB)
	t.Log(string(output))
	require.NoError(t, err)
	require.Equal(t, 1, exitCode)

	// Make files the same
	// 使文件内容一致
	require.NoError(t, os.WriteFile(fileB, []byte("hello\n"), 0644))

	// diff returns exit code 0 when files match
	// diff 在文件一致时返回退出码 0
	output, exitCode, err = osexec.NewCommandConfig().WithExpectCode(1).ExecTake("diff", fileA, fileB)
	t.Log(string(output))
	require.NoError(t, err)
	require.Equal(t, 0, exitCode)
}
