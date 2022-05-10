package main

import (
	"04_Sort/pkg/sort"
	"04_Sort/pkg/utils"
	"fmt"
	"io/ioutil"
	"log"
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

	-M - о̶р̶т̶и̶р̶о̶в̶а̶т̶ь̶ ̶п̶о̶ ̶н̶а̶з̶в̶а̶н̶и̶ю̶ ̶м̶е̶с̶я̶ц̶а̶
	-b - и̶г̶н̶о̶р̶и̶р̶о̶в̶а̶т̶ь̶ ̶х̶в̶о̶с̶т̶о̶в̶ы̶е̶ ̶п̶р̶о̶б̶е̶л̶ы̶
	-c - р̶о̶в̶е̶р̶я̶т̶ь̶ ̶о̶т̶с̶о̶р̶т̶и̶р̶о̶в̶а̶н̶ы̶ ̶л̶и̶ ̶д̶а̶н̶н̶ы̶е̶
	-h - с̶о̶р̶т̶и̶р̶о̶в̶а̶т̶ь̶ ̶п̶о̶ ̶ч̶и̶с̶л̶о̶в̶о̶м̶у̶ ̶з̶н̶а̶ч̶е̶н̶и̶ю̶ ̶с̶ ̶у̶ч̶е̶т̶о̶м̶ ̶с̶у̶ф̶ф̶и̶к̶с̶о̶в̶
*/

func main() {
	flags, k := utils.ParseFlags()
	bytes, ok := ioutil.ReadAll(os.Stdin)
	if ok != nil {
		log.Fatal(ok)
		os.Exit(1)
	}
	fmt.Println(sort.NewSortUtil(bytes, flags, k).Run())
}
