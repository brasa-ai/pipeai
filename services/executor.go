package services

import (
	"os"
	"os/exec"

	"github.com/your-handle/pipeai/helpers"
)

// Run executes the generated shell command.
func Run(cmdStr string) error {
	helpers.Log.Debug().Str("cmd", cmdStr).Msg("executing")
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
