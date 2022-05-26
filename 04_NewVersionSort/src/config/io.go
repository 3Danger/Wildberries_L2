package config

import (
	"io"
	"os"
	"strings"
)

func inData() (in io.Reader) {
	var ok error
	in = os.Stdin
	for i := 1; i < len(os.Args); i++ {
		os.Args[i] = strings.TrimSpace(os.Args[i])
		if os.Args[i][0] != '-' && !strings.Contains(os.Args[i-1], "-k") {
			if in, ok = os.Open(os.Args[i]); ok != nil {
				panic(ok)
			}
			return in
		}
	}
	return in
}
