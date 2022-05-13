package parse

import (
	"fmt"
	"io"
	"log"
	"os"
	"sort"
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

type Segment struct {
	a, b int
}

func (s Segment) GetA() int { return s.a }
func (s Segment) GetB() int { return s.b }

func (s Segment) isInner(oth Segment) bool {
	return s.a <= oth.a && s.b >= oth.b
}

func (s Segment) isTouchesEdge(oth Segment) bool {
	if s.a == INFINITY || s.a <= oth.a {
		if s.b == INFINITY || s.b >= oth.a {
			return true
		}
		return false
	}
	return oth.isTouchesEdge(s)
}

func (s *Segment) setIn(oth Segment) bool {
	if s.isInner(oth) {
		return true
	}
	if s.isTouchesEdge(oth) {
		if s.a != INFINITY && s.a > oth.a {
			s.a = oth.a
		}
		if s.b != INFINITY && s.b < oth.b && oth.b != NOTSETED {
			s.b = oth.b
		}
		return true
	}
	return false
}

type Config struct {
	F    []Segment
	D    byte
	S    bool
	Read io.Reader
}

func NewConfig() *Config {
	var conf = new(Config)
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
	if conf.F == nil {
		fmt.Println("cut: you must specify a list of bytes, characters, or fields")
		os.Exit(0)
	}
	return postScriptF(conf)
}

func postScriptF(config *Config) *Config {
	var res []Segment
	var tmp Segment
	f := config.F
	sort.Slice(f, func(i, j int) bool {
		if f[i].a == -1 {
			return true
		}
		if f[i].a == f[j].a {
			return f[i].b < f[j].b
		}
		return f[i].a < f[j].a
	})
	tmp = f[0]
	for _, v := range f {
		if !tmp.setIn(v) {
			res = append(res, tmp)
			tmp = v
		}
	}
	res = append(res, tmp)
	config.F = res
	return config
}

func GetFile(filePath string) io.Reader {
	open, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return open
}

func parseF(data string) (resultF []Segment) {
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
			resultF = append(resultF, Segment{INFINITY, b})
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
			resultF = append(resultF, Segment{a, b})
		}
	}
	return resultF
}
