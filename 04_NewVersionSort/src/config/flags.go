package config

import (
	"flag"
)

/*
	Утилита mysort
	Отсортировать строки в файле по аналогии с консольной утилитой mysort (man mysort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

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

type Keys map[byte]interface{}

//TODO delete after debug
func DebugNewKeys() *Keys {
	return newKeys()
}

//TODO delete after debug
func (k *Keys) DebugSetKeys(keys ...byte) {
	for _, v := range keys {
		(*k)[v] = v
	}
}

func newKeys() *Keys {
	var (
		k                   int
		n, r, u, M, b, c, h bool
	)

	flag.IntVar(&k, "k", 000000, "указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)")
	flag.BoolVar(&n, "n", false, "сортировать по числовому значению")
	flag.BoolVar(&r, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(&u, "u", false, "не выводить повторяющиеся строки")
	flag.BoolVar(&M, "M", false, "сортировать по названию месяца")
	flag.BoolVar(&b, "b", false, "игнорировать хвостовые пробелы")
	flag.BoolVar(&c, "c", false, "проверять отсортированы ли данные")
	flag.BoolVar(&h, "h", false, "сортировать по числовому значению с учетом суффиксов")
	flag.Parse()
	return &Keys{'k': k, 'n': n, 'r': r, 'u': u, 'M': M, 'b': b, 'c': c, 'h': h}
}
