package example2_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/osexec"
)

// TestConfigBasic demonstrates basic CommandConfig usage
// TestConfigBasic 演示 CommandConfig 的基本用法
func TestConfigBasic(t *testing.T) {
	// Create a config with working path using t.TempDir()
	// 使用 t.TempDir() 创建带有工作路径的配置
	tempDIR := t.TempDir()

	config := osexec.NewCommandConfig().WithPath(tempDIR)
	output, err := config.Exec("pwd")
	require.NoError(t, err)
	require.Contains(t, string(output), tempDIR)
	t.Logf("Current path: %s", string(output))
}

// TestConfigWithEnvs demonstrates config with environment variables
// TestConfigWithEnvs 演示带环境变量的配置
func TestConfigWithEnvs(t *testing.T) {
	// Create config with custom environment variables
	// 创建带自定义环境变量的配置
	config := osexec.NewCommandConfig().
		WithEnvs([]string{"MY_VAR=test123"})

	output, err := config.Exec("bash", "-c", "echo $MY_VAR")
	require.NoError(t, err)
	require.Contains(t, string(output), "test123")
	t.Logf("Environment variable: %s", string(output))
}

// TestConfigWithDebug demonstrates config with debug mode
// TestConfigWithDebug 演示带调试模式的配置
func TestConfigWithDebug(t *testing.T) {
	// Enable debug mode to see command execution details
	// 启用调试模式以查看命令执行详情
	config := osexec.NewCommandConfig().
		WithDebug()

	output, err := config.Exec("echo", "Hello, Debug!")
	require.NoError(t, err)
	require.Contains(t, string(output), "Hello, Debug!")
	t.Logf("Output: %s", string(output))
}

// TestConfigChained demonstrates chaining multiple config options
// TestConfigChained 演示链式设置多个配置选项
func TestConfigChained(t *testing.T) {
	// Create temp DIR
	// 创建临时 DIR
	tempDIR := t.TempDir()

	// Chain multiple configuration options
	// 链式设置多个配置选项
	config := osexec.NewCommandConfig().
		WithPath(tempDIR).
		WithEnvs([]string{"DEMO=chained"}).
		WithDebugMode(osexec.SHOW_COMMAND)

	output, err := config.Exec("bash", "-c", "echo $DEMO in $(pwd)")
	require.NoError(t, err)
	require.Contains(t, string(output), "chained")
	t.Logf("Output: %s", string(output))
}
