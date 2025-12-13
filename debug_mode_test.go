package osexec_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
)

// TestNewDebugMode tests debug mode creation from boolean values
// TestNewDebugMode 测试从布尔值创建调试模式
func TestNewDebugMode(t *testing.T) {
	require.Equal(t, osexec.DEBUG, osexec.NewDebugMode(true))
	require.Equal(t, osexec.QUIET, osexec.NewDebugMode(false))
}
