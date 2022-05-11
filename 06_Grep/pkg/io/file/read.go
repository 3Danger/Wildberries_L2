package file

import (
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
