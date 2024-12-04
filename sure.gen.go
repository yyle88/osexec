package osexec

import "github.com/yyle88/sure"

type CommandConfig88Must struct{ c *CommandConfig }

func (c *CommandConfig) Must() *CommandConfig88Must {
	return &CommandConfig88Must{c: c}
}
func (T *CommandConfig88Must) WithEnvs(envs []string) (res *CommandConfig) {
	res = T.c.WithEnvs(envs)
	return res
}
func (T *CommandConfig88Must) WithPath(path string) (res *CommandConfig) {
	res = T.c.WithPath(path)
	return res
}
func (T *CommandConfig88Must) WithShellType(shellType string) (res *CommandConfig) {
	res = T.c.WithShellType(shellType)
	return res
}
func (T *CommandConfig88Must) WithShellFlag(shellFlag string) (res *CommandConfig) {
	res = T.c.WithShellFlag(shellFlag)
	return res
}
func (T *CommandConfig88Must) WithShell(shellType string, shellFlag string) (res *CommandConfig) {
	res = T.c.WithShell(shellType, shellFlag)
	return res
}
func (T *CommandConfig88Must) WithBash() (res *CommandConfig) {
	res = T.c.WithBash()
	return res
}
func (T *CommandConfig88Must) WithZsh() (res *CommandConfig) {
	res = T.c.WithZsh()
	return res
}
func (T *CommandConfig88Must) WithSh() (res *CommandConfig) {
	res = T.c.WithSh()
	return res
}
func (T *CommandConfig88Must) WithDebugMode(debugMode bool) (res *CommandConfig) {
	res = T.c.WithDebugMode(debugMode)
	return res
}
func (T *CommandConfig88Must) Exec(name string, args ...string) (res []byte) {
	res, err1 := T.c.Exec(name, args...)
	sure.Must(err1)
	return res
}

type CommandConfig88Soft struct{ c *CommandConfig }

func (c *CommandConfig) Soft() *CommandConfig88Soft {
	return &CommandConfig88Soft{c: c}
}
func (T *CommandConfig88Soft) WithEnvs(envs []string) (res *CommandConfig) {
	res = T.c.WithEnvs(envs)
	return res
}
func (T *CommandConfig88Soft) WithPath(path string) (res *CommandConfig) {
	res = T.c.WithPath(path)
	return res
}
func (T *CommandConfig88Soft) WithShellType(shellType string) (res *CommandConfig) {
	res = T.c.WithShellType(shellType)
	return res
}
func (T *CommandConfig88Soft) WithShellFlag(shellFlag string) (res *CommandConfig) {
	res = T.c.WithShellFlag(shellFlag)
	return res
}
func (T *CommandConfig88Soft) WithShell(shellType string, shellFlag string) (res *CommandConfig) {
	res = T.c.WithShell(shellType, shellFlag)
	return res
}
func (T *CommandConfig88Soft) WithBash() (res *CommandConfig) {
	res = T.c.WithBash()
	return res
}
func (T *CommandConfig88Soft) WithZsh() (res *CommandConfig) {
	res = T.c.WithZsh()
	return res
}
func (T *CommandConfig88Soft) WithSh() (res *CommandConfig) {
	res = T.c.WithSh()
	return res
}
func (T *CommandConfig88Soft) WithDebugMode(debugMode bool) (res *CommandConfig) {
	res = T.c.WithDebugMode(debugMode)
	return res
}
func (T *CommandConfig88Soft) Exec(name string, args ...string) (res []byte) {
	res, err1 := T.c.Exec(name, args...)
	sure.Soft(err1)
	return res
}

type CommandConfig88Omit struct{ c *CommandConfig }

func (c *CommandConfig) Omit() *CommandConfig88Omit {
	return &CommandConfig88Omit{c: c}
}
func (T *CommandConfig88Omit) WithEnvs(envs []string) (res *CommandConfig) {
	res = T.c.WithEnvs(envs)
	return res
}
func (T *CommandConfig88Omit) WithPath(path string) (res *CommandConfig) {
	res = T.c.WithPath(path)
	return res
}
func (T *CommandConfig88Omit) WithShellType(shellType string) (res *CommandConfig) {
	res = T.c.WithShellType(shellType)
	return res
}
func (T *CommandConfig88Omit) WithShellFlag(shellFlag string) (res *CommandConfig) {
	res = T.c.WithShellFlag(shellFlag)
	return res
}
func (T *CommandConfig88Omit) WithShell(shellType string, shellFlag string) (res *CommandConfig) {
	res = T.c.WithShell(shellType, shellFlag)
	return res
}
func (T *CommandConfig88Omit) WithBash() (res *CommandConfig) {
	res = T.c.WithBash()
	return res
}
func (T *CommandConfig88Omit) WithZsh() (res *CommandConfig) {
	res = T.c.WithZsh()
	return res
}
func (T *CommandConfig88Omit) WithSh() (res *CommandConfig) {
	res = T.c.WithSh()
	return res
}
func (T *CommandConfig88Omit) WithDebugMode(debugMode bool) (res *CommandConfig) {
	res = T.c.WithDebugMode(debugMode)
	return res
}
func (T *CommandConfig88Omit) Exec(name string, args ...string) (res []byte) {
	res, err1 := T.c.Exec(name, args...)
	sure.Omit(err1)
	return res
}
