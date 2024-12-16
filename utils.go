package osexec

import (
	"fmt"

	"github.com/yyle88/eroticgo"
)

type CMC = CommandConfig

func NewCMC() *CMC {
	return NewCommandConfig()
}

var debugModeOpen = true

func SetDebugMode(enable bool) {
	debugModeOpen = enable
}

func showMessage(message string) {
	fmt.Println(eroticgo.BLUE.Sprint("---"))
	fmt.Println(eroticgo.BLUE.Sprint(message))
	fmt.Println(eroticgo.BLUE.Sprint("---"))
}

func showWarning(message string) {
	fmt.Println(eroticgo.RED.Sprint("---"))
	fmt.Println(eroticgo.RED.Sprint(message))
	fmt.Println(eroticgo.RED.Sprint("---"))
}
