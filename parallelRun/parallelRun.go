package main

import (
	"bufio"
	"flag"
	"github.com/liserjrqlxue/simple-util"
	"io"
	"log"
	"os"
	"strings"
)

var (
	threshold = flag.Int(
		"threshold",
		12,
		"threshold to used",
	)
	list = flag.String(
		"list",
		"",
		"cmdline list to parallel run",
	)
)

func main() {
	flag.Parse()
	if *list == "" {
		flag.Usage()
		log.Print("-list is required")
		os.Exit(0)
	}
	log.Printf("Parallel Run Start:%v", os.Args)
	var line string

	file, err := os.Open(*list)
	simple_util.CheckErr(err)
	defer simple_util.DeferClose(file)

	c := make(chan bool, *threshold)
	reader := bufio.NewReader(file)
	var i = 0
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		line = strings.TrimSuffix(line, "\n")
		i++
		c <- true
		go func(i int, cmd []string) {
			log.Printf("Task[%5d] Start:%v", i, cmd)
			err = simple_util.RunCmd(cmd[0], cmd[1:]...)
			if err != nil {
				log.Printf("Task[%5d] Error:%v", i, err)
			} else {
				log.Printf("Task[%5d] Done", i)
			}
			<-c
		}(i, strings.Split(line, " "))
	}
	if err != io.EOF {
		log.Printf("Error:%v", err)
	}
	for i := 0; i < *threshold; i++ {
		c <- true
	}
	for i := 0; i < *threshold; i++ {
		<-c
	}
	log.Printf("Parallel Run Done")
}
