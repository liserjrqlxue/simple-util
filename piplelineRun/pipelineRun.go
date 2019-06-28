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
		"sge",
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
		switch *mode {
		case "local":
			err = simple_util.RunCmd("bash", cmd[0])
		case "sge":
			newChan := make(chan string)
			go SGEsubmmit(cmd[0], oldChan, newChan, appendArgs, cmd[1:])
			oldChan = newChan
		default:
			log.Printf("Error Mode:[%s]", *mode)
		}
	}
	firstChan <- ""
	<-oldChan
}

var jobSubmitted = regexp.MustCompile(`^Your job (\d+) \("\S+"\) has been submiited$`)

func SGEsubmmit(script string, oldChan <-chan string, newChan chan<- string, submitArgs, scriptArgs []string) (error error) {
	hjid := <-oldChan
	if hjid != "" {
		submitArgs = append(submitArgs, "-hold_jid", hjid)
	}
	args := append(submitArgs, script)
	args = append(args, scriptArgs...)
	c := exec.Command("qsub", args...)
	submitLog, err := c.CombinedOutput()
	if err != nil {
		log.Fatalf("Error:%v %v", submitLog, err)
		newChan <- ""
		return
	}
	// Your job (\d+) \("script"\) has been submitted
	submitLogs := jobSubmitted.FindSubmatch(submitLog)
	log.Printf("%+v", submitLog)
	log.Printf("%+v", submitLogs)
	jid := string(submitLogs[1])
	newChan <- jid
	return
}
