package osexec_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/runpath"
)

// TestNewExecConfig tests creating exec config with shell and environment setup
// TestNewExecConfig 测试创建带 shell 和环境设置的执行配置
func TestNewExecConfig(t *testing.T) {
	execConfig := osexec.NewExecConfig()
	execConfig.WithShell("bash", "-c")
	execConfig.WithEnvs([]string{"A=1", "B=2"})
	data, err := execConfig.Exec("echo", "$A", "$B")
	require.NoError(t, err)
	t.Log(string(data))
}

// TestNewExecConfig_StreamExec tests creating exec config with path, shell and environment setup
// TestNewExecConfig_StreamExec 测试创建带路径、shell 和环境设置的执行配置
func TestNewExecConfig_StreamExec(t *testing.T) {
	root := runpath.PARENT.Path()
	t.Log(root)

	execConfig := osexec.NewExecConfig()
	execConfig.WithPath(root)
	execConfig.WithShell("bash", "-c")
	execConfig.WithEnvs([]string{"A=1", "B=2"})
	data, err := execConfig.StreamExec("echo", "$A", "$B")
	require.NoError(t, err)
	t.Log(string(data))
}
