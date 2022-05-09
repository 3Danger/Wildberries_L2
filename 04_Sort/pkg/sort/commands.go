package sort

import (
	"04_Sort/pkg/utils"
	"reflect"
	"strconv"
	"unicode"
)

/*
	Реализовать поддержку утилитой следующих ключей:

	-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
	-n — сортировать по числовому значению
	-r — сортировать в обратном порядке
	-u — не выводить повторяющиеся строки

	Дополнительно

	Реализовать поддержку утилитой следующих ключей:

	-M — сортировать по названию месяца
	-b — игнорировать хвостовые пробелы
	-c — проверять отсортированы ли данные
	-h — сортировать по числовому значению с учетом суффиксов
*/

type CommandMaker struct{}

func (CommandMaker) GetCommands(sortStruct *Sort) (IComparator, []ICommand) {
	var comparators IComparator
	command := make([]ICommand, 0)

	if ok, _ := (*sortStruct.keys)['n']; ok {
		comparators = NumValueSort{}
	} else {
		comparators = DefaultSort{}
	}
	if ok, _ := (*sortStruct.keys)['r']; ok {
		command = append(command, Reverse{})
	}
	if ok, _ := (*sortStruct.keys)['c']; ok {
		command = append(command, Unique{})
	}
	return comparators, command

}

// NumValueSort -n — сортировать по числовому значению
type NumValueSort struct{}

//drwxrwxr-x 4 csamuro csamuro    4096 May  9 11:55 .
func (NumValueSort) Compare(a, b []rune, k int, delim rune) bool {
	sa, sb := string(a), string(b)
	ai := utils.GetIndex(a, delim, k)
	bi := utils.GetIndex(b, delim, k)
	if ai == -1 || bi == -1 {
		return ai < 0
	}
	if ai >= len(a) || bi >= len(b) {
		return bi >= len(b)
	}
	ssa := string(a[ai])
	ssb := string(b[bi])
	_, _, _, _ = sa, sb, ssa, ssb
	parseInt := func(a []rune) int {
		l := 0
		for l < len(a) && unicode.IsDigit(a[l]) {
			l++
		}
		l64, ok := strconv.ParseInt(string(a[:l]), 10, 32)
		if ok != nil || l == len(a) {
			l64 = -1
		}
		return int(l64)
	}
	massa, massb := parseInt(a[ai:]), parseInt(b[bi:])
	if massa != massb {
		if massa == 0 || massb == 0 {
			return massa != 0
		}
		return massa > massb
	} else {
		//return string(a[ai:]) < string(b[bi:])
		return !utils.StringComparator(a[ai:], b[bi:])
	}
}

// DefaultSort — сортировать по умолчанию
type DefaultSort struct{}

func (DefaultSort) Compare(a, b []rune, k int, delim rune) bool {
	strA, strB := string(a), string(b)
	_, _ = strA, strB
	ai := utils.GetIndex(a, delim, k)
	bi := utils.GetIndex(b, delim, k)
	lena, lenb := len(a), len(b)
	if lena == 0 || lenb == 0 {
		return lenb == 0
	}
	return utils.StringComparator(a[ai:], b[bi:])
	//for ; ai < lena && bi < lenb; ai, bi = ai+1, bi+1 {
	//	if (a[ai] == delim || b[bi] == delim) && a[ai] != b[bi] {
	//		return a[ai] != delim
	//	}
	//	if a[ai] != b[bi] {
	//		return string(a[ai:]) < string(b[bi:])
	//	}
	//}
	//return lena > lenb
}

// Reverse -r — сортировать в обратном порядке
type Reverse struct{}

func (Reverse) Exec(data []string) []string {
	swapper := reflect.Swapper(data)
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		swapper(i, j)
	}
	return data
}

// Unique -u — не выводить повторяющиеся строки
type Unique struct{}

func (Unique) Exec(data []string) []string {
	tmp := make(map[string]struct{}, len(data))
	for _, v := range data {
		tmp[v] = struct{}{}
	}
	data = make([]string, 0, len(tmp))
	for k, _ := range tmp {
		data = append(data, k)
	}
	return data
}
