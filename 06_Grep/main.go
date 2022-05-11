package main

import (
	"fmt"
	"grep/pkg"
	"grep/pkg/io/file"
	"regexp"
)

func main() {
	grep := pkg.NewGrep()
	run := grep.Run()
	fmt.Println(run)
}

func test() {
	readFile := file.ReadFile("test.txt")
	compile, err := regexp.Compile("f.*")
	if err != nil {
		return
	}
	for _, v := range readFile {
		var index [][]int
		if index = compile.FindAllIndex([]byte(v), -1); index == nil {
			continue
		}
		fmt.Println(index)
	}
	_ = readFile
}
