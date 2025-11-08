package osexec

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestNewDebugMode tests debug mode creation from boolean values
// TestNewDebugMode 测试从布尔值创建调试模式
func TestNewDebugMode(t *testing.T) {
	require.Equal(t, DEBUG, NewDebugMode(true))
	require.Equal(t, QUIET, NewDebugMode(false))
}

// Test_isShowCommand tests if commands should be shown in different debug modes
// Test_isShowCommand 测试不同调试模式下是否应显示命令
func Test_isShowCommand(t *testing.T) {
	require.True(t, isShowCommand(DEBUG))
	require.True(t, isShowCommand(SHOW_COMMAND))
	require.False(t, isShowCommand(QUIET))
	require.False(t, isShowCommand(SHOW_OUTPUTS))
}

// Test_isShowOutputs tests if outputs should be shown in different debug modes
// Test_isShowOutputs 测试不同调试模式下是否应显示输出
func Test_isShowOutputs(t *testing.T) {
	require.True(t, isShowOutputs(DEBUG))
	require.True(t, isShowOutputs(SHOW_OUTPUTS))
	require.False(t, isShowOutputs(QUIET))
	require.False(t, isShowOutputs(SHOW_COMMAND))
}
