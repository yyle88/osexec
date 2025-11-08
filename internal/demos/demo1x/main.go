package main

import (
	"fmt"

	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
)

func main() {
	// Execute simple command
	// 执行简单命令
	output, err := osexec.Exec("echo", "abc")
	must.Done(err)
	fmt.Println("Output:", string(output))

	// Execute command in specific path
	// 在指定路径中执行命令
	output, err = osexec.ExecInPath("/tmp", "pwd")
	must.Done(err)
	fmt.Println("Current path:", string(output))

	// Execute with environment variables
	// 使用环境变量执行命令
	envs := []string{"MY_VAR=hello", "ANOTHER_VAR=world"}
	output, err = osexec.ExecInEnvs(envs, "printenv", "MY_VAR")
	must.Done(err)
	fmt.Println("Environment value:", string(output))
}
