package simple_util

import (
	"os"
	"os/exec"
)

func RunCmd(name string, args ...string) error {
	c := exec.Command(name, args...)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	err := c.Run()
	return err
}
