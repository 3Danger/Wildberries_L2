package exact_time

import "github.com/beevik/ntp"

func GetExactTime(host string) (string, error) {
	time, ok := ntp.Time(host)
	if ok != nil {
		return "", ok
	}
	return time.String(), nil
}
