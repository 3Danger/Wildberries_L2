package exact_time

import (
	"github.com/beevik/ntp"
	"time"
)

func GetExactTime(host string) (*time.Time, error) {
	tm, ok := ntp.Time(host)
	if ok != nil {
		return nil, ok
	}
	return &tm, nil
}
