package utils

import (
	"fmt"

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
	if a.E != nil {
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
