package builtins

import (
	"errors"
	"microshell/shell/commands/common"
	"strconv"
	"syscall"
)

type Kill struct {
	common.Command
}

func (k *Kill) Run() (uintptr, error) {
	var (
		pid int
		ok  error
	)
	args := k.Args()
	length := len(args)
	if length > 1 {
		for _, v := range args {
			if pid, ok = strconv.Atoi(v); ok != nil {
				return 0, errors.New("kill: " + ok.Error())
			} else {
				if ok = syscall.Kill(pid, syscall.SIGKILL); ok != nil {
					return 0, errors.New("kill: " + ok.Error())
				}
			}
		}
	}
	return 0, nil
}
