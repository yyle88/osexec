package utils

import (
	"fmt"

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
func WarpMessage(output []byte, err error, debugMode bool) ([]byte, error) {
	if err != nil {
		if debugMode {
			if len(output) > 0 {
				ShowWarning(string(output))
			} else {
				ShowWarning(err.Error())
			}
		}
		return output, erero.Wro(err)
	}
	if debugMode {
		if len(output) > 0 {
			ShowMessage(string(output))
		}
	}
	return output, nil
}
