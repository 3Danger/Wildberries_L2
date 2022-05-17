package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

/*
Утилита telnet

Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port
go-telnet mysite.ru 8080
go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту
	(ip или доменное имя + порт) по протоколу TCP.
	После подключения STDIN программы должен записываться в сокет,
	а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу
	(через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
	Если сокет закрывается со стороны сервера, программа должна также завершаться.
	При подключении к несуществующему сервер, программа должна завершаться через timeout
*/

type Config struct {
	TimeOut    time.Duration
	Host, Port *string
}

func GetArgs() (conf Config) {
	var (
		tmreg, numreg     *regexp.Regexp
		tmIndex, numIndex []int
		host, port, tmp   string
		ok                error
		tm                time.Duration
	)

	numreg, ok = regexp.Compile("\\d+")
	tmreg, ok = regexp.Compile("(ms|m|s)$")

	const timeoutFlag = "-timeout="
	for i := 1; i < len(os.Args); i++ {
		if n := strings.Index(os.Args[i], timeoutFlag); n > 0 {
			if tmIndex = tmreg.FindIndex([]byte(os.Args[i])); tmIndex == nil {
				fmt.Println("wrong argument near " + timeoutFlag)
				os.Exit(0)
			}
			num, _ := strconv.Atoi(os.Args[i][tmIndex[0]:tmIndex[1]])
			if numIndex = numreg.FindIndex([]byte(os.Args[i])); numIndex == nil {
				tm = time.Second * time.Duration(num)
			} else {
				switch os.Args[i][numIndex[0]:numIndex[1]] {
				case "ms":
					tm = time.Millisecond * time.Duration(num)
				case "m":
					tm = time.Minute * time.Duration(num)
				default:
					tm = time.Second * time.Duration(num)
				}
			}
		}
	}
}

func main() {

}
