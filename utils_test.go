package osexec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSetEnableDebug(t *testing.T) {
	SetEnableDebug(true)
	require.True(t, enableDebug)
}

func TestNewCMC(t *testing.T) {
	commandConfig := NewCMC()
	commandConfig.WithShell("bash", "-c")
	commandConfig.WithEnvs([]string{"A=1", "B=2"})
	data, err := commandConfig.Exec("echo", "$A", "$B")
	require.NoError(t, err)
	t.Log(string(data))
}
