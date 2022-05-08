package exact_time

import ntp "github.com/beevik/ntp"

func GetExactTime() (string, error) {
	time, ok := ntp.Time("time.apple.com")
	if ok != nil {
		return "", ok
	}
	return time.String(), nil
}
