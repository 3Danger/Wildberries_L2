package typesort

import (
	"04_Sort2/src/mysort"
	"regexp"
	"sort"
)

/*
	-M — сортировать по названию месяца
*/

var monthes = []*regexp.Regexp{
	makeMonth("jan"),
	makeMonth("feb"),
	makeMonth("mar"),
	makeMonth("apr"),
	makeMonth("may"),
	makeMonth("jun"),
	makeMonth("jul"),
	makeMonth("aug"),
	makeMonth("sep"),
	makeMonth("oct"),
	makeMonth("nov"),
	makeMonth("dec"),
}

func makeMonth(short string) *regexp.Regexp {
	m, _ := regexp.Compile("(?i)^(" + short + ")")
	return m
}

type Month struct {
	K int
}

func (m Month) Do(data []mysort.StringerElem) {
	sort.SliceStable(data, func(i, j int) bool {
		var (
			a, b       int64
			astr, bstr string
		)
		astr = data[i].GetWord(m.K)
		bstr = data[j].GetWord(m.K)
		a = getMassOfMonth(astr)
		b = getMassOfMonth(bstr)
		return a < b
	})
}

func getMassOfMonth(data string) int64 {
	for i := 0; i < 12; i++ {
		if monthes[i].MatchString(data) {
			return int64(i)
		}
	}
	return -1
}
