package Address

import (
	"fmt"
	"strings"
)

type Address struct {
	country, region, city, post, home, userName, phone string
}

func (a Address) PrintInfo() {
	fmt.Println(strings.Join([]string{a.country, a.region, a.city, a.post, a.home, a.userName, a.phone}, "\n"))
}
