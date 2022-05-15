package main

import (
	"fmt"
	"microshell/input_processing"
	"os"
	"syscall"
	"time"
)

/*
Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качестве аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*


Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

func main() {
	p := make([]int, 2)
	syscall.Pipe(p)
	syscall.Write(p[1], []byte("Hello"))
	b := make([]byte, 20)
	syscall.Read(p[0], b)
	fmt.Println(string(b))

	input_processing.ReadLine()
	//pid, _, _ := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	//if pid == 0 {
	//	err := syscall.Exec("/usr/bin/cat", []string{"", "-e"}, os.Environ())
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//	return
	//}
	//syscall.Wait4(int(pid), nil, 0, nil)
	//
	fmt.Println("after")
	//fmt.Println(exec.Command("cat").Output())

	//fmt.Println("BEFORE exec")
	//defer fmt.Println("AFTER exec")
	//output, err := exec.Command("/usr/bin/ls", "-la").Output()
	//if err != nil {
	//	return
	//}
	//fmt.Println(string(output))

}

func main2() {
	fmt.Println("BEFORE exec")
	defer fmt.Println("AFTER exec")
	//syscall.Exec("/usr/bin/ls", []string{"", "-l", "-a"}, os.Environ())

	ret, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
	if err != 0 {
		os.Exit(2)
	}
	if ret == 0 {
		time.Sleep(time.Second * 2)
		err := syscall.Exec("/usr/bin/ls", []string{"", "-l", "-a"}, os.Environ())
		if err != nil {
			return
		}
		os.Exit(0)
	}
	fmt.Println("bef")
	syscall.Wait4(int(ret), nil, 0, nil)
	fmt.Println("aft")

	//pid, err := syscall.ForkExec("/usr/bin/cat", os.Args, nil)
	//if err != nil {
	//	return
	//}
	//fmt.Println(pid)
	//if err = syscall.Kill(pid, syscall.SIGKILL); err != nil {
	//	log.Fatal(err)
	//}

}

/*
Отсутствие в стандартном пакете syscall чистого fork меня не остановило и даже не вызвало никаких подозрений. Я просто сделал примерно так (упрощено):

ret, _, err := syscall.Syscall(syscall.SYS_FORK, 0, 0, 0)
if err != 0 {
	os.Exit(2)
}
if ret > 0 {
	// родительский процесс
	os.Exit(0)
}
*/
