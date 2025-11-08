package osexec

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
)

// TestSetDebugMode tests setting package-wide debug mode
// TestSetDebugMode 测试设置包级别调试模式
func TestSetDebugMode(t *testing.T) {
	SetDebugMode(true)
	require.True(t, debugModeOpen)
}

// TestNewCMC tests creating command config with shell and environment setup
// TestNewCMC 测试创建带 shell 和环境设置的命令配置
func TestNewCMC(t *testing.T) {
	cmc := NewCMC()
	cmc.WithShell("bash", "-c")
	cmc.WithEnvs([]string{"A=1", "B=2"})
	data, err := cmc.Exec("echo", "$A", "$B")
	require.NoError(t, err)
	t.Log(string(data))
}

// TestNewOsCommand tests creating os command with path, shell and environment setup
// TestNewOsCommand 测试创建带路径、shell 和环境设置的 os 命令
func TestNewOsCommand(t *testing.T) {
	root := runpath.PARENT.Path()
	t.Log(root)

	osCommand := NewOsCommand()
	osCommand.WithPath(root)
	osCommand.WithShell("bash", "-c")
	osCommand.WithEnvs([]string{"A=1", "B=2"})
	data, err := osCommand.StreamExec("echo", "$A", "$B")
	require.NoError(t, err)
	t.Log(string(data))
}
