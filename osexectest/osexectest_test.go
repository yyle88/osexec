package osexectest_test

import (
	"testing"

	"github.com/yyle88/osexec/osexectest"
)

// TestMain checks bash command exists before running tests
// TestMain 在运行测试前检查 bash 命令是否存在
func TestMain(m *testing.M) {
	osexectest.ExitIfCommandNotFound(m, "bash")
	m.Run()
}

// TestSkipIfCommandNotFound tests SkipIfCommandNotFound function
// TestSkipIfCommandNotFound 测试 SkipIfCommandNotFound 函数
func TestSkipIfCommandNotFound(t *testing.T) {
	osexectest.SkipIfCommandNotFound(t, "zsh")
}
