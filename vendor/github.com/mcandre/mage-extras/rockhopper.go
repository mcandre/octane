package mageextras

import (
	"os"
	"os/exec"
)

// Rockhopper runs rockhopper.
func Rockhopper(args ...string) error {
	cmd := exec.Command("rockhopper")
	cmd.Args = append(cmd.Args, args...)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
