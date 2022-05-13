package main

import (
	"fmt"
	"grep/pkg/Grep"
)

func main() {
	grep := Grep.NewGrep()
	run := grep.Run()
	fmt.Print(run)
}
