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

func ExecInPath(path string, name string, args ...string) ([]byte, error) {
	if path == "" {
		return nil, erero.New("CAN NOT EXEC IN BLANK DIRECTORY PATH")
	}
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN NAME STRING")
	}
	zaplog.ZAPS.P1.LOG.Debug("EXEC_IN_PATH:", zap.String("CMD", cmsExecInPath(path, name, args...)))
	cmd := exec.Command(name, args...)
	cmd.Dir = path
	return cmd.CombinedOutput()
}

func ExecInEnvs(envs []string, name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN NAME STRING")
	}
	zaplog.ZAPS.P1.LOG.Debug("EXEC_IN_ENVS:", zap.String("CMD", cmsExecInEnvs(envs, name, args...)))
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
		return nil, erero.New("CAN NOT EXECUTE WITH WRONG SHELL FLAG")
	}
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN NAME STRING")
	}
	zaplog.ZAPS.P1.LOG.Debug("EXEC_XSH_RUN:", zap.String("CMD", cmsExecXshRun(shellType, shellFlag, name, args...)))
	sub := name + " " + strings.Join(args, " ")
	cmd := exec.Command(shellType, "-c", sub)
	return cmd.CombinedOutput()
}

func cmsExecInPath(path string, name string, args ...string) string {
	return strings.TrimSpace(fmt.Sprintf("cd %s && %s", path, cvtExecToBash(name, args...)))
}

func cvtExecToBash(name string, args ...string) string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", name, strings.Join(args, " ")))
}

func cmsExecInEnvs(envs []string, name string, args ...string) string {
	return strings.TrimSpace(fmt.Sprintf("%s %s", strings.Join(envs, " "), cvtExecToBash(name, args...)))
}

func cmsExecXshRun(shellType, shellFlag string, name string, args ...string) string {
	return strings.TrimSpace(fmt.Sprintf("%s %s '%s'", shellType, shellFlag, cvtExecToBash(name, args...)))
}

func Exec(cmx *CMX, name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN NAME STRING")
	}
	var cmd *exec.Cmd
	if cmx.ShellType != "" {
		if cmx.ShellFlag != "-c" { //这里暂不支持别的
			return nil, erero.New("CAN NOT EXECUTE WITH WRONG SHELL FLAG")
		}
		sub := name + " " + strings.Join(args, " ")
		cmd = exec.Command(cmx.ShellType, cmx.ShellFlag, sub)
	} else {
		cmd = exec.Command(name, args...)
	}
	if cmx.Path != "" {
		cmd.Dir = cmx.Path
	}
	if len(cmx.Envs) > 0 {
		cmd.Env = append(os.Environ(), cmx.Envs...)
	}
	return cmd.CombinedOutput()
}
