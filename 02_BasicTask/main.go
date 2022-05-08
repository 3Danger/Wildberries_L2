package main

import (
	"fmt"
	"github.com/3Danger/Wildberries_L2/02_BasicTask/exact_time"
	"log"
)

func main() {
	time, err := exact_time.GetExactTime("time.apple.com")
	if err != nil {
		log.Panic(err)
		return
	}
	fmt.Println(time)
}
