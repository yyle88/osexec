package osexec

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"slices"
	"strings"
	"sync"

	"github.com/yyle88/done"
	"github.com/yyle88/erero"
	"github.com/yyle88/eroticgo"
	"github.com/yyle88/osexec/internal/utils"
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
	DebugMode bool     // enable debug mode. // 是否启用调试模式，即打印调试的日志。
	MatchPipe func(line string) bool
	MatchMore bool
	TakeExits map[int]string
}

// NewCommandConfig creates and returns a new CommandConfig instance.
// NewCommandConfig 创建并返回一个新的 CommandConfig 实例。
func NewCommandConfig() *CommandConfig {
	return &CommandConfig{
		DebugMode: debugModeOpen, // Initial value is consistent with the debugModeOpen variable. // 初始值与 debugModeOpen 变量保持一致。
		TakeExits: make(map[int]string),
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

// WithDebug sets the debug mode to true for CommandConfig and returns the updated instance.
// WithDebug 将 CommandConfig 的调试模式设置为 true 并返回更新后的实例。
func (c *CommandConfig) WithDebug() *CommandConfig {
	return c.WithDebugMode(true)
}

// WithMatchPipe sets the match pipe function for CommandConfig and returns the updated instance.
// WithMatchPipe 设置 CommandConfig 的匹配管道函数并返回更新后的实例。
func (c *CommandConfig) WithMatchPipe(matchPipe func(line string) bool) *CommandConfig {
	c.MatchPipe = matchPipe
	return c
}

// WithMatchMore sets the match more flag for CommandConfig and returns the updated instance.
// WithMatchMore 设置 CommandConfig 的匹配更多标志并返回更新后的实例。
func (c *CommandConfig) WithMatchMore(matchMore bool) *CommandConfig {
	c.MatchMore = matchMore
	return c
}

// WithTakeExits sets the accepted exit codes for CommandConfig and returns the updated instance.
// WithTakeExits 设置 CommandConfig 的接受退出码集合并返回更新后的实例。
func (c *CommandConfig) WithTakeExits(takeExits map[int]string) *CommandConfig {
	//这里需要复制 map 避免出问题，其次是不要使用 clone 以免外面传的是 nil 就不好啦
	expMap := make(map[int]string, len(takeExits))
	for k, v := range takeExits {
		expMap[k] = v
	}
	//这里完全覆盖而不是增补，是因为覆盖更符合预期，否则还得写增补逻辑
	c.TakeExits = expMap
	return c
}

// WithExpectExit adds an expected exit code to CommandConfig and returns the updated instance.
// WithExpectExit 向 CommandConfig 添加一个期望的退出码并返回更新后的实例。
func (c *CommandConfig) WithExpectExit(exitCode int, reason string) *CommandConfig {
	c.TakeExits[exitCode] = reason
	return c
}

// WithExpectCode adds an expected exit code to CommandConfig and returns the updated instance.
// WithExpectCode 向 CommandConfig 添加一个期望的退出码并返回更新后的实例。
func (c *CommandConfig) WithExpectCode(exitCode int) *CommandConfig {
	c.TakeExits[exitCode] = "EXPECTED-EXIT-CODES"
	return c
}

// Exec executes a shell command with the specified name and arguments, using the CommandConfig configuration.
// Exec 使用 CommandConfig 的配置执行带有指定名称和参数的 shell 命令。
func (c *CommandConfig) Exec(name string, args ...string) ([]byte, error) {
	if err := c.validateConfig(name, args); err != nil {
		return nil, erero.Ero(err)
	}
	command := c.prepareCommand(name, args)
	return utils.WarpResults(done.VAE(command.CombinedOutput()), c.DebugMode, c.TakeExits)
}

func (c *CommandConfig) validateConfig(name string, args []string) error {
	if name == "" {
		return erero.New("can-not-execute-with-empty-command-name")
	}
	if c.ShellFlag == "" && c.ShellType == "" {
		if strings.Contains(name, " ") {
			return erero.New("can-not-contains-space-in-command-name")
		}
	}
	if c.ShellFlag != "" {
		if c.ShellType == "" {
			return erero.New("can-not-execute-with-wrong-shell-command")
		}
	}
	if c.ShellType != "" {
		if c.ShellFlag != "-c" {
			return erero.New("can-not-execute-with-wrong-shell-options")
		}
	}
	if c.DebugMode {
		debugMessage := c.makeCommandMessage(name, args)
		utils.ShowCommand(debugMessage)
		zaplog.ZAPS.Skip1.LOG.Debug("EXEC:", zap.String("CMD", debugMessage))
	}
	return nil
}

func (c *CommandConfig) prepareCommand(name string, args []string) *exec.Cmd {
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
	return cmd
}

// makeCommandMessage constructs a command-line string based on the CommandConfig and given command name and arguments.
// makeCommandMessage 根据 CommandConfig 和指定的命令名称及参数构造命令行字符串。
func (c *CommandConfig) makeCommandMessage(name string, args []string) string {
	var pts = printgo.NewPTS()
	if c.Path != "" {
		pts.Fprintf("cd %s && ", c.Path)
	}
	if len(c.Envs) > 0 {
		pts.Fprintf("%s ", strings.Join(c.Envs, " "))
	}
	if c.ShellType != "" && c.ShellFlag != "" {
		pts.Fprintf("%s %s '%s'", c.ShellType, c.ShellFlag, escapeSingleQuotes(makeCommandMessage(name, args)))
	} else {
		pts.Fprintf("%s %s", name, strings.Join(args, " "))
	}
	return pts.String()
}

// StreamExec executes a shell command with the specified name and arguments, using the CommandConfig configuration, and returns the output as a byte slice.
// StreamExec 使用 CommandConfig 的配置执行带有指定名称和参数的 shell 命令，并返回输出的字节切片。
func (c *CommandConfig) StreamExec(name string, args ...string) ([]byte, error) {
	return c.ExecInPipe(name, args...)
}

// ExecInPipe executes a shell command with the specified name and arguments, using the CommandConfig configuration, and returns the output as a byte slice.
// ExecInPipe 使用 CommandConfig 的配置执行带有指定名称和参数的 shell 命令，并返回输出的字节切片。
func (c *CommandConfig) ExecInPipe(name string, args ...string) ([]byte, error) {
	if err := c.validateConfig(name, args); err != nil {
		return nil, erero.Ero(err)
	}
	command := c.prepareCommand(name, args)

	stdoutPipe, err := command.StdoutPipe()
	if err != nil {
		return nil, erero.Wro(err)
	}

	stderrPipe, err := command.StderrPipe()
	if err != nil {
		return nil, erero.Wro(err)
	}

	stdoutReader := bufio.NewReader(stdoutPipe)
	stderrReader := bufio.NewReader(stderrPipe)
	if err := command.Start(); err != nil {
		return nil, erero.Wro(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(2)
	var errMatch = false
	var stderrBuffer = printgo.NewPTX()
	go func() {
		defer wg.Done()
		errMatch = c.readPipe(stderrReader, stderrBuffer, "REASON", eroticgo.RED)
	}()
	var outMatch = false
	var stdoutBuffer = printgo.NewPTX()
	go func() {
		defer wg.Done()
		outMatch = c.readPipe(stdoutReader, stdoutBuffer, "OUTPUT", eroticgo.GREEN)
	}()
	wg.Wait()

	if outMatch {
		return utils.WarpMessage(done.VAE(stdoutBuffer.Bytes(), nil), c.DebugMode)
	}

	if errMatch { //比如 "go: upgraded github.com/xx/xx vxx => vxx" 这就不算错误，而是正确的
		return utils.WarpMessage(done.VAE(stderrBuffer.Bytes(), nil), c.DebugMode)
	}

	if stderrBuffer.Len() > 0 {
		return utils.WarpMessage(done.VAE(stdoutBuffer.Bytes(), erero.New(stderrBuffer.String())), c.DebugMode)
	} else {
		return utils.WarpMessage(done.VAE(stdoutBuffer.Bytes(), nil), c.DebugMode)
	}
}

// readPipe reads from the provided reader and writes to the provided PTX buffer, using the specified debug message and colors.
// readPipe 从提供的 reader 读取数据并写入提供的 PTX 缓冲区，使用指定的调试消息和颜色。
func (c *CommandConfig) readPipe(reader *bufio.Reader, ptx *printgo.PTX, debugMessage string, erotic eroticgo.COLOR) (matched bool) {
	for {
		streamLine, _, err := reader.ReadLine()

		if c.DebugMode {
			zaplog.SUG.Debugln(debugMessage, erotic.Sprint(string(streamLine)))
		}

		if (c.MatchMore || !matched) && c.MatchPipe != nil {
			if c.MatchPipe(string(streamLine)) {
				matched = true
			}
		}

		if err != nil {
			if err == io.EOF {
				ptx.Write(streamLine)
				return matched
			}
			panic(erero.Wro(err)) //panic: 读取结果出错很罕见
		} else {
			ptx.Write(streamLine)
			ptx.Println()
		}
	}
}

// ShallowClone creates a shallow copy of the CommandConfig instance.
// ShallowClone 拷贝个新的 CommandConfig 实例，以便于实现总配置和子配置分隔.
func (c *CommandConfig) ShallowClone() *CommandConfig {
	newConfig := new(CommandConfig)
	*newConfig = *c
	newConfig.Envs = slices.Clone(c.Envs)      //这里为了避免踩内存还是得拷贝一份
	newConfig.TakeExits = make(map[int]string) //这里很简单因为不同的子命令期望的错误码不同，这里就不克隆这个“有预期的错误码表”，避免错误被忽略
	return newConfig
}

// GetSubClone creates a shallow copy of the CommandConfig instance with a new path and returns the updated instance.
// GetSubClone 创建一个带有新路径的 CommandConfig 实例的浅拷贝并返回更新后的实例。
func (c *CommandConfig) GetSubClone(path string) *CommandConfig {
	return c.ShallowClone().WithPath(path)
}
