package osexec_test

import (
	"testing"

	"github.com/yyle88/osexec"
	"github.com/yyle88/runpath"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/sure"
	"github.com/yyle88/sure/sure_cls_gen"
	"github.com/yyle88/syntaxgo/syntaxgo_ast"
	"github.com/yyle88/syntaxgo/syntaxgo_reflect"
)

// TestGen generates Must/Soft/Omit wrapper methods for CommandConfig
// TestGen 为 CommandConfig 生成 Must/Soft/Omit 包装方法
func TestGen(t *testing.T) {
	options := sure_cls_gen.NewClassGenOptions(runpath.PARENT.Path()).
		WithNewClassNameParts("88").
		WithNamingPatternType(sure_cls_gen.STYLE_SUFFIX_CAMELCASE_TYPE).
		MoreErrorHandlingModes(sure.MUST, sure.SOFT, sure.OMIT)

	config := &sure_cls_gen.ClassGenConfig{
		ClassGenOptions: options,
		PackageName:     syntaxgo_reflect.GetPkgName(osexec.CommandConfig{}),
		ImportOptions:   syntaxgo_ast.NewPackageImportOptions(),
		OutputPath:      runtestpath.SrcPath(t),
	}

	sure_cls_gen.GenerateClasses(config, osexec.CommandConfig{})
}

// TestCommandConfig88Must_Exec tests the Must wrapper execution
// TestCommandConfig88Must_Exec 测试 Must 包装器的执行
func TestCommandConfig88Must_Exec(t *testing.T) {
	output := osexec.NewCommandConfig().
		WithEnvs([]string{"A=1", "B=2"}).
		WithDebugMode(osexec.DEBUG).
		WithBash().
		Must().
		Exec("echo", "$A", "$B")
	t.Log(string(output))
}
