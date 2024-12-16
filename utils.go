package osexec

type CMC = CommandConfig

func NewCMC() *CMC {
	return NewCommandConfig()
}

type OsCommand = CommandConfig

func NewOsCommand() *OsCommand {
	return NewCommandConfig()
}

var debugModeOpen = true

func SetDebugMode(enable bool) {
	debugModeOpen = enable
}
