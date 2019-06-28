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
		"script list to pipeline run",
	)
)

func main() {
	log.Printf("Pipeline Run Start:%v", os.Args)

	flag.Parse()

	var appendArgs = []string{"-cwd"}
	if *queue != "" {
		appendArgs = append(appendArgs, "-q", *queue)
	}
	if *proj != "" {
		appendArgs = append(appendArgs, "-P", *proj)
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
		cmd := strings.Split(line, "\t")
		log.Printf("Task[%5d] Start:%v", i, cmd)
		newChan := make(chan string)
		switch *mode {
		case "local":
			go LocalRun(cmd[0], oldChan, newChan, cmd[1:])
		case "sge":
			go SGEsubmmit(cmd[0], oldChan, newChan, appendArgs, cmd[1:])
		default:
			log.Printf("Error Mode:[%s]", *mode)
		}
		oldChan = newChan
	}
	firstChan <- ""
	log.Printf("last job[%s] submitted\n", <-oldChan)
}

var sgeJobId = regexp.MustCompile(`^Your job (\d+) \("\S+"\) has been submitted\n$`)

func SGEsubmmit(script string, oldChan <-chan string, newChan chan<- string, submitArgs, scriptArgs []string) {
	hjid := <-oldChan
	if hjid != "" {
		submitArgs = append(submitArgs, "-hold_jid", hjid)
	}
	args := append(submitArgs, script)
	args = append(args, scriptArgs...)
	c := exec.Command("qsub", args...)
	log.Print("qsub ", strings.Join(args, " "))
	submitLogBytes, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("Error:%v %v", submitLogBytes, err)
	}
	// Your job (\d+) \("script"\) has been submitted
	submitLog := string(submitLogBytes)
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

func LocalRun(script string, oldChan <-chan string, newChan chan<- string, scriptArgs []string) {
	<-oldChan
	var args = []string{script}
	args = append(args, scriptArgs...)
	log.Print("bash ", strings.Join(args, " "))
	err := simple_util.RunCmd("bash", args...)
	if err != nil {
		log.Fatalf("Error:%v", err)
	}
	newChan <- ""
	return
}
