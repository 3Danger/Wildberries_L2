package main

import (
	"04_Sort/pkg/sort"
	"04_Sort/pkg/utils"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {
	flags, k := utils.ParseFlags()
	all, err := ioutil.ReadAll(os.Stdin)
	//open, err := os.Open("./txt2")
	if err != nil {
		return
	}
	//all, err := ioutil.ReadAll(open)
	//if err != nil {
	//	log.Panic(err)
	//}
	//fmt.Println("BEFORE\n", string(bytes.Runes(all)))
	//fmt.Println("AFTER\n", sort.NewSortUtil(all, flags, k).Run())
	fmt.Println(sort.NewSortUtil(all, flags, k).Run())
}
