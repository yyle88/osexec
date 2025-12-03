// Package osexectest provides testing utilities for command execution tests
// Contains functions to skip tests when commands are not available in PATH
// Logs command paths when commands are found on the system
//
// osexectest 提供命令执行测试的辅助工具
// 包含在命令不可用时跳过测试的辅助函数
// 在系统中找到命令时会记录命令路径
package osexectest

import (
	"os"
	"os/exec"
	"testing"

	"github.com/yyle88/must"
	"github.com/yyle88/zaplog"
)

// SkipIfCommandNotFound skips the test if the specified command is not found in PATH
// Logs the command path when found
//
// SkipIfCommandNotFound 如果指定的命令在 PATH 中找不到则跳过测试
// 找到命令时会打印命令路径
func SkipIfCommandNotFound(t *testing.T, name string) {
	must.Full(t)

	absPath, err := exec.LookPath(name)
	if err != nil {
		t.Skipf("[%s] is not available on this system, skipping test case", name)
	}
	t.Logf("[%s] is available on this system at: (%s)", name, absPath)
}

// ExitIfCommandNotFound exits TestMain with code 0 if the specified command is not found
// Use this in TestMain to skip tests when a command is unavailable
//
// ExitIfCommandNotFound 如果指定的命令不存在则以退出码 0 退出 TestMain
// 在 TestMain 中使用，当命令不可用时跳过测试
func ExitIfCommandNotFound(m *testing.M, name string) {
	ExitWithCodeIfCommandNotFound(must.Full(m), name, 0)
}

// ExitWithCodeIfCommandNotFound exits TestMain with specified code if the command is not found
// Use code 0 to skip tests gracefully, non-zero to indicate failure
//
// ExitWithCodeIfCommandNotFound 如果命令不存在则以指定退出码退出 TestMain
// 使用 0 表示优雅跳过测试，非零表示失败
func ExitWithCodeIfCommandNotFound(m *testing.M, name string, code int) {
	must.Full(m)

	absPath, err := exec.LookPath(name)
	if err != nil {
		if code == 0 {
			zaplog.SUG.Debugf("[%s] is not available on this system, skipping test case", name)
		} else {
			zaplog.SUG.Errorf("[%s] is not available on this system, skipping test case", name)
		}
		os.Exit(code)
	}
	zaplog.SUG.Infof("[%s] is available on this system at: (%s)", name, absPath)
}

// SkipIfEnvNotSet skips the test if the specified environment variable is not set
// Logs the environment variable value when found
//
// SkipIfEnvNotSet 如果指定的环境变量未设置则跳过测试
// 找到环境变量时会打印其值
func SkipIfEnvNotSet(t *testing.T, name string) {
	must.Full(t)

	value, exists := os.LookupEnv(name)
	if !exists {
		t.Skipf("[%s] environment variable is not set, skipping test case", name)
	}
	t.Logf("[%s] environment variable is set with value: (%s)", name, value)
}

// ExitIfEnvNotSet exits TestMain with code 0 if the specified environment variable is not set
// Use this in TestMain to skip tests when an environment variable is unavailable
//
// ExitIfEnvNotSet 如果指定的环境变量未设置则以退出码 0 退出 TestMain
// 在 TestMain 中使用，当环境变量不可用时跳过测试
func ExitIfEnvNotSet(m *testing.M, name string) {
	ExitWithCodeIfEnvNotSet(must.Full(m), name, 0)
}

// ExitWithCodeIfEnvNotSet exits TestMain with specified code if the environment variable is not set
// Use code 0 to skip tests gracefully, non-zero to indicate failure
//
// ExitWithCodeIfEnvNotSet 如果环境变量未设置则以指定退出码退出 TestMain
// 使用 0 表示优雅跳过测试，非零表示失败
func ExitWithCodeIfEnvNotSet(m *testing.M, name string, code int) {
	must.Full(m)

	value, exists := os.LookupEnv(name)
	if !exists {
		if code == 0 {
			zaplog.SUG.Debugf("[%s] environment variable is not set, skipping test case", name)
		} else {
			zaplog.SUG.Errorf("[%s] environment variable is not set, skipping test case", name)
		}
		os.Exit(code)
	}
	zaplog.SUG.Infof("[%s] environment variable is set with value: (%s)", name, value)
}
