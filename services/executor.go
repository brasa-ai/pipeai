package services

import (
	"os"
	"os/exec"

	"github.com/your-handle/pipeai/helpers"
)

// Run executes the generated shell command.
func Run(cmdStr string) error {
	helpers.Log.Debug().Str("cmd", cmdStr).Msg("executing")
	shell := "bash"
	args := []string{"-c"}
	if os.PathSeparator == '\\' { // Windows
		shell = "cmd"
		args = []string{"/c"}
	}
	cmd := exec.Command(shell, append(args, cmdStr)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
