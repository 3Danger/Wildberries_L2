package main

import (
	"fmt"
	"github.com/3Danger/Wildberries_L2/02_BasicTask/exact_time"
)

func main() {
	time, err := exact_time.GetExactTime("time.apple.com")
	if err != nil {
		return
	}
	fmt.Println(time)
}
