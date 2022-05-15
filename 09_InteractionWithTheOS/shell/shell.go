package shell

import (
	"errors"
	"os"
	"strings"
)

type Shell struct {
	paths []string
	pwd   string
}

func (s Shell) Paths() []string {
	return s.paths
}

func NewShell() *Shell {
	shell := new(Shell)
	//shell.env = initEnv()
	env := os.Getenv("PATH")
	if env != "" {
		shell.paths = strings.Split(env, ":")
	}
	return shell
}

//func initEnv() map[string]string {
//	environ := os.Environ()
//	env := make(map[string]string, len(environ))
//	for _, v := range environ {
//		split := strings.Split(v, "=")
//		switch len(split) {
//		case 1:
//			env[split[0]] = ""
//		case 2:
//			env[split[0]] = split[1]
//		}
//	}
//	return env
//}

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

func (s Shell) checkAllPath(ut string) (res string, notOk error) {
	for _, v := range s.paths {
		if _, notOk = checkFile(v + "/" + ut); notOk == nil {
			return v + "/" + ut, nil
		}
	}
	return "", errors.New(ut + ": command not found")
}

func (s Shell) whereUtil(ut string) (res string, ok error) {
	if strings.HasPrefix(ut, "./") {
		return checkFile(ut[2:])
	}
	return s.checkAllPath(ut)
}
