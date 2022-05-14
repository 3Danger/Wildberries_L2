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
	conf *parse.Config
}

func NewCut(conf *parse.Config, data []string) Cut {
	if data == nil {
		bytes, ok := ioutil.ReadAll(conf.Read)
		if ok != nil {
			log.Fatal(ok)
		}
		return Cut{strings.Split(string(bytes), "\n"), conf}
	}
	return Cut{data, conf}
}

func (c *Cut) SetData(data []string) {
	c.data = data
}

func (c Cut) GetPoints(data []byte) []int {
	var x = []int{0}
	for i, v := range data {
		if v == c.conf.D {
			x = append(x, i+1)
		}
	}
	x = append(x, len(data))
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (c Cut) getBytes(data []byte, seg [][2]int) string {
	_, _ = data, seg
	var res []byte
	points := c.GetPoints(data)
	lengthPoints := len(points)

	//TODO извлечение данных по индексам
	for _, v := range seg {
		v[1] = min(v[1], lengthPoints-1)
		if v[0] == parse.TIRE {
			res = append(res, data[:points[v[1]]]...)
		} else if v[1] == parse.TIRE {
			res = append(res, data[points[v[0]-1]:]...)
		} else if v[1] == parse.NOTHING {
			if v[0] >= lengthPoints {
				continue
			}
			res = append(res, data[points[v[0]-1]:points[v[0]]]...)
		} else {
			if v[0] >= lengthPoints {
				continue
			}
			res = append(res, data[points[v[0]-1]:points[v[1]]]...)
		}
	}
	return strings.Trim(string(res), string(c.conf.D))
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
	for i := range c.data {
		if c.hasDelim([]byte(c.data[i])) {
			tmp = append(tmp, string(c.getBytes([]byte(c.data[i]), c.conf.F)))
		}
	}
	return strings.Join(tmp, "\n")
}

func (c Cut) getResult() string {
	var tmp = make([]string, 0, len(c.data))
	for i := range c.data {
		if c.hasDelim([]byte(c.data[i])) {
			tmp = append(tmp, c.getBytes([]byte(c.data[i]), c.conf.F))
		} else {
			tmp = append(tmp, c.data[i])
		}
	}
	return strings.Join(tmp, "\n")
}

func (c Cut) GetResult() string {
	if c.conf.S {
		return c.getResultWithDelim()
	}
	return c.getResult()
}
