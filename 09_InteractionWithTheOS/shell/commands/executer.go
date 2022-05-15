package commands

import (
	"fmt"
	"log"
	"syscall"
	"time"
)

type ICommand interface {
	Run() (uintptr, error)
	Writer() int
	SetWriter(int)
	Reader() int
	SetReader(int)
}

func ExecuteAll(executable []ICommand) {
	var (
		ok   error
		pid  uintptr
		pids []uintptr
	)
	for _, e := range executable {
		time.Sleep(time.Second)
		if pid, ok = e.Run(); ok != nil {
			log.Fatal(ok)
		} else if pid > 0 {
			pids = append(pids, pid)
		}
	}
	for _, pid = range pids {
		if _, ok = syscall.Wait4(int(pid), nil, 0, nil); ok != nil {
			log.Fatal(ok)
		}
		fmt.Println(pid, "is done")
	}
}
