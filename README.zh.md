# osexec

调用 Golang 的 `os/exec` 包的简单工具。

## 英文文档

[English README](README.md)

## 功能

- **自定义执行配置**：支持通过自定义环境变量、工作路径和 Shell 选项来执行命令。

## 安装

```bash  
go get github.com/yyle88/osexec  
```  

## `CommandConfig` 结构体及其方法

`CommandConfig` 结构体提供了一种灵活的方式来配置和执行命令。您可以通过链式接口设置自定义的环境变量、工作路径、Shell 类型以及调试选项。

### `NewCommandConfig() *CommandConfig`

创建并返回一个新的 `CommandConfig` 实例。

#### 示例：

```go  
config := osexec.NewCommandConfig()
```  

### 链式方法

- **WithEnvs(envs []string) *CommandConfig**：设置自定义环境变量。
- **WithPath(path string) *CommandConfig**：设置工作路径。
- **WithShellType(shellType string) *CommandConfig**：设置 Shell 类型（例如，`bash`）。
- **WithShellFlag(shellFlag string) *CommandConfig**：设置 Shell 标志（例如，`-c`）。
- **WithShell(shellType, shellFlag string) *CommandConfig**：设置 Shell 类型和标志。
- **WithBash() *CommandConfig**：将命令配置为使用 `bash -c`。
- **WithZsh() *CommandConfig**：将命令配置为使用 `zsh -c`。
- **WithSh() *CommandConfig**：将命令配置为使用 `sh -c`。
- **WithDebugMode(debugMode bool) *CommandConfig**：启用或禁用调试模式。

#### 示例：

```go
package main

import (
	"fmt"
	"github.com/yyle88/osexec"
)

func main() {
	// 创建一个新的 CommandConfig 实例，设置工作目录和调试模式
	config := osexec.NewCommandConfig().
		WithPath("/path/to/directoryName").
		WithDebugMode(true)

	output, err := config.Exec("echo", "Hello, World!")
	if err != nil {
		fmt.Println("Reason:", err)
	} else {
		fmt.Println("Output:", string(output))
	}
}
```

---

## 许可

项目采用 MIT 许可证，详情请参阅 [LICENSE](LICENSE)。

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
