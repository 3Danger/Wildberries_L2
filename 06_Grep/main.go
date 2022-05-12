package main

import (
	"fmt"
	"grep/pkg/Grep"
	"grep/pkg/io/file"
	"regexp"
)

func main() {
	grep := Grep.NewGrep()
	run := grep.Run()
	fmt.Print(run)
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
