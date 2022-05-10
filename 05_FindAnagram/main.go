package main

import (
	an "FindAnagram/pkg/anagram"
	"fmt"
	"strings"
)

func main() {
	combinations := an.GetCombinations([]rune("hel"))
	data := make([]string, len(combinations))
	for i := range data {
		data[i] = string(combinations[i])
	}
	fmt.Println(strings.Join(data, "\n"))
	fmt.Println(len(data))

	//	123
	//131	-
	//	132
	//133	-
	//211	-
	//212	-
	//	213
	//221	-
	//222	-
	//223	-
	//	231
	//232	-
	//233	-
	//311	-
	//	312
	//313	-
	//	321
	//322	-
	//323	-
	//331	-
	//332	-
	//333	--
	//111	--
	//112	-
	//113	-
	//121	-
	//122	-
}
