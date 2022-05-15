package builtins

import (
	"errors"
	"microshell/shell/commands/common"
	"os"
)

//Run() error

type Cd struct {
	common.Command
}

func (c Cd) Run() (ok error) {
	args := c.Args()
	if len(args) != 2 {
		return errors.New("cd: too many arguments")
	}
	if ok := os.Chdir(args[1]); ok != nil {
		return errors.New("cd: " + ok.Error())
	}
	return nil
}
