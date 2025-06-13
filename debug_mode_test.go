package osexec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isShowCommand(t *testing.T) {
	require.True(t, isShowCommand(DEBUG))
	require.True(t, isShowCommand(SHOW_COMMAND))
	require.False(t, isShowCommand(QUIET))
	require.False(t, isShowCommand(SHOW_OUTPUTS))
}

func Test_isShowOutputs(t *testing.T) {
	require.True(t, isShowOutputs(DEBUG))
	require.True(t, isShowOutputs(SHOW_OUTPUTS))
	require.False(t, isShowOutputs(QUIET))
	require.False(t, isShowOutputs(SHOW_COMMAND))
}
