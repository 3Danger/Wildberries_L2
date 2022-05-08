package exact_time

import (
	"fmt"
	"testing"
)

func TestGetExactTime(t *testing.T) {
	tm, err := GetExactTime("time.apple.com")
	if err != nil {
		return
	}
	fmt.Println(tm)
}
