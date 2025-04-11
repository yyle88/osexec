package utils

import (
	"errors"
	"fmt"
	"os/exec"
	"strconv"

	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/eroticgo"
)

func ShowCommand(message string) {
	fmt.Println(eroticgo.BLUE.Sprint("---"))
	fmt.Println(eroticgo.BLUE.Sprint(message))
	fmt.Println(eroticgo.BLUE.Sprint("---"))
}

func ShowMessage(message string) {
	fmt.Println(eroticgo.GREEN.Sprint("---"))
	fmt.Println(eroticgo.GREEN.Sprint(message))
	fmt.Println(eroticgo.GREEN.Sprint("---"))
}

func ShowWarning(message string) {
	fmt.Println(eroticgo.RED.Sprint("---"))
	fmt.Println(eroticgo.RED.Sprint(message))
	fmt.Println(eroticgo.RED.Sprint("---"))
}

// WarpMessage handles the output of the executed command and wraps errors.
// WarpMessage 处理执行命令的输出，并在出现错误时封装错误信息。
func WarpMessage(a *done.Vae[byte], debugMode bool) ([]byte, error) {
	return WarpResults(a, debugMode, map[int]string{})
}

func WarpResults(a *done.Vae[byte], debugMode bool, expectedExitCodes map[int]string) ([]byte, error) {
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
