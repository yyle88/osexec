package osexec

type CMX struct {
	Envs      []string // 可选的环境变量
	Path      string   // 可选的执行路径
	ShellType string
	ShellFlag string
}

func NewCMX() *CMX {
	return &CMX{}
}

// WithEnvs 设置环境变量并返回 CMX，以支持链式调用。
func (c *CMX) WithEnvs(envs []string) *CMX {
	c.Envs = envs
	return c
}

// WithPath 设置执行路径并返回 CMX，以支持链式调用。
func (c *CMX) WithPath(path string) *CMX {
	c.Path = path
	return c
}

// WithShellType 设置 shell 类型并返回 CMX，以支持链式调用。
func (c *CMX) WithShellType(shellType string) *CMX {
	c.ShellType = shellType
	return c
}

// WithShellFlag 设置 shell 标志并返回 CMX，以支持链式调用。
func (c *CMX) WithShellFlag(shellFlag string) *CMX {
	c.ShellFlag = shellFlag
	return c
}

func (cmx *CMX) Exec(name string, args ...string) ([]byte, error) {
	return Exec(cmx, name, args...)
}
