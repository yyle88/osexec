package osexec

import (
	"os"
	"os/exec"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/tern"
	"github.com/yyle88/tern/zerotern"
)

type CMC struct {
	Envs      []string // 可选的环境变量
	Path      string   // 可选的执行路径
	ShellType string
	ShellFlag string
}

func NewCMC() *CMC {
	return &CMC{}
}

// WithEnvs 设置环境变量并返回 CMC，以支持链式调用。
func (M *CMC) WithEnvs(envs []string) *CMC {
	M.Envs = envs
	return M
}

// WithPath 设置执行路径并返回 CMC，以支持链式调用。
func (M *CMC) WithPath(path string) *CMC {
	M.Path = path
	return M
}

// WithShellType 设置 shell 类型并返回 CMC，以支持链式调用。
func (M *CMC) WithShellType(shellType string) *CMC {
	M.ShellType = shellType
	return M
}

// WithShellFlag 设置 shell 标志并返回 CMC，以支持链式调用。
func (M *CMC) WithShellFlag(shellFlag string) *CMC {
	M.ShellFlag = shellFlag
	return M
}

func (M *CMC) WithBash() *CMC {
	M.ShellType = "bash"
	M.ShellFlag = "-c"
	return M
}

func (M *CMC) WithZsh() *CMC {
	M.ShellType = "zsh"
	M.ShellFlag = "-c"
	return M
}

func (M *CMC) WithSh() *CMC {
	M.ShellType = "sh"
	M.ShellFlag = "-c"
	return M
}

func (M *CMC) Exec(name string, args ...string) ([]byte, error) {
	if name == "" {
		return nil, erero.New("CAN NOT EXECUTE WITH EMPTY COMMAND NAME")
	}
	if strings.Contains(name, " ") {
		return nil, erero.New("CAN NOT CONTAINS SPACE IN NAME STRING")
	}
	if M.ShellType != "" && M.ShellFlag != "-c" { //这里暂不支持别的
		return nil, erero.New("CAN NOT EXECUTE WITH WRONG SHELL FLAG")
	}

	cmd := tern.BFF(M.ShellType != "", func() *exec.Cmd {
		sub := name + " " + strings.Join(args, " ")
		return exec.Command(M.ShellType, M.ShellFlag, sub)
	}, func() *exec.Cmd {
		return exec.Command(name, args...)
	})
	cmd.Dir = zerotern.VV(M.Path, "") //这样目的在于提示 path 很有可能是空的
	cmd.Env = tern.BF(len(M.Envs) > 0, func() []string {
		return append(os.Environ(), M.Envs...)
	})
	return cmd.CombinedOutput()
}