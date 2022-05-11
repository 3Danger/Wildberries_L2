package main

import (
	"fmt"
	"grep/pkg/Config"
	"grep/pkg/io/file"
	"regexp"
)

func main() {
	fmt.Println(Config.GetConfig())
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
