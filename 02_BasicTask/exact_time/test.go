package exact_time

import (
	"fmt"
	"testing"
)

func TestGetExactTime(t *testing.T) {
	time, err := GetExactTime("time.apple.com")
	if err != nil {
		return
	}
	fmt.Println(time)
}
