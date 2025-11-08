// Package osexec provides simple and flexible command execution utilities
//
// osexec 提供简单而灵活的命令执行工具包
//
// Provides enhanced abstraction on os/exec with customizable configurations
// 提供增强的 os/exec 抽象接口，支持可定制的配置选项
//
// Supports custom environment variables, working paths, and shell options
// 支持自定义环境变量、工作路径和 shell 选项
//
// Includes debug modes and intelligent error handling capabilities
// 包含调试模式和智能错误处理能力
package osexec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/osexec/internal/utils"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Exec executes a command.
// Exec 执行一个命令且获得结果。
func Exec(name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("can-not-execute-with-blank-command-name")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("can-not-contains-space-in-command-name")
	}
	if debugModeOpen {
		debugMessage := strings.TrimSpace(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
		utils.ShowCommand(debugMessage)
		zaplog.ZAPS.Skip1.LOG.Debug("EXEC:", zap.String("CMD", debugMessage))
	}
	command := exec.Command(name, args...)
	return utils.WarpMessage(done.VAE(command.CombinedOutput()), debugModeOpen)
}

// ExecInPath executes a command in a specified DIR.
// ExecInPath 在指定的 DIR 中执行一个命令。
func ExecInPath(path string, name string, args ...string) ([]byte, error) {
	if path == "" {
		return nil, erero.New("can-not-execute-in-blank-DIR-path")
	}
	if name == "" {
		return nil, erero.New("can-not-execute-with-blank-command-name")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("can-not-contains-space-in-command-name")
	}
	if debugModeOpen {
		debugMessage := strings.TrimSpace(fmt.Sprintf("cd %s && %s", path, makeCommandMessage(name, args)))
		utils.ShowCommand(debugMessage)
		zaplog.ZAPS.Skip1.LOG.Debug("EXEC_IN_PATH:", zap.String("CMD", debugMessage))
	}
	command := exec.Command(name, args...)
	command.Dir = path
	return utils.WarpMessage(done.VAE(command.CombinedOutput()), debugModeOpen)
}

// ExecInEnvs executes a command with custom environment variables.
// ExecInEnvs 使用自定义环境变量执行一个命令。
func ExecInEnvs(envs []string, name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("can-not-execute-with-blank-command-name")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("can-not-contains-space-in-command-name")
	}
	if debugModeOpen {
		debugMessage := strings.TrimSpace(fmt.Sprintf("%s %s", strings.Join(envs, " "), makeCommandMessage(name, args)))
		utils.ShowCommand(debugMessage)
		zaplog.ZAPS.Skip1.LOG.Debug("EXEC_IN_ENVS:", zap.String("CMD", debugMessage))
	}
	command := exec.Command(name, args...)
	command.Env = append(os.Environ(), envs...) // Add custom environment variables on top of system envs // 在系统环境变量基础上添加自定义环境变量
	return utils.WarpMessage(done.VAE(command.CombinedOutput()), debugModeOpen)
}

// ExecXshRun executes a command using a specific shell type and shell flag.
// ExecXshRun 使用指定的 shell 类型和 shell 标志执行一个命令。
func ExecXshRun(shellType, shellFlag string, name string, args ...string) ([]byte, error) {
	if shellType == "" {
		return nil, erero.New("can-not-execute-with-wrong-shell-command")
	}
	if shellFlag != "-c" {
		return nil, erero.New("can-not-execute-with-wrong-shell-options")
	}
	if name == "" {
		return nil, erero.New("can-not-execute-with-blank-command-name")
	}
	if debugModeOpen {
		debugMessage := strings.TrimSpace(fmt.Sprintf("%s %s '%s'", shellType, shellFlag, escapeSingleQuotes(makeCommandMessage(name, args))))
		utils.ShowCommand(debugMessage)
		zaplog.ZAPS.Skip1.LOG.Debug("EXEC_XSH_RUN:", zap.String("CMD", debugMessage))
	}
	command := exec.Command(shellType, "-c", name+" "+strings.Join(args, " "))
	return utils.WarpMessage(done.VAE(command.CombinedOutput()), debugModeOpen)
}

// makeCommandMessage formats a command name and its arguments into a single command-line string.
// makeCommandMessage 将命令名称及其参数格式化为一个命令行字符串。
func makeCommandMessage(name string, args []string) string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
}

// escapeSingleQuotes escapes single quotes in a string to ensure safe use in shell commands
// escapeSingleQuotes 转义字符串中的单引号，确保在 shell 命令中安全使用
func escapeSingleQuotes(input string) string {
	return strings.ReplaceAll(input, "'", `'\''`)
}
