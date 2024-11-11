package osexec_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/runpath"
)

func TestCMX_ExecInPath(t *testing.T) {
	root := runpath.PARENT.Path()
	data, err := osexec.NewCMX().WithPath(root).Exec("ls", "-a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Contains(t, string(data), runpath.Name())
}

func TestCMX_ExecInEnvs(t *testing.T) {
	data, err := osexec.NewCMX().WithEnvs([]string{"a=100", "b=200"}).Exec("bash", "-c", "echo $a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Equal(t, "100", strings.TrimSpace(string(data)))
}

func TestCMX_ExecXshRun(t *testing.T) {
	data, err := osexec.NewCMX().WithShellType("bash").WithShellFlag("-c").Exec("echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}
