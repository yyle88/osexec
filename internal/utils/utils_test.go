package utils_test

import (
	"testing"

	"github.com/yyle88/osexec/internal/utils"
)

func TestShowCommand(t *testing.T) {
	utils.ShowCommand("ls")
}

func TestShowMessage(t *testing.T) {
	utils.ShowMessage("ok")
}
