[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/osexec/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/osexec/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/osexec)](https://pkg.go.dev/github.com/yyle88/osexec)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/osexec/master.svg)](https://coveralls.io/github/yyle88/osexec?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/osexec.svg)](https://github.com/yyle88/osexec/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/osexec)](https://goreportcard.com/report/github.com/yyle88/osexec)

# osexec

Simple utilities to use Golang's `os/exec` package.

## CHINESE README

[中文说明](README.zh.md)

## Features

- **Custom Execution Configurations**: Execute commands with customizable environment variables, working paths, and shell options.

## Installation

```bash  
go get github.com/yyle88/osexec  
```  

## `CommandConfig` Structure and Methods

`CommandConfig` structure provides a flexible way to configure and execute commands. You can set custom environment variables, directories, shell types, and debug options using a chainable interface.

### `NewCommandConfig() *CommandConfig`

Creates and returns a new `CommandConfig` instance.

#### Example:

```go  
config := osexec.NewCommandConfig()
```  

### Chainable Methods

- **WithEnvs(envs []string) *CommandConfig**: Sets custom environment variables.
- **WithPath(path string) *CommandConfig**: Sets the working path.
- **WithShellType(shellType string) *CommandConfig**: Sets the shell type (e.g., `bash`).
- **WithShellFlag(shellFlag string) *CommandConfig**: Sets the shell flag (e.g., `-c`).
- **WithShell(shellType, shellFlag string) *CommandConfig**: Sets shell type and flag.
- **WithBash() *CommandConfig**: Configures the command to use `bash -c`.
- **WithZsh() *CommandConfig**: Configures the command to use `zsh -c`.
- **WithSh() *CommandConfig**: Configures the command to use `sh -c`.
- **WithDebugMode(debugMode bool) *CommandConfig**: Enables / disables debug mode.

#### Example:

```go
package main

import (
	"fmt"
	"github.com/yyle88/osexec"
)

func main() {
	// Create a new CommandConfig instance and set the working directory and debug mode
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

## License

MIT License. See [LICENSE](LICENSE).

---

## Support

Welcome to contribute to this project by submitting pull requests and reporting issues.

If you find this package valuable, give me some stars on GitHub! Thank you!!!

**Thank you for your support!**

**Happy Coding with `osexec`!** 🎉

Give me stars. Thank you!!!

## GitHub Stars

[![starring](https://starchart.cc/yyle88/osexec.svg?variant=adaptive)](https://starchart.cc/yyle88/osexec)
