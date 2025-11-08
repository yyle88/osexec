package utils_test

import (
	"testing"

	"github.com/yyle88/osexec/internal/utils"
)

// TestShowCommand tests displaying command string
// TestShowCommand 测试显示命令字符串
func TestShowCommand(t *testing.T) {
	utils.ShowCommand("ls")
}

// TestShowMessage tests displaying message string
// TestShowMessage 测试显示消息字符串
func TestShowMessage(t *testing.T) {
	utils.ShowMessage("ok")
}
