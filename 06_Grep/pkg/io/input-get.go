package io

import (
	"grep/pkg/Config"
	"io/ioutil"
	"log"
	"strings"
)

func GetData(conf Config.Conf) []string {
	var (
		raw []byte
		ok  error
	)
	if raw, ok = ioutil.ReadAll(conf.Input); ok != nil {
		log.Fatal(ok)
	}
	return strings.Split(string(raw), "\n")
}
