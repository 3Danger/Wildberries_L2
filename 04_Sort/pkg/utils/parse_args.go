package utils

import (
	"flag"
	"strings"
)

func ParseFlags() (*map[byte]bool, int) {
	var n, r, u, M, b, c, h bool
	var k int
	flag.IntVar(&k, "k", 0, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()
	return &map[byte]bool{'n': n, 'r': r, 'u': u, 'M': M, 'b': b, 'c': c, 'h': h}, k - 1
	//return &map[byte]bool{'n': true, 'r': r, 'u': u, 'M': M, 'b': b, 'c': c, 'h': h}, 6
}

func IsNumber(r rune) bool {
	return r >= '0' && r <= '9'
}

func GetIndex(data []rune, delim rune, column int) (index int) {
	var last rune = 'k'
	for i, v := range data {
		if v == delim {
			if v != last {
				column--
			}
		} else if column == 0 {
			return i
		}
		last = v
	}
	return strings.LastIndexFunc(string(data), func(r rune) bool { return r != delim })
}

//func GetIndexOfNumbers(data []rune, delim rune, column int) (index int) {
//	var last rune = 'k'
//	for i, v := range data {
//		if !IsNumber(v) {
//			if IsNumber(last) {
//				column--
//			}
//		} else if column == 0 {
//			return i
//		}
//		last = v
//	}
//	return -1
//return strings.LastIndexFunc(string(data), func(r rune) bool { return r != delim })
//}
