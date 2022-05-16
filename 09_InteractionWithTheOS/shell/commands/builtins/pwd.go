package builtins

import (
	"microshell/shell/commands/common"
	"sync"
)

type Pwd struct {
	common.Command
}

func (p Pwd) Run(group *sync.WaitGroup) (ok error) {

	group.Done()
	return nil
}
