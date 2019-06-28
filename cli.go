package simple_util

import (
	"log"
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

func RunTask(c <-chan bool, cmd []string, task string, index int, ok *bool) {
	log.Printf("Task[%s.%04d] Start:%v", task, index, cmd)
	err := RunCmd(cmd[0], cmd[1:]...)
	if err != nil {
		log.Printf("Task[%s.%04d] Error:%v", task, index, err)
		*ok = false
	} else {
		log.Printf("Task[%s.%04d] Done", task, index)
	}
	<-c
}

func ParallelRun(cmds [][]string, threshold int, tag string) (ok bool) {
	ok = true
	c := make(chan bool, threshold)
	for i, cmd := range cmds {
		c <- true
		go RunTask(c, cmd, tag, i, &ok)
	}
	for i := 0; i < threshold; i++ {
		c <- true
	}
	for i := 0; i < threshold; i++ {
		<-c
	}
	if ok {
		log.Print("Parallel Run Done With Error")
	} else {
		log.Print("Parallel Run Done")
	}
	return
}
