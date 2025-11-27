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

// CommandConfig represents the configuration when executing shell commands
// CommandConfig 表示执行 shell 命令时的配置
type CommandConfig struct {
	Envs      []string                     // Custom environment variables // 自定义环境变量
	Path      string                       // Execution path // 执行路径
	ShellType string                       // Type of shell to use, e.g., bash, zsh // shell 类型，例如 bash，zsh
	ShellFlag string                       // Shell flag, e.g., "-c" // Shell 参数，例如 "-c"
	DebugMode DebugMode                    // Debug mode setting // 调试模式设置
	MatchPipe func(outputLine string) bool // Function to match output lines in pipe mode // 管道模式下匹配输出行的函数
	MatchMore bool                         // Continue matching even when matched // 即使匹配成功也继续匹配
	TakeExits map[int]string               // Map of expected exit codes with reasons // 预期退出码及其原因的映射表
}

// NewCommandConfig creates and returns a new CommandConfig instance.
// NewCommandConfig 创建并返回一个新的 CommandConfig 实例。
func NewCommandConfig() *CommandConfig {
	return &CommandConfig{
		DebugMode: NewDebugMode(debugModeOpen), // consistent with the debugModeOpen variable. // 初始值与 debugModeOpen 保持一致。
		MatchPipe: func(outputLine string) bool { return false },
		TakeExits: make(map[int]string),
	}
}

// WithEnvs sets the environment variables and returns the updated instance
// WithEnvs 设置环境变量并返回更新后的实例
func (c *CommandConfig) WithEnvs(envs []string) *CommandConfig {
	c.Envs = envs
	return c
}

// WithPath sets the execution path and returns the updated instance
// WithPath 设置执行路径并返回更新后的实例
func (c *CommandConfig) WithPath(path string) *CommandConfig {
	c.Path = path
	return c
}

// WithShellType sets the shell type and returns the updated instance
// WithShellType 设置 shell 类型并返回更新后的实例
func (c *CommandConfig) WithShellType(shellType string) *CommandConfig {
	c.ShellType = shellType
	return c
}

// WithShellFlag sets the shell flag and returns the updated instance
// WithShellFlag 设置 shell 参数并返回更新后的实例
func (c *CommandConfig) WithShellFlag(shellFlag string) *CommandConfig {
	c.ShellFlag = shellFlag
	return c
}

// WithShell sets both the shell type and shell flag, returns the updated instance
// WithShell 同时设置 shell 类型和 shell 参数，返回更新后的实例
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

// WithDebug sets the debug mode to true and returns the updated instance
// WithDebug 将调试模式设置 true 并返回更新后的实例
func (c *CommandConfig) WithDebug() *CommandConfig {
	return c.WithDebugMode(DEBUG)
}

// WithDebugMode sets the debug mode and returns the updated instance
// WithDebugMode 设置调试模式并返回更新后的实例
func (c *CommandConfig) WithDebugMode(debugMode DebugMode) *CommandConfig {
	c.DebugMode = debugMode
	return c
}

// WithMatchPipe sets the match pipe function and returns the updated instance
// WithMatchPipe 设置匹配管道函数并返回更新后的实例
func (c *CommandConfig) WithMatchPipe(matchPipe func(outputLine string) bool) *CommandConfig {
	c.MatchPipe = matchPipe
	return c
}

// WithMatchMore sets the match more flag and returns the updated instance
// WithMatchMore 设置匹配更多标志并返回更新后的实例
func (c *CommandConfig) WithMatchMore(matchMore bool) *CommandConfig {
	c.MatchMore = matchMore
	return c
}

// WithTakeExits sets the accepted exit codes and returns the updated instance
// WithTakeExits 设置接受退出码集合并返回更新后的实例
func (c *CommandConfig) WithTakeExits(takeExits map[int]string) *CommandConfig {
	// Clone map to avoid shared reference issues, and avoid using maps.Clone when input could be nil
	// 复制 map 避免共享引用问题，不使用 maps.Clone 以防输入为 nil
	expMap := make(map[int]string, len(takeExits))
	for k, v := range takeExits {
		expMap[k] = v
	}
	// Replace instead of merge, as replacement matches expected pattern
	// 完全替换而非合并，因为替换更符合预期模式
	c.TakeExits = expMap
	return c
}

