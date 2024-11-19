# osexec
Simple utilities to use Golang's `os/exec` package.

## Features

- **ExecInPath**: Execute commands in a specific path.
- **ExecInEnvs**: Execute commands with a custom set of environment variables.
- **ExecXshRun**: Execute commands using a shell with a specific shell type.

## Installation

To install `osexec`, use the following command:

```bash
go get github.com/yyle88/osexec
```

## Functions

### `ExecInPath(path string, name string, args ...string) ([]byte, error)`

Executes a command in the specified path.

#### Parameters:
- `path`: The path where the command should be executed.
- `name`: The name of the command to execute.
- `args`: Arguments to pass to the command.

#### Example:

```go
output, err := osexec.ExecInPath("/path/to/dir", "echo", "Hello, World!")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Output:", string(output))
}
```

### `ExecInEnvs(envs []string, name string, args ...string) ([]byte, error)`

Executes a command with custom environment variables.

#### Parameters:
- `envs`: A list of environment variables to set for the command.
- `name`: The name of the command to execute.
- `args`: Arguments to pass to the command.

#### Example:

```go
envs := []string{"CUSTOM_ENV=1"}
output, err := osexec.ExecInEnvs(envs, "echo", "Custom Env Variable")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Output:", string(output))
}
```

### `ExecXshRun(shellType, shellFlag string, name string, args ...string) ([]byte, error)`

Executes a command using a shell (e.g., `bash`, `sh`) with a specific shell type and flags.

#### Parameters:
- `shellType`: The shell type to use (e.g., `"bash"`, `"sh"`).
- `shellFlag`: The flag to pass to the shell (e.g., `"-c"`).
- `name`: The name of the command to execute.
- `args`: Arguments to pass to the command.

#### Example:

```go
output, err := osexec.ExecXshRun("bash", "-c", "echo", "Shell-based Command")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Output:", string(output))
}
```

---

## `CMC` Structure and Methods

The `CMC` structure provides a flexible way to configure and execute commands. You can set custom environment variables, directories, shell types, and shell flags using a chainable interface.

### `NewCMC() *CMC`

Creates and returns a new `CMC` instance.

### Chainable Methods

- **WithEnvs(envs []string) *CMC**: Sets custom environment variables for the command.
- **WithPath(path string) *CMC**: Sets the working path for the command.
- **WithShellType(shellType string) *CMC**: Sets the shell type (e.g., `bash`).
- **WithShellFlag(shellFlag string) *CMC**: Sets the shell flag (e.g., `-c`).

### `Exec(cmc *CMC, name string, args ...string) ([]byte, error)`

Executes the command based on the configuration in the `CMC` instance.

#### Example:

```go
cmc := osexec.NewCMC().
    WithEnvs([]string{"CUSTOM_ENV=1"}).
    WithPath("/path/to/dir").
    WithShellType("bash").
    WithShellFlag("-c")

output, err := cmc.Exec("echo", "Hello from CMC!")
if err != nil {
    fmt.Println("Error:", err)
} else {
    fmt.Println("Output:", string(output))
}
```

## Error Handling

Each function in this package returns an error if the execution fails. Be sure to handle errors properly in your application.

---

## Contribution

Feel free to fork this repository, make improvements, and submit pull requests. Issues and feature requests are also welcome!

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## Thank You

If you find this package valuable, give it a star on GitHub! Thank you!!!

---
