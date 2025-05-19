//go:build !windows
// +build !windows

package main

import (
	"github.com/apernet/hysteria/app/v2/cmd"
)

func run() {
	cmd.Execute()
}
