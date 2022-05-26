package postprocess

import (
	mysort "04_Sort2/src/mysort"
	"reflect"
)

/*
	-r — сортировать в обратном порядке
*/

type Reverse struct{ K int }

func (r *Reverse) Do(data []mysort.StringerElem) {
	swap := reflect.Swapper(data)
	end := len(data) - 1
	for i := 0; i < end; i, end = i+1, end-1 {
		swap(i, end)
	}
}
