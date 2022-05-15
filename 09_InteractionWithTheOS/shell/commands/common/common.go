package common

import (
	"log"
	"os"
	"syscall"
)

// (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

type Command struct {
	args   []string
	env    []string
	writer int
	reader int
}

func (c *Command) Args() []string {
	return c.args
}

func (c *Command) Writer() int {
	return c.writer
}

func (c *Command) SetWriter(writer int) {
	c.writer = writer
}

func (c *Command) Reader() int {
	return c.reader
}

func (c *Command) SetReader(reader int) {
	c.reader = reader
}

func NewCommand(args, env []string, writer, reader int) *Command {
	return &Command{args, env, writer, reader}
}

func (c Command) DupAll() (ok error) {
	if ok = syscall.Dup2(c.writer, 0); ok != nil {
		return ok
	}
	if ok = syscall.Dup2(c.reader, 1); ok != nil {
		return ok
	}
	return nil
}

func (c Command) ForkMe() (pid uintptr) {
	pid, _, _ = syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	return pid
}

func (c Command) Run() (pid uintptr, ok error) {
	pid = c.ForkMe()
	if pid == 0 {
		if ok = c.DupAll(); ok != nil {
			log.Fatal(ok)
		}
		if ok = syscall.Exec(c.args[0], c.args, c.env); ok != nil {
			log.Fatal(ok)
		}
		os.Exit(1)
	}
	return pid, nil
}
