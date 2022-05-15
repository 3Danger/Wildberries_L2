package builtins

import (
	"log"
	"microshell/shell/commands/common"
	"os"
	"syscall"
)

type Echo struct {
	common.Command
}

func (e *Echo) Run() (pid uintptr, ok error) {
	var (
		n   int
		buf []byte
	)
	pid = e.ForkMe()
	if pid == 0 {
		if ok = e.DupAll(); ok != nil {
			log.Fatal()
		}
		buf = make([]byte, 0, 1)
		for n, ok = syscall.Read(e.Reader(), buf); n > 0; {
			if _, ok = syscall.Write(e.Writer(), buf); ok != nil {
				log.Fatal(ok)
			}
		}
		os.Exit(0)
	}
	return pid, nil
}
