package osexec

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/runpath"
)

func TestSetDebugMode(t *testing.T) {
	SetDebugMode(true)
	require.True(t, debugModeOpen)
}

func TestNewCMC(t *testing.T) {
	cmc := NewCMC()
	cmc.WithShell("bash", "-c")
	cmc.WithEnvs([]string{"A=1", "B=2"})
	data, err := cmc.Exec("echo", "$A", "$B")
	require.NoError(t, err)
	t.Log(string(data))
}

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
