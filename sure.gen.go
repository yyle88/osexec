// Code generated using sure/sure_cls_gen. DO NOT EDIT.
// This file was auto generated via github.com/yyle88/sure
// Generated from: sure.gen_test.go:27 -> osexec.TestGen
// ========== SURE:DO-NOT-EDIT-SECTION:END ==========

package osexec

import (
	"os/exec"

	"github.com/yyle88/sure"
)

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
func (T *CommandConfig88Must) WithDebug() (res *CommandConfig) {
	res = T.c.WithDebug()
	return res
}
func (T *CommandConfig88Must) WithDebugMode(debugMode DebugMode) (res *CommandConfig) {
	res = T.c.WithDebugMode(debugMode)
	return res
}
func (T *CommandConfig88Must) WithMatchPipe(matchPipe func(outputLine string) bool) (res *CommandConfig) {
	res = T.c.WithMatchPipe(matchPipe)
	return res
}
func (T *CommandConfig88Must) WithMatchMore(matchMore bool) (res *CommandConfig) {
	res = T.c.WithMatchMore(matchMore)
	return res
}
func (T *CommandConfig88Must) WithTakeExits(takeExits map[int]string) (res *CommandConfig) {
	res = T.c.WithTakeExits(takeExits)
	return res
}
func (T *CommandConfig88Must) WithExpectExit(exitCode int, reason string) (res *CommandConfig) {
	res = T.c.WithExpectExit(exitCode, reason)
	return res
}
func (T *CommandConfig88Must) WithExpectCode(exitCode int) (res *CommandConfig) {
	res = T.c.WithExpectCode(exitCode)
	return res
}
func (T *CommandConfig88Must) Exec(name string, args ...string) (res []byte) {
	res, err1 := T.c.Exec(name, args...)
	sure.Must(err1)
	return res
}
func (T *CommandConfig88Must) ExecWith(name string, args []string, prepare func(command *exec.Cmd)) (res []byte) {
	res, err1 := T.c.ExecWith(name, args, prepare)
	sure.Must(err1)
	return res
}
func (T *CommandConfig88Must) ExecTake(name string, args ...string) (res []byte, res1 int) {
	res, res1, err2 := T.c.ExecTake(name, args...)
	sure.Must(err2)
	return res, res1
}
func (T *CommandConfig88Must) IsShowCommand() (res bool) {
	res = T.c.IsShowCommand()
	return res
}
func (T *CommandConfig88Must) IsShowOutputs() (res bool) {
	res = T.c.IsShowOutputs()
	return res
}
func (T *CommandConfig88Must) StreamExec(name string, args ...string) (res []byte) {
	res, err1 := T.c.StreamExec(name, args...)
	sure.Must(err1)
	return res
}
func (T *CommandConfig88Must) ExecInPipe(name string, args ...string) (res []byte) {
	res, err1 := T.c.ExecInPipe(name, args...)
	sure.Must(err1)
	return res
}
func (T *CommandConfig88Must) NewConfig() (res *CommandConfig) {
	res = T.c.NewConfig()
	return res
}
func (T *CommandConfig88Must) SubConfig(path string) (res *CommandConfig) {
	res = T.c.SubConfig(path)
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
func (T *CommandConfig88Soft) WithDebug() (res *CommandConfig) {
	res = T.c.WithDebug()
	return res
}
func (T *CommandConfig88Soft) WithDebugMode(debugMode DebugMode) (res *CommandConfig) {
	res = T.c.WithDebugMode(debugMode)
	return res
}
func (T *CommandConfig88Soft) WithMatchPipe(matchPipe func(outputLine string) bool) (res *CommandConfig) {
	res = T.c.WithMatchPipe(matchPipe)
	return res
}
func (T *CommandConfig88Soft) WithMatchMore(matchMore bool) (res *CommandConfig) {
	res = T.c.WithMatchMore(matchMore)
	return res
}
func (T *CommandConfig88Soft) WithTakeExits(takeExits map[int]string) (res *CommandConfig) {
	res = T.c.WithTakeExits(takeExits)
	return res
}
func (T *CommandConfig88Soft) WithExpectExit(exitCode int, reason string) (res *CommandConfig) {
	res = T.c.WithExpectExit(exitCode, reason)
	return res
}
func (T *CommandConfig88Soft) WithExpectCode(exitCode int) (res *CommandConfig) {
	res = T.c.WithExpectCode(exitCode)
	return res
}
func (T *CommandConfig88Soft) Exec(name string, args ...string) (res []byte) {
	res, err1 := T.c.Exec(name, args...)
	sure.Soft(err1)
	return res
}
func (T *CommandConfig88Soft) ExecWith(name string, args []string, prepare func(command *exec.Cmd)) (res []byte) {
	res, err1 := T.c.ExecWith(name, args, prepare)
	sure.Soft(err1)
	return res
}
func (T *CommandConfig88Soft) ExecTake(name string, args ...string) (res []byte, res1 int) {
	res, res1, err2 := T.c.ExecTake(name, args...)
	sure.Soft(err2)
	return res, res1
}
func (T *CommandConfig88Soft) IsShowCommand() (res bool) {
	res = T.c.IsShowCommand()
	return res
}
func (T *CommandConfig88Soft) IsShowOutputs() (res bool) {
	res = T.c.IsShowOutputs()
	return res
}
func (T *CommandConfig88Soft) StreamExec(name string, args ...string) (res []byte) {
	res, err1 := T.c.StreamExec(name, args...)
	sure.Soft(err1)
	return res
}
func (T *CommandConfig88Soft) ExecInPipe(name string, args ...string) (res []byte) {
	res, err1 := T.c.ExecInPipe(name, args...)
	sure.Soft(err1)
	return res
}
func (T *CommandConfig88Soft) NewConfig() (res *CommandConfig) {
	res = T.c.NewConfig()
	return res
}
func (T *CommandConfig88Soft) SubConfig(path string) (res *CommandConfig) {
	res = T.c.SubConfig(path)
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
func (T *CommandConfig88Omit) WithDebug() (res *CommandConfig) {
	res = T.c.WithDebug()
	return res
}
func (T *CommandConfig88Omit) WithDebugMode(debugMode DebugMode) (res *CommandConfig) {
	res = T.c.WithDebugMode(debugMode)
	return res
}
func (T *CommandConfig88Omit) WithMatchPipe(matchPipe func(outputLine string) bool) (res *CommandConfig) {
	res = T.c.WithMatchPipe(matchPipe)
	return res
}
func (T *CommandConfig88Omit) WithMatchMore(matchMore bool) (res *CommandConfig) {
	res = T.c.WithMatchMore(matchMore)
	return res
}
func (T *CommandConfig88Omit) WithTakeExits(takeExits map[int]string) (res *CommandConfig) {
	res = T.c.WithTakeExits(takeExits)
	return res
}
func (T *CommandConfig88Omit) WithExpectExit(exitCode int, reason string) (res *CommandConfig) {
	res = T.c.WithExpectExit(exitCode, reason)
	return res
}
func (T *CommandConfig88Omit) WithExpectCode(exitCode int) (res *CommandConfig) {
	res = T.c.WithExpectCode(exitCode)
	return res
}
func (T *CommandConfig88Omit) Exec(name string, args ...string) (res []byte) {
	res, err1 := T.c.Exec(name, args...)
	sure.Omit(err1)
	return res
}
func (T *CommandConfig88Omit) ExecWith(name string, args []string, prepare func(command *exec.Cmd)) (res []byte) {
	res, err1 := T.c.ExecWith(name, args, prepare)
	sure.Omit(err1)
	return res
}
func (T *CommandConfig88Omit) ExecTake(name string, args ...string) (res []byte, res1 int) {
	res, res1, err2 := T.c.ExecTake(name, args...)
	sure.Omit(err2)
	return res, res1
}
func (T *CommandConfig88Omit) IsShowCommand() (res bool) {
	res = T.c.IsShowCommand()
	return res
}
func (T *CommandConfig88Omit) IsShowOutputs() (res bool) {
	res = T.c.IsShowOutputs()
	return res
}
func (T *CommandConfig88Omit) StreamExec(name string, args ...string) (res []byte) {
	res, err1 := T.c.StreamExec(name, args...)
	sure.Omit(err1)
	return res
}
func (T *CommandConfig88Omit) ExecInPipe(name string, args ...string) (res []byte) {
	res, err1 := T.c.ExecInPipe(name, args...)
	sure.Omit(err1)
	return res
}
func (T *CommandConfig88Omit) NewConfig() (res *CommandConfig) {
	res = T.c.NewConfig()
	return res
}
func (T *CommandConfig88Omit) SubConfig(path string) (res *CommandConfig) {
	res = T.c.SubConfig(path)
	return res
}
