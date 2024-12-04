package osexec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Exec executes a command.
// Exec 执行一个命令且获得结果。
func Exec(name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN COMMAND NAME")
	}
	if enableDebug {
		debugMessage := strings.TrimSpace(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
		showMessage(debugMessage)
		zaplog.ZAPS.P1.LOG.Debug("EXEC:", zap.String("CMD", debugMessage))
	}
	command := exec.Command(name, args...)
	return warpOutput(command.CombinedOutput())
}

// ExecInPath executes a command in a specified directory.
// ExecInPath 在指定的目录中执行一个命令。
func ExecInPath(path string, name string, args ...string) ([]byte, error) {
	if path == "" {
		return nil, erero.New("CAN NOT EXEC IN BLANK DIRECTORY PATH")
	}
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN COMMAND NAME")
	}
	if enableDebug {
		debugMessage := strings.TrimSpace(fmt.Sprintf("cd %s && %s", path, formatCommandMessage(name, args)))
		showMessage(debugMessage)
		zaplog.ZAPS.P1.LOG.Debug("EXEC_IN_PATH:", zap.String("CMD", debugMessage))
	}
	command := exec.Command(name, args...)
	command.Dir = path
	return warpOutput(command.CombinedOutput())
}

// ExecInEnvs executes a command with custom environment variables.
// ExecInEnvs 使用自定义环境变量执行一个命令。
func ExecInEnvs(envs []string, name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN COMMAND NAME")
	}
	if enableDebug {
		debugMessage := strings.TrimSpace(fmt.Sprintf("%s %s", strings.Join(envs, " "), formatCommandMessage(name, args)))
		showMessage(debugMessage)
		zaplog.ZAPS.P1.LOG.Debug("EXEC_IN_ENVS:", zap.String("CMD", debugMessage))
	}
	command := exec.Command(name, args...)
	command.Env = os.Environ() // Add custom environment variables
	command.Env = append(command.Env, envs...)
	return warpOutput(command.CombinedOutput())
}

// ExecXshRun executes a command using a specific shell type and shell flag.
// ExecXshRun 使用指定的 shell 类型和 shell 标志执行一个命令。
func ExecXshRun(shellType, shellFlag string, name string, args ...string) ([]byte, error) {
	if shellType == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY SHELL TYPE")
	}
	if shellFlag != "-c" {
		return nil, erero.New("CAN NOT EXECUTE WITH WRONG SHELL OPTIONS")
	}
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN COMMAND NAME")
	}
	if enableDebug {
		debugMessage := strings.TrimSpace(fmt.Sprintf("%s %s '%s'", shellType, shellFlag, escapeSingleQuotes(formatCommandMessage(name, args))))
		showMessage(debugMessage)
		zaplog.ZAPS.P1.LOG.Debug("EXEC_XSH_RUN:", zap.String("CMD", debugMessage))
	}
	command := exec.Command(shellType, "-c", name+" "+strings.Join(args, " "))
	return warpOutput(command.CombinedOutput())
}

// warpOutput handles the output of the executed command and wraps errors if they occur.
// warpOutput 处理执行命令的输出，并在出现错误时封装错误信息。
func warpOutput(output []byte, erx error) ([]byte, error) {
	if erx != nil {
		if enableDebug {
			if len(output) > 0 {
				showWarning(string(output))
			} else {
				showWarning(erx.Error())
			}
		}
		return output, erero.Wro(erx)
	}
	return output, nil
}

// formatCommandMessage formats a command name and its arguments into a single command-line string.
// formatCommandMessage 将命令名称及其参数格式化为一个命令行字符串。
func formatCommandMessage(name string, args []string) string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
}

// escapeSingleQuotes escapes single quotes in a string for safe use in shell commands.
// escapeSingleQuotes 转义字符串中的单引号，以便在 shell 命令中安全使用。
func escapeSingleQuotes(input string) string {
	return strings.ReplaceAll(input, "'", `'\''`)
}
