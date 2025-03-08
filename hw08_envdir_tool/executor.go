package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func RunCmd(cmd []string, env Environment) (returnCode int) {
	if len(cmd) == 0 {
		return 1
	}

	for key, val := range env {
		if val.NeedRemove {
			os.Unsetenv(key)
		} else {
			os.Setenv(key, val.Value)
		}
	}

	binPath, err := exec.LookPath(cmd[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: command not found: %s\n", cmd[0])
		return 127
	}

	command := exec.Command(binPath, cmd[1:]...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
	command.Env = os.Environ()

	if err := command.Run(); err != nil {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return exitErr.ExitCode()
		}
		return 1
	}
	return 0
}
