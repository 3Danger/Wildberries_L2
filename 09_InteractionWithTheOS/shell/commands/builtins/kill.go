package builtins

import (
	"errors"
	"log"
	"microshell/shell/commands/common"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

type Kill struct {
	common.Command
}

func (k *Kill) Run(group *sync.WaitGroup) (ok error) {
	var (
		pid    int
		PsList string
		tmp    []byte
	)
	args := k.Args()
	if tmp, ok = exec.Command("ps", "-a").Output(); ok != nil {
		log.Fatal(ok)
	}
	PsList = string(tmp)
	_ = PsList
	for _, killingApp := range args[1:] {
		appList := strings.Split(PsList, "\n")
		for _, strokeApp := range appList {
			if strings.HasSuffix(strokeApp, killingApp) {
				strokeApp = strings.TrimLeft(strokeApp, " \t")
				sp := strings.IndexByte(strokeApp, byte(' '))
				if pid, ok = strconv.Atoi(strokeApp[:sp]); ok != nil {
					log.Fatal(ok)
				} else if ok = syscall.Kill(pid, syscall.SIGKILL); ok != nil {
					return errors.New("kill: " + ok.Error())
				}
			}
		}
	}
	k.CloseFds()
	group.Done()
	return nil
}
