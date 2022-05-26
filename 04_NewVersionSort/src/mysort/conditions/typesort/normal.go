package typesort

import (
	"04_Sort2/src/mysort"
	"sort"
)

type Normal struct{ K int }

func (n Normal) Do(data []mysort.StringerElem) {
	sort.SliceStable(data, func(i, j int) bool {
		return data[i].GetWord(n.K) < data[j].GetWord(n.K)
	})
}
