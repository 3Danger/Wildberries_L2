package main

import (
	"flag"
	"fmt"
	"github.com/3Danger/Wildberries_L2/02_BasicTask/exact_time"
	"os"
)

func main() {
	var host string
	flag.StringVar(&host, "h", "time.apple.com", "set host address")
	flag.Parse()
	tm, err := exact_time.GetExactTime(host)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(tm.String())
}
