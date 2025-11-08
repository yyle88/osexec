[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/osexec/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/osexec/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/osexec)](https://pkg.go.dev/github.com/yyle88/osexec)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/osexec/main.svg)](https://coveralls.io/github/yyle88/osexec?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23%2C%201.24%2C%201.25-lightgrey.svg)](https://github.com/yyle88/osexec)
[![GitHub Release](https://img.shields.io/github/release/yyle88/osexec.svg)](https://github.com/yyle88/osexec/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/osexec)](https://goreportcard.com/report/github.com/yyle88/osexec)

# osexec

Simple utilities to use Golang's `os/exec` package.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Features

- **Custom Execution Configurations**: Execute commands with customizable environment variables, working paths, and shell options
- **Chainable API**: Fluent interface when building command configurations
- **Shell Support**: Built-in support with bash, zsh, and sh shells
- **Debug Modes**: Multiple debug levels for command and output management
- **Exit Code Handling**: Accept specific exit codes as success
- **Environment Variables**: Simple environment variable management
- **Path Management**: Execute commands in specific paths

## Installation

```bash
go get github.com/yyle88/osexec
```

## Quick Start

### Basic Usage

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

⬆️ **[Source](internal/demos/demo1x/main.go)**

### Advanced Usage

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

⬆️ **[Source](internal/demos/demo2x/main.go)**

## CommandConfig - Advanced Usage

`CommandConfig` provides a flexible method to configure and execute commands with chainable methods.

### Creating Configuration

```go
config := osexec.NewCommandConfig()
```

### Shell Execution

Execute commands using different shells:

```go
// Using bash
config := osexec.NewCommandConfig().WithBash()
output, err := config.Exec("echo $SHELL")

// Using zsh
config := osexec.NewCommandConfig().WithZsh()
output, err := config.Exec("echo 'ZSH Command'")

// Using sh
config := osexec.NewCommandConfig().WithSh()
output, err := config.Exec("pwd")
```

### Complex Shell Commands

```go
config := osexec.NewCommandConfig().WithBash()

// Pipe commands
output, err := config.Exec("echo 'apple\nbanana\norange' | grep 'banana'")

// Command with variables
config.WithEnvs([]string{"GREETING=Hello", "NAME=World"})
output, err = config.Exec("echo", "$GREETING $NAME!")
```

### Debug Modes

Manage command and output options:

```go
// Complete debug mode - shows both command and output
config := osexec.NewCommandConfig().WithDebug()

// Show command just
config := osexec.NewCommandConfig().WithDebugMode(osexec.SHOW_COMMAND)

// Show outputs just
config := osexec.NewCommandConfig().WithDebugMode(osexec.SHOW_OUTPUTS)

// Quiet mode - no debug output
config := osexec.NewCommandConfig().WithDebugMode(osexec.QUIET)
```

### Exit Code Handling

Accept specific exit codes as success:

```go
// Accept exit code 1 as success with reason
config := osexec.NewCommandConfig().
	WithExpectExit(1, "DIFFERENCES FOUND")

output, err := config.Exec("diff", "file1.txt", "file2.txt")
// err becomes nil even if diff returns exit code 1

// Accept multiple exit codes
config := osexec.NewCommandConfig().
	WithTakeExits(map[int]string{
		1: "DIFFERENCES FOUND",
		2: "TROUBLE",
	})
```

### Chainable Configuration

Combine multiple configuration options:

```go
config := osexec.NewCommandConfig().
	WithPath("/path/to/project").
	WithEnvs([]string{"ENV=production"}).
	WithBash().
	WithDebugMode(osexec.SHOW_COMMAND).
	WithExpectCode(1)

output, err := config.Exec("command-name", "arg1", "arg2")
```

## API Reference

### Configuration Methods

- **WithEnvs(envs []string)**: Sets custom environment variables
- **WithPath(path string)**: Sets the working path
- **WithShellType(shellType string)**: Sets the shell type (e.g., `bash`)
- **WithShellFlag(shellFlag string)**: Sets the shell flag (e.g., `-c`)
- **WithShell(shellType, shellFlag string)**: Sets both shell type and flag
- **WithBash()**: Configures to use `bash -c`
- **WithZsh()**: Configures to use `zsh -c`
- **WithSh()**: Configures to use `sh -c`
- **WithDebug()**: Enables complete debug mode
- **WithDebugMode(debugMode DebugMode)**: Sets specific debug mode
- **WithExpectExit(exitCode int, reason string)**: Adds expected exit code with reason
- **WithExpectCode(exitCode int)**: Adds expected exit code
- **WithTakeExits(takeExits map[int]string)**: Sets multiple expected exit codes

### Execution Methods

- **Exec(name string, args ...string)**: Executes command and returns output
- **ExecWith(name string, args []string, runWith func(*exec.Cmd))**: Executes with custom command setup
- **StreamExec(name string, args ...string)**: Executes command with pipe handling
- **ExecInPipe(name string, args ...string)**: Executes with stdout/stderr pipe processing

### Debug Modes

- **QUIET**: No debug output
- **DEBUG**: Complete debug mode with command and output
- **SHOW_COMMAND**: Show command just
- **SHOW_OUTPUTS**: Show outputs just

---

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-09-26 07:39:27.188023 +0000 UTC -->

## 📄 License

MIT License. See [LICENSE](LICENSE).

---

## 🤝 Contributing

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Found a mistake?** Open an issue on GitHub with reproduction steps
- 💡 **Have a feature idea?** Create an issue to discuss the suggestion
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![starring](https://starchart.cc/yyle88/osexec.svg?variant=adaptive)](https://starchart.cc/yyle88/osexec)
