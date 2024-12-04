package osexec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func Exec(name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN COMMAND NAME")
	}
	if enableDebug {
		debugMessage := strings.TrimSpace(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
		zaplog.ZAPS.P1.LOG.Debug("EXEC:", zap.String("CMD", debugMessage))
	}
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}

func ExecInPath(path string, name string, args ...string) ([]byte, error) {
	if path == "" {
		return nil, erero.New("CAN NOT EXEC IN BLANK DIRECTORY PATH")
	}
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN COMMAND NAME")
	}
	if enableDebug {
		debugMessage := strings.TrimSpace(fmt.Sprintf("cd %s && %s", path, formatCommandMessage(name, args)))
		zaplog.ZAPS.P1.LOG.Debug("EXEC_IN_PATH:", zap.String("CMD", debugMessage))
	}
	cmd := exec.Command(name, args...)
	cmd.Dir = path
	return cmd.CombinedOutput()
}

func ExecInEnvs(envs []string, name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN COMMAND NAME")
	}
	if enableDebug {
		debugMessage := strings.TrimSpace(fmt.Sprintf("%s %s", strings.Join(envs, " "), formatCommandMessage(name, args)))
		zaplog.ZAPS.P1.LOG.Debug("EXEC_IN_ENVS:", zap.String("CMD", debugMessage))
	}
	cmd := exec.Command(name, args...)
	cmd.Env = os.Environ() //这里没事是安全的
	cmd.Env = append(cmd.Env, envs...)
	return cmd.CombinedOutput()
}

func ExecXshRun(shellType, shellFlag string, name string, args ...string) ([]byte, error) {
	if shellType == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY SHELL TYPE")
	}
	if shellFlag != "-c" { //这里虽然必然要填 -c 但是依然要作为参数，否则外部调用时就会很怪异
		return nil, erero.New("CAN NOT EXECUTE WITH WRONG SHELL OPTIONS")
	}
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN COMMAND NAME")
	}
	if enableDebug {
		debugMessage := strings.TrimSpace(fmt.Sprintf("%s %s '%s'", shellType, shellFlag, escapeSingleQuotes(formatCommandMessage(name, args))))
		zaplog.ZAPS.P1.LOG.Debug("EXEC_XSH_RUN:", zap.String("CMD", debugMessage))
	}
	cmd := exec.Command(shellType, "-c", name+" "+strings.Join(args, " "))
	return cmd.CombinedOutput()
}

// 把命令和参数拼接为命令行字符串
func formatCommandMessage(name string, args []string) string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
}

// 把单引号转义为 '\”（shell 中安全嵌套单引号的方式）
func escapeSingleQuotes(input string) string {
	return strings.ReplaceAll(input, "'", `'\''`)
}
