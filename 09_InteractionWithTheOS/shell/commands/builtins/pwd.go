package builtins

import (
	"microshell/shell/commands/common"
)

type Pwd struct {
	common.Command
}

func (p Pwd) Run() (pid uintptr, ok error) {

	return 0, nil
}
