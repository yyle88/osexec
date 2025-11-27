[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/osexec/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/osexec/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/osexec)](https://pkg.go.dev/github.com/yyle88/osexec)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/osexec/main.svg)](https://coveralls.io/github/yyle88/osexec?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://github.com/yyle88/osexec)
[![GitHub Release](https://img.shields.io/github/release/yyle88/osexec.svg)](https://github.com/yyle88/osexec/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/osexec)](https://goreportcard.com/report/github.com/yyle88/osexec)

# osexec

调用 Golang `os/exec` 包的简单工具

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 功能特性

- **自定义执行配置**：支持通过自定义环境变量、工作路径和 shell 选项来执行命令
- **链式 API**：构建命令配置时使用流畅的接口
- **Shell 支持**：内置支持 bash、zsh 和 sh shells
- **调试模式**：多种调试级别管理命令和输出
- **退出码处理**：接受特定退出码作为成功
- **环境变量**：简单管理环境变量
- **路径管理**：在特定路径中执行命令

## 安装

```bash
go get github.com/yyle88/osexec
```

## 快速开始

### 基础用法

```go
package main

import (
	"fmt"

	"github.com/yyle88/must"
	"github.com/yyle88/osexec"
)

func main() {
	// Execute simple command
	output, err := osexec.Exec("echo", "abc")
	must.Done(err)
	fmt.Println("Output:", string(output))

	// Execute command in specific path
	output, err = osexec.ExecInPath("/tmp", "pwd")
	must.Done(err)
	fmt.Println("Current path:", string(output))

	// Execute with environment variables
	envs := []string{"MY_VAR=hello", "ANOTHER_VAR=world"}
	output, err = osexec.ExecInEnvs(envs, "printenv", "MY_VAR")
	must.Done(err)
	fmt.Println("Environment value:", string(output))
}
```

⬆️ **[源码](internal/demos/demo1x/main.go)**

### 高级用法

```go
package main

import (
	"fmt"

	"github.com/yyle88/done"
	"github.com/yyle88/osexec"
)

func main() {
	// Create configuration with chainable methods
	config := osexec.NewCommandConfig().
		WithBash().
		WithDebugMode(osexec.SHOW_COMMAND)

	// Execute shell command
	output, err := config.Exec("echo $HOME")
	done.Done(err)
	fmt.Println("Home path:", string(output))

	// Execute command with custom environment
	config = config.NewConfig().
		WithEnvs([]string{"GREETING=Hello", "NAME=Go"}).
		WithBash()

	output, err = config.Exec("echo", "$GREETING $NAME!")
	done.Done(err)
	fmt.Println("Message:", string(output))
}
```

⬆️ **[源码](internal/demos/demo2x/main.go)**

## CommandConfig - 高级用法

`CommandConfig` 提供灵活的方法来配置和执行命令，支持链式方法调用

### 创建配置

```go
config := osexec.NewCommandConfig()
```

### Shell 执行

使用不同的 shell 执行命令：

```go
// 使用 bash
config := osexec.NewCommandConfig().WithBash()
output, err := config.Exec("echo $SHELL")

// 使用 zsh
config := osexec.NewCommandConfig().WithZsh()
output, err := config.Exec("echo 'ZSH Command'")

// 使用 sh
config := osexec.NewCommandConfig().WithSh()
output, err := config.Exec("pwd")
```

### 复杂 Shell 命令

```go
config := osexec.NewCommandConfig().WithBash()

// 管道命令
output, err := config.Exec("echo 'apple\nbanana\norange' | grep 'banana'")

// 带变量的命令
config.WithEnvs([]string{"GREETING=Hello", "NAME=World"})
output, err = config.Exec("echo", "$GREETING $NAME!")
```

### 调试模式

管理命令和输出选项：

```go
// 完整调试模式 - 显示命令和输出
config := osexec.NewCommandConfig().WithDebug()

// 仅显示命令
config := osexec.NewCommandConfig().WithDebugMode(osexec.SHOW_COMMAND)

// 仅显示输出
config := osexec.NewCommandConfig().WithDebugMode(osexec.SHOW_OUTPUTS)

// 静默模式 - 无调试输出
config := osexec.NewCommandConfig().WithDebugMode(osexec.QUIET)
```

### 退出码处理

接受特定退出码作为成功：

```go
// 接受退出码 1 作为成功
config := osexec.NewCommandConfig().
	WithExpectExit(1, "DIFFERENCES FOUND")

output, err := config.Exec("diff", "file1.txt", "file2.txt")
// 即使 diff 返回退出码 1，err 也会是 nil

// 接受多个退出码
config := osexec.NewCommandConfig().
	WithTakeExits(map[int]string{
		1: "DIFFERENCES FOUND",
		2: "TROUBLE",
	})
```

### 获取退出码

使用 `ExecTake` 获取退出码进行精细控制：

```go
// ExecTake 返回输出、退出码和错误
output, exitCode, err := osexec.NewCommandConfig().
	WithExpectCode(1).
	ExecTake("diff", "file1.txt", "file2.txt")

// 文件不同时 exitCode = 1，文件相同时 exitCode = 0
fmt.Println("Exit code:", exitCode)
```

### 链式配置

组合多个配置选项：

```go
config := osexec.NewCommandConfig().
	WithPath("/path/to/project").
	WithEnvs([]string{"ENV=production"}).
	WithBash().
	WithDebugMode(osexec.SHOW_COMMAND).
	WithExpectCode(1)

output, err := config.Exec("command-name", "arg1", "arg2")
```

## API 参考

### 配置方法

- **WithEnvs(envs []string)**：设置自定义环境变量
- **WithPath(path string)**：设置工作路径
- **WithShellType(shellType string)**：设置 shell 类型（如 `bash`）
- **WithShellFlag(shellFlag string)**：设置 shell 参数（如 `-c`）
- **WithShell(shellType, shellFlag string)**：同时设置 shell 类型和参数
- **WithBash()**：配置使用 `bash -c`
- **WithZsh()**：配置使用 `zsh -c`
- **WithSh()**：配置使用 `sh -c`
- **WithDebug()**：启用完整调试模式
- **WithDebugMode(debugMode DebugMode)**：设置特定调试模式
- **WithExpectExit(exitCode int, reason string)**：添加期望的退出码及原因
- **WithExpectCode(exitCode int)**：添加期望的退出码
- **WithTakeExits(takeExits map[int]string)**：设置多个期望的退出码

### 执行方法

- **Exec(name string, args ...string)**：执行命令并返回输出
- **ExecTake(name string, args ...string)**：执行命令并返回输出、退出码和错误
- **ExecWith(name string, args []string, prepare func(*exec.Cmd))**：使用自定义命令设置执行
- **StreamExec(name string, args ...string)**：使用管道处理执行命令
- **ExecInPipe(name string, args ...string)**：使用 stdout/stderr 管道处理执行

### 调试模式

- **QUIET**：无调试输出
- **DEBUG**：完整调试模式，显示命令和输出
- **SHOW_COMMAND**：仅显示命令
- **SHOW_OUTPUTS**：仅显示输出

## 测试工具

`osexectest` 包提供了用于编写命令执行测试的辅助函数。

### 因缺少命令而跳过测试

在编写依赖外部命令（例如 `zsh`、`git`、`tree`）的测试时，如果测试环境中缺少所需命令，最佳实践是跳过这些测试。`SkipIfCommandNotFound` 函数可以帮助您轻松实现这一点。

```go
package my_test

import (
    "testing"

    "github.com/yyle88/osexec/osexectest"
)

func TestSomethingThatNeedsZsh(t *testing.T) {
    // 如果系统中未安装 'zsh'，此测试将自动跳过。
    osexectest.SkipIfCommandNotFound(t, "zsh")

    // ... 使用 'zsh' 的其余测试代码
}
```
这可以避免在未安装特定命令行工具的环境中出现测试失败。

---

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-25 03:52:28.131064 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们完善文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，协助解决性能问题
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：面向用户的更改需要更新文档
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来贡献此项目。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub 标星点赞

[![Stargazers](https://starchart.cc/yyle88/osexec.svg?variant=adaptive)](https://starchart.cc/yyle88/osexec)
