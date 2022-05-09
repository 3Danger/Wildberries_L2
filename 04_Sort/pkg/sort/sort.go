package sort

import (
	bytes2 "bytes"
	"reflect"
	"strings"
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

type Sort struct {
	data [][]rune
	keys *map[byte]bool
	k    int
}

type IComparator interface {
	Compare(a, b []rune, k int, delim rune) bool
}

type ICommand interface {
	Exec([]string) []string
}

func NewSortUtil(bytes []byte, keys *map[byte]bool, k int) *Sort {
	bytes = bytes2.Trim(bytes, "\n")
	tmp := strings.Split(string(bytes), "\n")
	runes := make([][]rune, len(tmp))
	for i, v := range tmp {
		runes[i] = []rune(v)
	}
	return &Sort{runes, keys, k}
}

//func (s *Sort) sort(f func(a, b []rune, k int, delim rune) bool) {
func (s *Sort) sort(cmp IComparator) *Sort {
	swapper := reflect.Swapper(s.data)
	for i := 0; i < len(s.data); i++ {
		for j := i + 1; j < len(s.data); j++ {
			if cmp.Compare(s.data[i], s.data[j], s.k, ' ') {
				swapper(i, j)
			}
		}
	}
	return s
}

func (s *Sort) toStrings() (result []string) {
	result = make([]string, len(s.data))
	for i, v := range s.data {
		result[i] = string(v)
	}
	return result
}

func (s *Sort) Run() string {
	command, iCommands := CommandMaker{}.GetCommands(s)
	toStrings := s.sort(command).toStrings()
	for _, v := range iCommands {
		toStrings = v.Exec(toStrings)
	}
	return strings.Join(toStrings, "\n")
}
