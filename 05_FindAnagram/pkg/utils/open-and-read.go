package utils

import (
	"bufio"
	"log"
	"os"
)

// ReadFromFile читает файл и возвращает результат в виде строчного слйса
func ReadFromFile(filePath string) []string {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	result := make([]string, 0, stat.Size()/8)
	buffer := bufio.NewScanner(file)
	for buffer.Scan() {
		result = append(result, buffer.Text())
	}

	return result
}
