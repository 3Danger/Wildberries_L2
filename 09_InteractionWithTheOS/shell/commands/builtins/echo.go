package builtins

import (
	"fmt"
	"log"
	"microshell/shell/commands/common"
	"syscall"
)

type Echo struct {
	common.Command
}

func (e *Echo) Run() (ok error) {
	var (
		n   int
		buf = make([]byte, 1)
	)
	fmt.Println("start", e.Args()[0])
	//pid := e.ForkMe()
	//if pid == 0 {
	//	if ok = e.DupAll(); ok != nil {
	//		log.Fatal()
	//	}
	//if ok = syscall.Fchmod()
	//if ok = syscall.Fchmod(e.Reader(), syscall.O_NONBLOCK); ok != nil {
	//	log.Fatal(ok)
	//}
	//if ok = syscall.Fchmod(e.Writer(), syscall.O_NONBLOCK); ok != nil {
	//	log.Fatal(ok)
	//}
	if ok = syscall.SetNonblock(e.Reader(), true); ok != nil {
		log.Fatal(ok)
	}
	n, ok = syscall.Read(e.Reader(), buf)
	for n > 0 && ok == nil {
		if _, ok = syscall.Write(e.Writer(), buf); ok != nil {
			log.Fatal(ok)
		}
		n, ok = syscall.Read(e.Reader(), buf)
	}
	if !isFlagN(e.Args()) {
		syscall.Write(e.Writer(), []byte{'\n'})
	}
	//os.Exit(0)
	//}
	//e.SetFd(pid)
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
