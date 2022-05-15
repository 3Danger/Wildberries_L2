package builtins

import "microshell/shell/commands/common"

type Ps struct {
	common.Command
}

func (p Ps) Run() (pid uintptr, ok error) {

	return 0, nil
}
