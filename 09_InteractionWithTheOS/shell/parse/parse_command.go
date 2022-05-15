package parse

import (
	"errors"
	"log"
	"microshell/shell/commands"
	"microshell/shell/commands/builtins"
	"microshell/shell/commands/common"
	"os"
	"strings"
	"syscall"
)

func customSplit(data, delim, ignore string) (result []string) {
	var ignored bool
	var pnt int
	var i int

	data = strings.Trim(data, delim)
	for ; i < len(data); i++ {
		if strings.ContainsRune(ignore, rune(data[i])) && (i == 0 || data[i-1] != '\\') {
			ignored = !ignored
		}
		if !ignored {
			if strings.ContainsRune(delim, rune(data[i])) && (i == 0 || data[i-1] != '\\') {
				result = append(result, data[pnt:i])
				for i+1 < len(data) && strings.ContainsRune(delim, rune(data[i+1])) {
					i++
				}
				pnt = i + 1
			}
		}
	}
	if i != pnt {
		result = append(result, data[pnt:i])
	}
	return result
}

func CreateCommands(input string, paths []string) (cms []commands.ICommand, ok error) {
	const ignore = "\"'"
	var pipex = make([]int, 2)
	var std = make([]int, 2)
	var out, in int

	if ok = syscall.Pipe(pipex); ok != nil {
		return nil, ok
	}
	if std[0], ok = syscall.Dup(0); ok != nil {
		log.Fatal(ok)
	}
	if std[1], ok = syscall.Dup(1); ok != nil {
		log.Fatal(ok)
	}

	out = std[1]
	in = pipex[1]
	groups := customSplit(input, ";", ignore)
	for _, group := range groups {
		pipeSplit := customSplit(group, "|", ignore)
		for _, cmdline := range pipeSplit {
			args := customSplit(cmdline, " ", ignore)
			cmd, err := createCommand(args, paths, in, out)
			if err != nil {
				return nil, err
			}
			cms = append(cms, cmd)

			out = pipex[0]
			if ok = syscall.Pipe(pipex); ok != nil {
				return nil, ok
			}
			in = pipex[1]
		}
		in = std[1]
		cms[len(cms)-1].SetWriter(in)
	}
	return cms, nil
}

func checkFile(ut string) (res string, notOk error) {
	stat, notOk := os.Stat(ut)
	if notOk != nil {
		return "", notOk
	} else if stat.IsDir() {
		return "", errors.New(ut + " is directory, can't execute")
	} else if stat.Mode()&0100 == 0 {
		return "", errors.New(ut + " isn't executable, pls make: \n$> chmod +x " + ut)
	}
	return ut, nil
}

func createCommand(args, paths []string, writer, reader int) (res commands.ICommand, notOk error) {
	switch args[0] {
	case "cd":
		return &builtins.Cd{Command: *common.NewCommand(args, os.Environ(), writer, reader)}, nil
	case "pwd":
		return &builtins.Pwd{Command: *common.NewCommand(args, os.Environ(), writer, reader)}, nil
	case "echo":
		return &builtins.Echo{Command: *common.NewCommand(args, os.Environ(), writer, reader)}, nil
	case "kill":
		return &builtins.Kill{Command: *common.NewCommand(args, os.Environ(), writer, reader)}, nil
	case "ps":
		return &builtins.Ps{Command: *common.NewCommand(args, os.Environ(), writer, reader)}, nil
	}
	for _, v := range paths {
		if _, notOk = checkFile(v + "/" + args[0]); notOk == nil {
			args[0] = v + "/" + args[0]
			return common.NewCommand(args, os.Environ(), writer, reader), nil
		}
	}
	return nil, errors.New(args[0] + ": command not found")
}
