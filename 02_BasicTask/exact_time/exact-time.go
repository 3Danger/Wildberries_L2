package exact_time

import ntp "github.com/beevik/ntp"

func GetExactTime(host string) (string, error) {
	time, ok := ntp.Time(host)
	if ok != nil {
		return "", ok
	}
	return time.String(), nil
}
