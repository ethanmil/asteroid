//go:build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Run() {
	cmd := exec.Command("go", "run", ".")

	// Pass all streams via terminal
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Errorf("failed to run: %s", err.Error())
	}
}
