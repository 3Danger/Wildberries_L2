package utils

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func TryItUntilTimeOut(until, sleep time.Duration, f func() error) (ok error) {
	now := time.Now()
	before := now.Add(until)
	for ok = f(); ok != nil && !time.Now().After(before); {
		if ok = f(); ok != nil {
			time.Sleep(sleep)
		}
	}
	return ok
}

var msgMut = sync.Mutex{}

func ServiceMessage(msg string) {
	msgMut.Lock()
	tire := strings.Repeat("-", (40-len(msg))>>1)
	bias := ""
	if len(msg)%2 == 1 {
		bias = "-"
	}
	fmt.Println("\r<" + tire + " " + msg + " " + bias + tire + ">")
	msgMut.Unlock()
}
