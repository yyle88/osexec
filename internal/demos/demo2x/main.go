package main

import (
	"fmt"

	"github.com/yyle88/done"
	"github.com/yyle88/osexec"
)

func main() {
	// Create configuration with chainable methods
	// 使用链式方法创建配置
	config := osexec.NewCommandConfig().
		WithBash().
		WithDebugMode(osexec.SHOW_COMMAND)

	// Execute shell command
	// 执行 shell 命令
	output, err := config.Exec("echo $HOME")
	done.Done(err)
	fmt.Println("Home path:", string(output))

	// Execute command with custom environment
	// 使用自定义环境执行命令
	config = config.NewConfig().
		WithEnvs([]string{"GREETING=Hello", "NAME=Go"}).
		WithBash()

	output, err = config.Exec("echo", "$GREETING $NAME!")
	done.Done(err)
	fmt.Println("Message:", string(output))
}
