// Package osexectest provides testing utilities for command execution tests
// Contains functions to skip tests when commands are not available in PATH
// Logs command paths when commands are found on the system
//
// osexectest 提供命令执行测试的辅助工具
// 包含在命令不可用时跳过测试的辅助函数
// 在系统中找到命令时会记录命令路径
package osexectest

import (
	"os/exec"
	"testing"
)

// SkipIfCommandNotFound skips the test if the specified command is not found in PATH
// Logs the command path when found
//
// SkipIfCommandNotFound 如果指定的命令在 PATH 中找不到则跳过测试
// 找到命令时会打印命令路径
func SkipIfCommandNotFound(t *testing.T, name string) {
	absPath, err := exec.LookPath(name)
	if err != nil {
		t.Skipf("[%s] is not available on this system, skipping test case", name)
	}
	t.Logf("[%s] is available on this system at: (%s)", name, absPath)
}
