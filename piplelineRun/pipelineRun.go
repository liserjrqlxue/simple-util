package main

import (
	"bufio"
	"flag"
	"github.com/liserjrqlxue/simple-util"
	"log"
	"os"
	"os/exec"
	"regexp"
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
		"",
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
		"script list to pipeline run:\nlocal:script args\nsge:script submitArgs",
	)
	cwd = flag.Bool(
		"cwd",
		false,
		"-cwd for SGE",
	)
)

func main() {
	log.Printf("Pipeline Run Start:%v", os.Args)

	flag.Parse()

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
		cmds := strings.Split(line, "\t")

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

func SGEsubmmit(i int, cmds []string, oldChan <-chan string, newChan chan<- string, submitArgs []string) {
	hjid := <-oldChan
	log.Printf("Task[%5d] Start:%v", i, cmds)
	if hjid != "" {
		submitArgs = append(submitArgs, "-hold_jid", hjid)
	}
	args := append(submitArgs, cmds[1:]...)
	args = append(args, cmds[0])
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
		jid := string(submitLogs[1])
		newChan <- jid
	} else {
		log.Fatalf("Error: jid parse error:%s->%+v", submitLog, submitLogs)
	}
	return
}

func LocalRun(i int, cmds []string, oldChan <-chan string, newChan chan<- string) {
	<-oldChan
	log.Printf("Task[%5d] Start:%v", i, cmds)
	log.Print("bash ", strings.Join(cmds, " "))
	err := simple_util.RunCmd("bash", cmds...)
	if err != nil {
		log.Fatalf("Error:%v", err)
	}
	newChan <- ""
	return
}
