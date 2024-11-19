package osexec_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/runpath"
)

func TestCMC_ExecInPath(t *testing.T) {
	root := runpath.PARENT.Path()
	data, err := osexec.NewCMC().WithPath(root).Exec("ls", "-a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Contains(t, string(data), runpath.Name())
}

func TestCMC_ExecInEnvs(t *testing.T) {
	data, err := osexec.NewCMC().WithEnvs([]string{"a=100", "b=200"}).Exec("bash", "-c", "echo $a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Equal(t, "100", strings.TrimSpace(string(data)))
}

func TestCMC_ExecXshRun(t *testing.T) {
	data, err := osexec.NewCMC().WithShellType("bash").WithShellFlag("-c").Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCMC_ExecXshRun_WithBash(t *testing.T) {
	data, err := osexec.NewCMC().WithBash().Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCMC_ExecXshRun_WithZsh(t *testing.T) {
	data, err := osexec.NewCMC().WithZsh().Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestCMC_ExecXshRun_WithSh(t *testing.T) {
	data, err := osexec.NewCMC().WithSh().Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}
