package commands

import (
	"log"
	"sync"
	"syscall"
)

type ICommand interface {
	Run(*sync.WaitGroup) error
	Writer() int
	SetWriter(int)
	Reader() int
	SetReader(int)
	Pid() uintptr
}

func ExecuteAll(executable []ICommand) {
	var (
		ok    error
		pid   uintptr
		pids  []uintptr
		group sync.WaitGroup
	)

	for _, e := range executable {
		group.Add(1)
		if ok = e.Run(&group); ok != nil {
			log.Fatal(ok)
		} else if pid = e.Pid(); pid > 0 {
			pids = append(pids, pid)
			if _, ok = syscall.Wait4(int(pid), nil, 0, nil); ok != nil {
				log.Fatal(ok)
			}
		}
		group.Wait()
	}
	//for _, pid = range pids {
	//	if _, ok = syscall.Wait4(int(pid), nil, 0, nil); ok != nil {
	//		log.Fatal(ok)
	//	}
	//fmt.Println(pid, "is done")
	//}
	//group.Wait()
}
