package builtins

import (
	"log"
	"microshell/shell/commands/common"
	"sync"
	"syscall"
)

type Echo struct {
	common.Command
}

func (e *Echo) Run(group *sync.WaitGroup) (ok error) {
	var (
		wasFlagN bool
		x        = 1
	)
	wasFlagN = isFlagN(e.Args())
	if wasFlagN {
		x++
	}
	if _, ok = syscall.Write(e.Writer(), []byte(e.Args()[x])); ok != nil {
		log.Fatal(ok)
	}
	if !wasFlagN {
		if _, ok = syscall.Write(e.Writer(), []byte{'\n'}); ok != nil {
			log.Fatal(ok)
		}
	}
	e.CloseFds()
	group.Done()
	return nil
}

func isFlagN(args []string) bool {
	var t, b bool
	for _, v := range args {
		for _, r := range v {
			if r == '-' {
				t = true
			} else if r == 'n' && t {
				b = true
			} else if !t || !b {
				t, b = false, false
			}
		}
	}
	return t && b
}
