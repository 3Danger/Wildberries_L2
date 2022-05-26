package utils

import (
	"io"
	"io/ioutil"
	"strings"
)

func ReadToString(r io.Reader) []string {
	if b, ok := ioutil.ReadAll(r); ok != nil {
		panic(ok)
	} else {
		return strings.Split(string(b), "\n")
	}
}
