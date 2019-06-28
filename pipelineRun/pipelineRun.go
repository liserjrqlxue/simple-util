package main

import (
	"bufio"
	"flag"
	"github.com/liserjrqlxue/simple-util"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var (
	proj = flag.String(
		"proj",
		"",
		"project for SGE(-P)",
	)
	queue = flag.String(
		"queue",
		"bc_b2c.q",
		"queue for SGE(-q)",
	)
	mode = flag.String(
		"mode",
		"local",
		"run mode:[local|sge]",
	)
	list = flag.String(
		"list",
		"",
		"script list to pipeline run:\nlocal\t:script\targs\nsge\t:scripts\tsubmitArgs",
	)
	cwd = flag.Bool(
		"cwd",
		false,
		"-cwd for SGE",
	)
	threshold = flag.Int(
		"threshold",
		12,
		"threshold for local mode",
	)
)

var (
	sep = regexp.MustCompile(`\s+`)
)

func main() {
	log.Printf("Pipeline Run Start:%v", os.Args)

	flag.Parse()
	if *list == "" {
		flag.Usage()
		log.Print("-list is required")
		os.Exit(0)
	}

	var submitArgs []string
	if *cwd {
		submitArgs = append(submitArgs, "-cwd")
	}
	if *queue != "" {
		submitArgs = append(submitArgs, "-q", *queue)
	}
	if *proj != "" {
		submitArgs = append(submitArgs, "-P", *proj)
	}

	file, err := os.Open(*list)
	simple_util.CheckErr(err)
	defer simple_util.DeferClose(file)

	var line string
	firstChan := make(chan string)
	oldChan := make(chan string)
	oldChan = firstChan
	reader := bufio.NewReader(file)
	var i = 0
	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		i++
		line = strings.TrimSuffix(line, "\n")
		cmds := sep.Split(line, -1)

		newChan := make(chan string)
		switch *mode {
		case "local":
			go LocalRun(i, cmds, oldChan, newChan)
		case "sge":
			go SGEsubmmit(i, cmds, oldChan, newChan, submitArgs)
		default:
			log.Printf("Error Mode:[%s]", *mode)
		}
		oldChan = newChan
	}
	firstChan <- ""
	log.Printf("last job[%s] submitted\n", <-oldChan)
}

var sgeJobId = regexp.MustCompile(`^Your job (\d+) \("\S+"\) has been submitted\n$`)

func submit(script, hjid string, submitArgs []string) (jid string) {
	if hjid != "" {
		submitArgs = append(submitArgs, "-hold_jid", hjid)
	}
	args := append(submitArgs, script)
	c := exec.Command("qsub", args...)
	log.Print("qsub ", strings.Join(args, " "))
	submitLogBytes, err := c.CombinedOutput()
	submitLog := string(submitLogBytes)
	if err != nil {
		log.Fatalf("Error:%v:[%v]", err, submitLog)
	}
	// Your job (\d+) \("script"\) has been submitted
	log.Print(submitLog)
	submitLogs := sgeJobId.FindStringSubmatch(submitLog)
	if len(submitLogs) == 2 {
		jid = string(submitLogs[1])
	} else {
		log.Fatalf("Error: jid parse error:%s->%+v", submitLog, submitLogs)
	}
	return
}

func SGEsubmmit(i int, cmds []string, oldChan <-chan string, newChan chan<- string, submitArgs []string) {
	hjid := <-oldChan
	log.Printf("Task[%5d] Start:%v", i, cmds)

	args := append(submitArgs, cmds[1:]...)
	var jids []string
	for j, cmd := range strings.Split(cmds[0], ",") {
		log.Printf("submit Task[%5d].%d:%s", i, j, cmd)
		jid := submit(cmd, hjid, args)
		if jid != "" {
			jids = append(jids, jid)
		}
	}

	newChan <- strings.Join(jids, ",")
	return
}

func LocalRun(i int, cmds []string, oldChan <-chan string, newChan chan<- string) {
	<-oldChan
	log.Printf("Task[%04d] Start:%v", i, cmds)
	var cmdSlice [][]string
	for _, cmd := range strings.Split(cmds[0], ",") {
		var cmdline []string
		cmdline = append(cmdline, "bash", cmd)
		cmdline = append(cmdline, cmds[1:]...)
		cmdSlice = append(cmdSlice, cmdline)
	}
	ok := simple_util.ParallelRun(cmdSlice, *threshold, strconv.Itoa(i))
	if !ok {
		log.Fatalf("Task[%04d] Run and Stop", i)
	}
	newChan <- ""
	return
}
