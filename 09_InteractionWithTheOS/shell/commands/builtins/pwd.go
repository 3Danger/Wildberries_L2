package builtins

import (
	"microshell/shell/commands/common"
)

type Pwd struct {
	common.Command
}

func (p Pwd) Run() (ok error) {

	return nil
}
