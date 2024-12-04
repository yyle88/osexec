package osexec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/tern"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type CommandConfig struct {
	Envs      []string // 可选的环境变量
	Path      string   // 可选的执行路径
	ShellType string
	ShellFlag string
	DebugMode bool
}

func NewCommandConfig() *CommandConfig {
	return &CommandConfig{
		DebugMode: enableDebug, //初始值和这个环境变量里的相同
	}
}

// WithEnvs 设置环境变量并返回 CommandConfig，以支持链式调用。
func (c *CommandConfig) WithEnvs(envs []string) *CommandConfig {
	c.Envs = envs
	return c
}

// WithPath 设置执行路径并返回 CommandConfig，以支持链式调用。
func (c *CommandConfig) WithPath(path string) *CommandConfig {
	c.Path = path
	return c
}

// WithShellType 设置 shell 类型并返回 CommandConfig，以支持链式调用。
func (c *CommandConfig) WithShellType(shellType string) *CommandConfig {
	c.ShellType = shellType
	return c
}

// WithShellFlag 设置 shell 标志并返回 CommandConfig，以支持链式调用。
func (c *CommandConfig) WithShellFlag(shellFlag string) *CommandConfig {
	c.ShellFlag = shellFlag
	return c
}

func (c *CommandConfig) WithShell(shellType, shellFlag string) *CommandConfig {
	c.ShellType = shellType
	c.ShellFlag = shellFlag
	return c
}

func (c *CommandConfig) WithBash() *CommandConfig {
	return c.WithShell("bash", "-c")
}

func (c *CommandConfig) WithZsh() *CommandConfig {
	return c.WithShell("zsh", "-c")
}

func (c *CommandConfig) WithSh() *CommandConfig {
	return c.WithShell("sh", "-c")
}

func (c *CommandConfig) WithDebugMode(debugMode bool) *CommandConfig {
	c.DebugMode = debugMode
	return c
}

func (c *CommandConfig) Exec(name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN COMMAND NAME")
	}
	if c.ShellType != "" && c.ShellFlag != "-c" { //这里暂不支持别的
		return nil, erero.New("CAN NOT EXECUTE WITH WRONG SHELL OPTIONS")
	}
	if c.DebugMode {
		debugMessage := c.formatCommandLine(name, args)
		zaplog.ZAPS.P1.LOG.Debug("EXEC:", zap.String("CMD", debugMessage))
	}
	cmd := tern.BFF(c.ShellType != "",
		func() *exec.Cmd {
			return exec.Command(c.ShellType, c.ShellFlag, name+" "+strings.Join(args, " "))
		},
		func() *exec.Cmd {
			return exec.Command(name, args...)
		})
	cmd.Dir = c.Path //这样目的在于提示 path 很有可能是空的
	cmd.Env = tern.BF(len(c.Envs) > 0, func() []string {
		return append(os.Environ(), c.Envs...)
	})
	return cmd.CombinedOutput()
}

func (c *CommandConfig) formatCommandLine(name string, args []string) string {
	var stb strings.Builder
	if c.Path != "" {
		stb.WriteString(fmt.Sprintf("cd %s && ", c.Path))
	}
	if len(c.Envs) > 0 {
		stb.WriteString(fmt.Sprintf("%s ", strings.Join(c.Envs, " ")))
	}
	if c.ShellType != "" && c.ShellFlag != "" {
		stb.WriteString(fmt.Sprintf("%s %s '%s'", c.ShellType, c.ShellFlag, escapeSingleQuotes(formatCommandMessage(name, args))))
	} else {
		stb.WriteString(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
	}
	return stb.String()
}
