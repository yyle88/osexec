package osexec_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
	"github.com/yyle88/runpath"
)

func TestExec(t *testing.T) {
	data, err := osexec.Exec("go", "version")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestExecInPath(t *testing.T) {
	root := runpath.PARENT.Path()
	data, err := osexec.ExecInPath(root, "ls", "-a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Contains(t, string(data), runpath.Name())
}

func TestExecInEnvs(t *testing.T) {
	data, err := osexec.ExecInEnvs([]string{"a=100", "b=200"}, "bash", "-c", "echo $a")
	require.NoError(t, err)
	t.Log(string(data))
	require.Equal(t, "100", strings.TrimSpace(string(data)))
}

func TestExecXshRun(t *testing.T) {
	data, err := osexec.ExecXshRun("bash", "-c", "echo", "$HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}

func TestExecXshRun1(t *testing.T) {
	data, err := osexec.ExecXshRun("bash", "-c", "echo $HOME")
	require.NoError(t, err)
	t.Log(string(data))
	require.NotEmpty(t, strings.TrimSpace(string(data)))
}
