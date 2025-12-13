// Package utils provides utilities supporting command execution and output formatting
//
// utils 提供命令执行和输出格式化工具
//
// Provides support functions to show command output with colors
// 提供带颜色的命令输出显示支持函数
//
// Handles command results wrapping and exit code processing
// 处理命令结果包装和退出码处理
package utils

import (
	"fmt"
	"os/exec"
	"strconv"

	"github.com/pkg/errors"
	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/eroticgo"
)

// ShowCommand displays command message in blue with dividing lines
// ShowCommand 用蓝色显示命令消息，带分隔线
func ShowCommand(message string) {
	fmt.Println(eroticgo.BLUE.Sprint("---"))
	fmt.Println(eroticgo.BLUE.Sprint(message))
	fmt.Println(eroticgo.BLUE.Sprint("---"))
}

// ShowMessage displays success message in green with dividing lines
// ShowMessage 用绿色显示成功消息，带分隔线
func ShowMessage(message string) {
	fmt.Println(eroticgo.GREEN.Sprint("---"))
	fmt.Println(eroticgo.GREEN.Sprint(message))
	fmt.Println(eroticgo.GREEN.Sprint("---"))
}

// ShowWarning displays warning message in red with dividing lines
// ShowWarning 用红色显示警告消息，带分隔线
func ShowWarning(message string) {
	fmt.Println(eroticgo.RED.Sprint("---"))
	fmt.Println(eroticgo.RED.Sprint(message))
	fmt.Println(eroticgo.RED.Sprint("---"))
}

// WrapOutputs handles command output that succeeds without errors
// WrapOutputs 处理成功的命令输出，无错误
func WrapOutputs(outputs []byte, debugMode bool) ([]byte, error) {
	return WrapMessage(done.VAE(outputs, nil), debugMode)
}

// WrapMessage handles the output of the executed command and wraps errors
// WrapMessage 处理执行命令的输出，并在出现错误时封装错误信息
func WrapMessage(a *done.Vae[byte], debugMode bool) ([]byte, error) {
	return WrapResults(a, debugMode, map[int]string{})
}

// WrapResults handles command output with expected exit codes support
// WrapResults 处理命令输出，支持预期的退出码
func WrapResults(a *done.Vae[byte], debugMode bool, expectedExitCodes map[int]string) ([]byte, error) {
	if a.E != nil {
		if len(expectedExitCodes) > 0 {
			if ext := new(exec.ExitError); errors.As(a.E, &ext) {
				if reason, ok := expectedExitCodes[ext.ExitCode()]; ok {
					if debugMode {
						if reason != "" {
							ShowMessage("EXIT-CODE:" + strconv.Itoa(ext.ExitCode()) + "-" + "REASON:" + reason)
						} else {
							ShowMessage("EXIT-CODE:" + strconv.Itoa(ext.ExitCode()))
						}
						if len(a.V) > 0 {
							ShowMessage(string(a.V))
						}
					}
					return a.V, nil
				}
			}
		}

		if debugMode {
			if len(a.V) > 0 {
				ShowWarning(string(a.V))
			} else {
				ShowWarning(a.E.Error())
			}
		}
		return a.V, erero.Wro(a.E)
	}
	if debugMode {
		if len(a.V) > 0 {
			ShowMessage(string(a.V))
		}
	}
	return a.V, nil
}

// WrapOutcome handles command output and returns exit code
// Combines output wrapping with exit code extraction
//
// WrapOutcome 处理命令输出并返回退出码
// 结合输出包装和退出码提取
func WrapOutcome(a *done.Vae[byte], debugMode bool, expectedExitCodes map[int]string) ([]byte, int, error) {
	output, err := WrapResults(a, debugMode, expectedExitCodes)
	if err != nil {
		exitCode := ExceptsCode(a.E)
		// Use errors.WithMessagef instead of erero to wrap without logging
		// 使用 errors.WithMessagef 而非 erero 包装错误，避免重复打印日志
		return output, exitCode, errors.WithMessagef(err, "command exit code: %d", exitCode)
	}
	if a.E != nil {
		return output, ExceptsCode(a.E), nil
	}
	return output, 0, nil
}

// ExceptsCode extracts exit code from command execution issue
// Returns 0 with no issue, the exit code if ExitError, -1 in remaining cases
//
// ExceptsCode 从命令执行问题中提取退出码
// 无问题返回 0，ExitError 返回退出码，其余情况返回 -1
func ExceptsCode(err error) int {
	if err == nil {
		return 0
	}
	var ext *exec.ExitError
	if errors.As(err, &ext) {
		return ext.ExitCode()
	}
	return -1
}
