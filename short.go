package osexec

type CMC = CommandConfig

func NewCMC() *CMC {
	return NewCommandConfig()
}

var enableDebug = true

func SetEnableDebug(enable bool) {
	enableDebug = enable
}
