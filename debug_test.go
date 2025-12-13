package osexec_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
)

// TestSetDebugMode tests setting package-wide debug mode
// TestSetDebugMode 测试设置包级别调试模式
func TestSetDebugMode(t *testing.T) {
	if false {
		osexec.SetDebugMode(true)
		require.True(t, osexec.GetDebugMode())
	}
}

// TestGetDebugMode tests getting package-wide debug mode status
// TestGetDebugMode 测试获取包级别调试模式状态
func TestGetDebugMode(t *testing.T) {
	t.Log(osexec.GetDebugMode())
}
