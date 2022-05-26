package typesort

import (
	"04_Sort2/src/mysort"
	"regexp"
	"sort"
	"strconv"
)

/*
	-h — сортировать по числовому значению с учетом суффиксов
*/

type Humanity struct {
	K int
}

func (h Humanity) Do(data []mysort.StringerElem) {
	sort.SliceStable(data, func(i, j int) bool {
		var (
			a, b       int64
			astr, bstr string
		)
		astr = data[i].GetWord(h.K)
		bstr = data[j].GetWord(h.K)
		a = calcMass(astr)
		b = calcMass(bstr)
		return a < b
	})
}

func calcMass(data string) (num int64) {
	var ok error

	if num, ok = strconv.ParseInt(data, 10, 32); ok == nil {
		return num
	}
	mass := []*reg{
		NewReg("pb", 1<<60),
		NewReg("tb", 1<<50),
		NewReg("gb", 1<<40),
		NewReg("mb", 1<<30),
		NewReg("kb", 1<<20),
		NewReg("k", 1<<20),
		NewReg("g", 1<<40),
		NewReg("m", 1<<30),
		NewReg("k", 1<<20),
	}
	for _, v := range mass {
		if v.Key.MatchString(data) {
			idx := v.Key.FindIndex([]byte(data))
			num, ok = strconv.ParseInt(data[:idx[0]], 10, 32)
			if ok != nil {
				break
			}
			return num * v.Val
		}
	}
	return -1
}

type reg struct {
	Key *regexp.Regexp
	Val int64
}

func NewReg(r string, val int64) *reg {
	rg, _ := regexp.Compile("(?i)" + r + "$")
	return &reg{rg, val}
}
