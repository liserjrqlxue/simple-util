package simple_util

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCmd(name string, args ...string) {
	c := exec.Command(name, args...)
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	err := c.Run()
	fmt.Println(err)
}
