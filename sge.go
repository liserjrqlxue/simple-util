package simple_util

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

var sgeJobId = regexp.MustCompile(`^Your job (\d+) \("\S+"\) has been submitted\n$`)

func SGEsubmit(i int, cmds []string, hjid string, submitArgs []string) string {
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

	return strings.Join(jids, ",")
}

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