// WithExpectExit adds an expected exit code and returns the updated instance
// WithExpectExit 添加期望的退出码并返回更新后的实例
func (c *CommandConfig) WithExpectExit(exitCode int, reason string) *CommandConfig {
	c.TakeExits[exitCode] = reason
	return c
}

// WithExpectCode adds an expected exit code and returns the updated instance
// WithExpectCode 添加期望的退出码并返回更新后的实例
func (c *CommandConfig) WithExpectCode(exitCode int) *CommandConfig {
	c.TakeExits[exitCode] = "EXPECTED-EXIT-CODES"
	return c
}

// Exec executes a shell command with the specified name and arguments, using the CommandConfig configuration.
// Exec 使用 CommandConfig 的配置执行带有指定名称和参数的 shell 命令。
func (c *CommandConfig) Exec(name string, args ...string) ([]byte, error) {
	const skipDepth = 1

	if err := c.checkConfig(name, args, skipDepth+1); err != nil {
		return nil, erero.Ero(err)
	}
	command := c.prepareCommand(name, args)
	return utils.WarpResults(done.VAE(command.CombinedOutput()), c.IsShowOutputs(), c.TakeExits)
}

// ExecWith executes a command with custom exec.Cmd preparation
// Allows setting stdin, extra env vars, and additional cmd fields via prepare callback
//
// ExecWith 执行命令，支持自定义 exec.Cmd 配置
// 通过 prepare 回调可设置 stdin、额外环境变量和其它 cmd 字段
func (c *CommandConfig) ExecWith(name string, args []string, prepare func(command *exec.Cmd)) ([]byte, error) {
	const skipDepth = 1

	if err := c.checkConfig(name, args, skipDepth+1); err != nil {
		return nil, erero.Ero(err)
	}
	command := c.prepareCommand(name, args)
	prepare(command)
	return utils.WarpResults(done.VAE(command.CombinedOutput()), c.IsShowOutputs(), c.TakeExits)
}

// ExecTake executes a command and returns output, exit code, and an issue if one exists
// Returns exit code enabling fine-grained handling of command outcomes
// Exit code 0 indicates success, non-zero indicates various conditions
//
// ExecTake 执行命令并返回输出、退出码和错误（如果有的话）
// 返回退出码以便精细处理命令结果
// 退出码 0 表示成功，非零表示各种情况
func (c *CommandConfig) ExecTake(name string, args ...string) ([]byte, int, error) {
	const skipDepth = 1

	if err := c.checkConfig(name, args, skipDepth+1); err != nil {
		return nil, -1, erero.Ero(err)
	}
	command := c.prepareCommand(name, args)
	return utils.WarpOutcome(done.VAE(command.CombinedOutput()), c.IsShowOutputs(), c.TakeExits)
}

// IsShowCommand checks if the command should be displayed based on the debug mode
// IsShowCommand 检查是否应根据调试模式显示命令
func (c *CommandConfig) IsShowCommand() bool {
	return isShowCommand(c.DebugMode)
}

// IsShowOutputs checks if the command results should be displayed based on the debug mode
// IsShowOutputs 检查是否应根据调试模式显示命令结果
func (c *CommandConfig) IsShowOutputs() bool {
	return isShowOutputs(c.DebugMode)
}

