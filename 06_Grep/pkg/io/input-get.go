package io

import (
	"grep/pkg/Grep/Config"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func ReadFile(filePath string) []string {
	readFile, ok := ioutil.ReadFile(filePath)
	if ok != nil {
		log.Fatal(ok)
		os.Exit(1)
	}
	split := strings.Split(string(readFile), "\n")
	return split
}

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
