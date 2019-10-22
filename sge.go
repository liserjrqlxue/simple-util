package simple_util

import (
	"log"
	"os/exec"
	"regexp"
	"strings"
)

var sgeJobId = regexp.MustCompile(`^Your job (\d+) \("\S+"\) has been submitted\n$`)

func SGEsubmit(cmds []string, hjid string, submitArgs []string) string {
	return Submit(cmds[0], hjid, submitArgs, cmds[1:])
}

func Submit(script, hjid string, submitArgs, args []string) (jid string) {
	if hjid != "" {
		submitArgs = append(submitArgs, "-hold_jid", hjid)
	}
	cmds := append(submitArgs, script)
	cmds = append(cmds, args...)
	c := exec.Command("qsub", cmds...)
	log.Print("qsub [", strings.Join(cmds, "] ["))
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
