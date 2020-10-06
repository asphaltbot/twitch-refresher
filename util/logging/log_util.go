package logging

import (
	"github.com/fatih/color"
	"os"
)

var (
	InfoLine  = color.New(color.FgCyan).PrintlnFunc()
	ErrorLine = color.New(color.FgRed).PrintlnFunc()
	fatalLine = color.New(color.BgRed, color.FgBlack).PrintlnFunc()
)

func FatalLine(content ...interface{}) {
	fatalLine(content...)
	os.Exit(1)
}