// checkConfig validates the command configuration and shows debug info if needed.
// checkConfig 验证命令配置并在需要时显示调试信息。
func (c *CommandConfig) checkConfig(name string, args []string, skipDepth int) error {
	if name == "" {
		return erero.New("can-not-execute-with-blank-command-name")
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
	if c.IsShowCommand() {
		debugMessage := c.makeCommandMessage(name, args)
		utils.ShowCommand(debugMessage)
		zaplog.ZAPS.Skip(skipDepth).LOG.Debug("EXEC:", zap.String("CMD", debugMessage))
	}
	return nil
}

// prepareCommand creates and configures an exec.Cmd based on the CommandConfig settings.
// prepareCommand 根据 CommandConfig 设置创建并配置 exec.Cmd。
func (c *CommandConfig) prepareCommand(name string, args []string) *exec.Cmd {
	cmd := tern.BFF(c.ShellType != "",
		func() *exec.Cmd {
			return exec.Command(c.ShellType, c.ShellFlag, name+" "+strings.Join(args, " "))
		},
		func() *exec.Cmd {
			return exec.Command(name, args...)
		})
	cmd.Dir = c.Path
	// Set environment variables: when c.Envs has no items, Go uses os.Environ()
	// 设置环境变量：当 c.Envs 没有项目时，Go 使用 os.Environ()
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
	const skipDepth = 1

	if err := c.checkConfig(name, args, skipDepth+1); err != nil {
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

	// Wait for command to complete and get exit status
	// 等待命令完成并获取退出状态
	erw := command.Wait()

	// When output matched, exit with success status (can succeed even if erw != nil)
	// 当输出匹配成功时，以成功状态退出（即使 erw != nil 也可以成功）
	if outMatch {
		return utils.WarpResults(done.VAE(stdoutBuffer.Bytes(), erw), c.IsShowOutputs(), c.TakeExits)
	}

	// If stderr matched, return with stderr data (e.g., "go: upgraded xxx")
	// 如果 stderr 匹配成功，返回 stderr 数据（比如 "go: upgraded xxx"）
	if errMatch {
		return utils.WarpResults(done.VAE(stderrBuffer.Bytes(), erw), c.IsShowOutputs(), c.TakeExits)
	}

	// No match found, check errors in sequence
	// 没有匹配，按顺序检查错误
	if erw != nil {
		// Command failed with non-zero exit code
		// 命令以非零退出码失败
		return utils.WarpResults(done.VAE(stdoutBuffer.Bytes(), erw), c.IsShowOutputs(), c.TakeExits)
	}

	if stderrBuffer.Len() > 0 {
		// Command succeeded but has stderr content
		// 命令成功但有 stderr 内容
		return utils.WarpMessage(done.VAE(stdoutBuffer.Bytes(), erero.New(stderrBuffer.String())), c.IsShowOutputs())
	}

	// Command succeeded with no errors
	// 命令成功且无错误
	return utils.WarpOutputs(stdoutBuffer.Bytes(), c.IsShowOutputs())
}

// readPipe reads from the provided reader and writes to the provided PTX buffer, using the specified debug message and colors.
// readPipe 从提供的 reader 读取数据并写入提供的 PTX 缓冲区，使用指定的调试消息和颜色。
func (c *CommandConfig) readPipe(reader *bufio.Reader, ptx *printgo.PTX, debugMessage string, erotic eroticgo.COLOR) (matched bool) {
	for {
		streamLine, _, err := reader.ReadLine()

		if c.IsShowOutputs() {
			zaplog.SUG.Debugln(debugMessage, erotic.Sprint(string(streamLine)))
		}

		if (c.MatchMore || !matched) && c.MatchPipe != nil && c.MatchPipe(string(streamLine)) {
			matched = true
		}

		if err != nil {
			if err == io.EOF {
				ptx.Write(streamLine)
				return matched
			}
			panic(erero.Wro(err)) // Panic on failure for read error, which is rare // 读取错误时 panic，这种情况很罕见
		}
		ptx.Write(streamLine)
		ptx.Println()
	}
}

// NewConfig creates a shallow clone of the CommandConfig instance.
// NewConfig 克隆一个新的 CommandConfig 实例，以便于实现总配置和子配置分隔.
func (c *CommandConfig) NewConfig() *CommandConfig {
	return &CommandConfig{
		Envs:      slices.Clone(c.Envs),                          // Clone to avoid data sharing issues // 克隆以避免数据共享问题
		Path:      c.Path,                                        // Use same path // 使用相同路径
		ShellType: "",                                            // Each command sets its own // 各命令自行设置
		ShellFlag: "",                                            // Each command sets its own // 各命令自行设置
		DebugMode: c.DebugMode,                                   // Use same debug mode // 使用相同调试模式
		MatchPipe: func(outputLine string) bool { return false }, // Each command sets its own // 各命令自行设置
		MatchMore: false,                                         // Each command sets its own // 各命令自行设置
		TakeExits: make(map[int]string),                          // New map as different commands expect different exit codes // 新建映射表，因不同命令期望不同退出码
	}
}

// SubConfig creates a shallow clone of the CommandConfig instance with a new path and returns the updated instance.
// SubConfig 创建一个带有新路径的 CommandConfig 实例的浅克隆并返回更新后的实例。
func (c *CommandConfig) SubConfig(path string) *CommandConfig {
	return c.NewConfig().WithPath(path)
}
