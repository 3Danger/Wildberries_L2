package cut

import (
	"cut/pkg/parse"
	"io/ioutil"
	"log"
	"strings"
)

/*
	Реализовать утилиту аналог консольной команды cut (man cut).
	Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

	Реализовать поддержку утилитой следующих ключей:
	-f - "fields" - выбрать поля (колонки)
	-d - "delimiter" - использовать другой разделитель
	-s - "separated" - только строки с разделителем
*/

type Cut struct {
	data []string
	conf parse.Config
}

func NewCut() Cut {
	config := parse.GenerateConfig()
	bytes, ok := ioutil.ReadAll(config.Read)
	if ok != nil {
		log.Fatal(ok)
	}
	return Cut{strings.Split(string(bytes), "\n"), config}
}

func (c Cut) GetIdx(data []byte, seg [2]int) (a, b int) {
	makeId := func(data []byte, seg, def int) int {
		for i, v := range data {
			if seg == 1 {
				return i
			}
			if v == c.conf.D {
				seg--
			}
		}
		return def
	}
	a = makeId(data, seg[0], 0)
	b = makeId(data, seg[1], len(data))
	return a, b
}

func (c Cut) getBytes(data []byte, segment [2]int) []byte {
	a, b := 0, 0
	if segment[0] == parse.INFINITY {
		a, b = c.GetIdx(data, segment)
		return data[:b]
	} else if segment[1] == parse.INFINITY {
		a, b = c.GetIdx(data, segment)
		return data[a:]
	} else if segment[1] == parse.NOTSETED {
		a, b = c.GetIdx(data, segment)
		return data[a:b]
	} else {
		a, b = c.GetIdx(data, segment)
		return data[a:b]
	}
}

func (c Cut) hasDelim(data []byte) bool {
	for _, v := range data {
		if c.conf.D == v {
			return true
		}
	}
	return false
}

func (c Cut) getResultWithDelim() string {
	var tmp = make([]string, 0, len(c.data))
	var sb strings.Builder
	for i := range c.data {
		if !c.hasDelim([]byte(c.data[i])) {
			continue
		}
		for _, seg := range c.conf.F {
			sb.WriteString(string(c.getBytes([]byte(c.data[i]), seg)))
		}
		tmp = append(tmp, sb.String())
		sb.Reset()
	}
	return strings.Join(tmp, "\n")
}

func (c Cut) getResult() string {
	var tmp = make([]string, 0, len(c.data))
	var sb strings.Builder
	for i := range c.data {
		for _, seg := range c.conf.F {
			sb.WriteString(string(c.getBytes([]byte(c.data[i]), seg)))
		}
		tmp = append(tmp, sb.String())
		sb.Reset()
	}
	return strings.Join(tmp, "\n")
}

func (c Cut) GetResult() string {
	if c.conf.S {
		return c.getResultWithDelim()
	}
	return c.getResult()
}
