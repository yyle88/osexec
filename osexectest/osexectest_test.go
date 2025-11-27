package osexectest_test

import (
	"testing"

	"github.com/yyle88/osexec/osexectest"
)

func TestSkipIfCommandNotFound(t *testing.T) {
	osexectest.SkipIfCommandNotFound(t, "zsh")
}
