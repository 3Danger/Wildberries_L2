package main

import (
	"flag"
	"fmt"
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
