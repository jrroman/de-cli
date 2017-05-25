package cmd

import (
	"fmt"
	"os/exec"
)

type RunCmd struct {}

func (c *RunCmd) Run(cmdname string, arg ...string) ([]byte, error) {
	output, err := exec.Command(cmdname, arg...).CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("Error executing command: %s\nError: %s\n", cmdname, err)
	}
	return output, nil
}