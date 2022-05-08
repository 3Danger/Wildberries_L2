package main

import (
	"flag"
	"fmt"
	"github.com/3Danger/Wildberries_L2/02_BasicTask/etime"
	"os"
)

func main() {
	var host string
	flag.StringVar(&host, "h", "time.apple.com", "set host address")
	flag.Parse()
	tm, err := etime.GetExactTime(host)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println(tm.String())
}
