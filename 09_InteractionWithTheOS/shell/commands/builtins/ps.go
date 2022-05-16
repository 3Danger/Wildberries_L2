package builtins

import (
	"microshell/shell/commands/common"
	"sync"
)

type Ps struct {
	common.Command
}

func (p Ps) Run(group *sync.WaitGroup) (ok error) {

	group.Done()
	return nil
}
