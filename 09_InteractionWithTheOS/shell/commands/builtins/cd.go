package builtins

import (
	"errors"
	"microshell/shell/commands/common"
	"os"
	"sync"
)

//Run() error

type Cd struct {
	common.Command
}

func (c Cd) Run(group *sync.WaitGroup) (ok error) {
	args := c.Args()
	c.CloseFds()
	if len(args) != 2 {
		return errors.New("cd: too many arguments")
	}
	if ok := os.Chdir(args[1]); ok != nil {
		return errors.New("cd: " + ok.Error())
	}
	group.Done()
	return nil
}
