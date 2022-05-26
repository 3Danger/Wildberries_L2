package typesort

import (
	mysort "04_Sort2/src/mysort"
	"sort"
	"strconv"
)

/*
	-n — сортировать по числовому значению
*/

type Numbereded struct {
	K int
}

func (n Numbereded) Do(data []mysort.StringerElem) {
	sort.SliceStable(data, func(i, j int) bool {
		var (
			aok, bok   error
			a, b       int64
			astr, bstr string
		)
		astr = data[i].GetWord(n.K)
		bstr = data[j].GetWord(n.K)
		a, aok = strconv.ParseInt(astr, 10, 32)
		b, bok = strconv.ParseInt(bstr, 10, 32)
		if aok != nil && aok == bok {
			return astr < bstr
		}
		return a < b
		//return data[i].GetWord(n.K) < data[j].GetWord(n.K)
	})
}
