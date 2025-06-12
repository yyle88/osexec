package osexec_test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/runpath"
)

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
		data, err := osexec.NewCommandConfig().WithPath(root).WithDebugShowCmd(true).Exec("ls", "-a")
		require.NoError(t, err)
		t.Log(string(data))
		require.Contains(t, string(data), runpath.Name())
	})

	t.Run("case4", func(t *testing.T) {
		data, err := osexec.NewCommandConfig().WithPath(root).WithDebugShowRes(true).Exec("ls", "-a")
		require.NoError(t, err)
		t.Log(string(data))
		require.Contains(t, string(data), runpath.Name())
	})
}

func TestCommandConfig_ExecInEnvs(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithEnvs([]string{"a=100", "b=200"}).Exec("bash", "-c", "echo $a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Equal(t, "100", strings.TrimSpace(string(data)))
}

func TestCommandConfig_ExecXshRun(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithShellType("bash").WithShellFlag("-c").Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCommandConfig_ExecXshRun_WithBash(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithBash().Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCommandConfig_ExecXshRun_WithBash2(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithBash().Exec("echo $HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCommandConfig_ExecXshRun_WithZsh(t *testing.T) {
	// 检测环境是否支持 zsh
	path, err := exec.LookPath("zsh")
	if err != nil { // 假如测试环境里没有 zsh 就会报错
		t.Skip("zsh is not available on this system, skipping test case")
	}
	t.Log(path)

	data, err := osexec.NewCommandConfig().WithZsh().Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCommandConfig_ExecXshRun_WithSh(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithDebugMode(false).WithSh().Exec("echo $HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCommandConfig_WithTakeExits(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithDebugMode(true).
		WithSh().
		WithTakeExits(map[int]string{1: "DIFFERENCES FOUND"}).
		Exec("diff", "-u", "go.mod", "go.sum")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCommandConfig_WithExpectExit(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithDebugMode(true).
		WithSh().
		WithExpectExit(1, "DIFFERENCES FOUND").
		Exec("diff", "-u", "go.mod", "go.sum")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCommandConfig_WithExpectCode(t *testing.T) {
	data, err := osexec.NewCommandConfig().WithDebugMode(true).
		WithSh().
		WithExpectCode(1). // DIFFERENCES FOUND // 当发现区别时就不算是有错误
		Exec("diff", "-u", "go.mod", "go.sum")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

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
