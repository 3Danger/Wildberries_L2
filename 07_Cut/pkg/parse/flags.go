package parse

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
	Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

	Реализовать поддержку утилитой следующих ключей:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем
*/

const INFINITY = -1
const NOTSETED = -2

type Config struct {
	F    [][2]int
	D    byte
	S    bool
	Read io.Reader
}

func GenerateConfig() Config {
	var conf Config
	conf.D = '\t'
	conf.Read = os.Stdin
	length := len(os.Args)
	for i := 1; i < length; i++ {
		if os.Args[i][0] == '-' && len(os.Args[i]) == 2 {
			switch os.Args[i][1] {
			case 'f':
				if i < length-1 {
					i++
					conf.F = parseF(os.Args[i])
				}
			case 'd':
				if i < length-1 {
					i++
					conf.D = os.Args[i][0]
				}
			case 's':
				conf.S = true
			}
		} else {
			conf.Read = GetFile(os.Args[i])
		}
	}
	return conf
}

func GetFile(filePath string) io.Reader {
	open, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return open
}

func parseF(data string) (resultF [][2]int) {
	var ok error
	a, b := 0, 0
	res := strings.Split(data, ",")
	for _, v := range res {
		v = strings.Trim(v, " \t")
		if len(v) == 0 {
			log.Fatal("cut: fields are numbered from 1")
		}
		two := strings.Split(v, "-")
		if len(two) == 0 {
			log.Fatal("cut: invalid range with no endpoint: -")
		}
		if v[0] == '-' {
			if b, ok = strconv.Atoi(two[0]); ok != nil {
				log.Fatal(ok)
			}
			b++
			resultF = append(resultF, [2]int{INFINITY, b})
		} else {
			if a, ok = strconv.Atoi(two[0]); ok != nil {
				log.Fatal(ok)
			}
			if len(two) > 1 {
				if b, ok = strconv.Atoi(two[1]); ok != nil {
					log.Fatal(ok)
				}
				b++
			} else if v[len(v)-1] == '-' {
				b = INFINITY
			} else {
				b = NOTSETED
			}
			if b != INFINITY && b != NOTSETED && a > b {
				log.Fatal("cut: invalid decreasing range")
			}
			resultF = append(resultF, [2]int{a, b})
		}
	}
	return resultF
}
