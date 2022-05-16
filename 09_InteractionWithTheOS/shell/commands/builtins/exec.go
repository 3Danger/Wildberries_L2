package builtins

import (
	"log"
	"microshell/shell/commands/common"
	"os"
	"sync"
	"syscall"
)

type Exec struct {
	common.Command
}

func (e *Exec) Run(group *sync.WaitGroup) (ok error) {
	e.CloseFds()
	_ = group
	if len(e.Args()) > 1 {
		if ok = syscall.Exec(e.Args()[1], e.Args()[1:], os.Environ()); ok != nil {
			log.Fatal(ok)
		}
	}
	return nil
}
