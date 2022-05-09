package utils

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func kValidate(k int) {
	if k <= 0 {
		if k == 0 {
			fmt.Fprintln(os.Stderr, "sort: field number is zero: invalid field specification ‘0’")
		} else {
			fmt.Fprintf(os.Stderr, "sort: invalid number at field start: invalid count at start of ‘%d’\n", k)
		}
		os.Exit(0)
	}
}

func ParseFlags() (*map[byte]bool, int) {
	var n, r, u, M, b, c, h bool
	var k int
	flag.IntVar(&k, "k", 1, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()
	kValidate(k)
	return &map[byte]bool{'n': n, 'r': r, 'u': u, 'M': M, 'b': b, 'c': c, 'h': h}, k - 1
	//return &map[byte]bool{'n': true, 'r': r, 'u': u, 'M': M, 'b': b, 'c': c, 'h': h}, 6
}

func XOR(a, b bool) bool {
	return (a && !b) || (!a && b)
}

func StringComparator(a, b []rune) bool {
	al, bl := len(a), len(b)
	ai, bi := 0, 0

	for a[ai] == '-' && ai < al-1 {
		ai++
	}
	for b[bi] == '-' && bi < bl-1 {
		bi++
	}

	for ; ai < al && bi < bl; ai, bi = ai+1, bi+1 {
		if a[ai] == b[bi] {
			continue
		}
		for (unicode.IsSpace(a[ai]) || a[ai] == '-') && ai < al-1 {
			ai++
		}
		for (unicode.IsSpace(b[bi]) || b[bi] == '-') && bi < bl-1 {
			bi++
		}
		if XOR(ai == al, bi == bl) {
			return ai == al
		}
		//if XOR(unicode.IsNumber(a[ai]), unicode.IsNumber(b[bi])) {
		//	return !unicode.IsDigit(a[ai])
		//}
		if unicode.IsNumber(a[ai]) && unicode.IsNumber(b[bi]) {
			return a[ai] > b[bi]
		}
		//if XOR(unicode.Is(unicode.Cyrillic, a[ai]), unicode.Is(unicode.Cyrillic, b[bi])) {
		//	return !unicode.Is(unicode.Cyrillic, a[ai])
		//}
		//
		//if XOR(unicode.ToLower(a[ai]), unicode.ToLower(b[bi])) {
		//	return !unicode.IsUpper(a[ai])
		//}
		//if XOR(unicode.IsSpace(a[ai]), unicode.IsSpace(b[bi])) {
		//	return !unicode.IsSpace(a[ai])
		//}

		//if XOR(unicode.IsGraphic(a[ai]), unicode.IsGraphic(b[bi])) {
		//	return !unicode.IsGraphic(a[ai])
		//}
		//if XOR(unicode.IsLetter(a[ai]), unicode.IsLetter(b[bi])) {
		//	return unicode.IsLetter(a[ai])
		//}
		return a[ai] > b[bi]
	}
	return al > bl
}

func GetIndex(data []rune, delim rune, column int) (index int) {
	var last = 'k'
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

//total 2000
//-r- 1 csamuro csamuro      24 May  1 07:20 go.mod
//-r- 1 csamuro csamuro       0 May  3 05:11 Русское_Название_Файла
//-r- 1 csamuro csamuro       Русское_Название_Файла
//-r-x 1 csamuro csamuro 2017788 May  4 09:30 main
//dr-x 2 csamuro csamuro    4096 May  2 07:49 .idea
