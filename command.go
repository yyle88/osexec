package osexec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/printgo"
	"github.com/yyle88/tern"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// CommandConfig represents the configuration for executing shell commands.
// CommandConfig 表示执行 shell 命令的配置。
type CommandConfig struct {
	Envs      []string // Optional environment variables. // 填写可选的环境变量。
	Path      string   // Optional execution path. // 填写可选的执行路径。
	ShellType string   // Optional type of shell to use, e.g., bash, zsh. // 填写可选的 shell 类型，例如 bash，zsh。
	ShellFlag string   // Optional shell flag, e.g., "-c". // 填写可选的 Shell 参数，例如 "-c"。
	DebugMode bool     // enable debug mode. // 是否启用调试模式。
}

// NewCommandConfig creates and returns a new CommandConfig instance.
// NewCommandConfig 创建并返回一个新的 CommandConfig 实例。
func NewCommandConfig() *CommandConfig {
	return &CommandConfig{
		DebugMode: debugModeOpen, // Initial value is consistent with the debugModeOpen variable. // 初始值与 debugModeOpen 变量保持一致。
	}
}

// WithEnvs sets the environment variables for CommandConfig and returns the updated instance.
// WithEnvs 设置 CommandConfig 的环境变量并返回更新后的实例。
func (c *CommandConfig) WithEnvs(envs []string) *CommandConfig {
	c.Envs = envs
	return c
}

// WithPath sets the execution path for CommandConfig and returns the updated instance.
// WithPath 设置 CommandConfig 的执行路径并返回更新后的实例。
func (c *CommandConfig) WithPath(path string) *CommandConfig {
	c.Path = path
	return c
}

// WithShellType sets the shell type for CommandConfig and returns the updated instance.
// WithShellType 设置 CommandConfig 的 shell 类型并返回更新后的实例。
func (c *CommandConfig) WithShellType(shellType string) *CommandConfig {
	c.ShellType = shellType
	return c
}

// WithShellFlag sets the shell flag for CommandConfig and returns the updated instance.
// WithShellFlag 设置 CommandConfig 的 shell 参数并返回更新后的实例。
func (c *CommandConfig) WithShellFlag(shellFlag string) *CommandConfig {
	c.ShellFlag = shellFlag
	return c
}

// WithShell sets both the shell type and shell flag for CommandConfig and returns the updated instance.
// WithShell 同时设置 CommandConfig 的 shell 类型和 shell 参数，并返回更新后的实例。
func (c *CommandConfig) WithShell(shellType, shellFlag string) *CommandConfig {
	c.ShellType = shellType
	c.ShellFlag = shellFlag
	return c
}

// WithBash sets the shell to bash with the "-c" flag and returns the updated instance.
// WithBash 设置 shell 为 bash 并附带 "-c" 参数，返回更新后的实例。
func (c *CommandConfig) WithBash() *CommandConfig {
	return c.WithShell("bash", "-c")
}

// WithZsh sets the shell to zsh with the "-c" flag and returns the updated instance.
// WithZsh 设置 shell 为 zsh 并附带 "-c" 参数，返回更新后的实例。
func (c *CommandConfig) WithZsh() *CommandConfig {
	return c.WithShell("zsh", "-c")
}

// WithSh sets the shell to sh with the "-c" flag and returns the updated instance.
// WithSh 设置 shell 为 sh 并附带 "-c" 参数，返回更新后的实例。
func (c *CommandConfig) WithSh() *CommandConfig {
	return c.WithShell("sh", "-c")
}

// WithDebugMode sets the debug mode for CommandConfig and returns the updated instance.
// WithDebugMode 设置 CommandConfig 的调试模式并返回更新后的实例。
func (c *CommandConfig) WithDebugMode(debugMode bool) *CommandConfig {
	c.DebugMode = debugMode
	return c
}

// Exec executes a shell command with the specified name and arguments, using the CommandConfig configuration.
// Exec 使用 CommandConfig 的配置执行带有指定名称和参数的 shell 命令。
func (c *CommandConfig) Exec(name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("can-not-execute-with-empty-command-name")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("can-not-contains-space-in-command-name")
	}
	if c.ShellFlag != "" {
		if c.ShellType == "" {
			return nil, erero.New("can-not-execute-with-wrong-shell-command")
		}
	}
	if c.ShellType != "" {
		if c.ShellFlag != "-c" {
			return nil, erero.New("can-not-execute-with-wrong-shell-options")
		}
	}
	if c.DebugMode {
		debugMessage := c.makeCommandMessage(name, args)
		showMessage(debugMessage)
		zaplog.ZAPS.P1.LOG.Debug("EXEC:", zap.String("CMD", debugMessage))
	}
	cmd := tern.BFF(c.ShellType != "",
		func() *exec.Cmd {
			return exec.Command(c.ShellType, c.ShellFlag, name+" "+strings.Join(args, " "))
		},
		func() *exec.Cmd {
			return exec.Command(name, args...)
		})
	cmd.Dir = c.Path
	cmd.Env = tern.BF(len(c.Envs) > 0, func() []string {
		return append(os.Environ(), c.Envs...)
	})
	return c.warpCommandOutput(cmd.CombinedOutput())
}

// warpCommandOutput processes the output and error from the command execution, adding debug information if necessary.
// warpCommandOutput 处理命令执行的输出和错误，并在需要时添加调试信息。
func (c *CommandConfig) warpCommandOutput(output []byte, erx error) ([]byte, error) {
	if erx != nil {
		if c.DebugMode {
			if len(output) > 0 {
				showWarning(string(output))
			} else {
				showWarning(erx.Error())
			}
		}
		return output, erero.Wro(erx)
	}
	return output, nil
}

// makeCommandMessage constructs a command-line string based on the CommandConfig and given command name and arguments.
// makeCommandMessage 根据 CommandConfig 和指定的命令名称及参数构造命令行字符串。
func (c *CommandConfig) makeCommandMessage(name string, args []string) string {
	var pts = printgo.NewPTS()
	if c.Path != "" {
		pts.WriteString(fmt.Sprintf("cd %s && ", c.Path))
	}
	if len(c.Envs) > 0 {
		pts.WriteString(fmt.Sprintf("%s ", strings.Join(c.Envs, " ")))
	}
	if c.ShellType != "" && c.ShellFlag != "" {
		pts.WriteString(fmt.Sprintf("%s %s '%s'", c.ShellType, c.ShellFlag, escapeSingleQuotes(makeCommandMessage(name, args))))
	} else {
		pts.WriteString(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
	}
	return pts.String()
}
