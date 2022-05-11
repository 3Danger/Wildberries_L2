package main

import (
	"FindAnagram/pkg/anagram"
	ut "FindAnagram/pkg/utils"
	"flag"
	"fmt"
)

//GetInput обрабатываем флаги
func GetInput() (filePath string) {
	flag.StringVar(&filePath, "f", "files/small_file.txt", "file name")
	flag.Parse()
	return
}

func main() {
	// Получаем данные
	filePath := GetInput()
	data := ut.ReadFromFile(filePath)

	// <<< Функция поиска всех множеств анаграмм по словарю.>>>
	result := anagram.FindAnagram(data)

	//Выводим результат
	for k, v := range *result {
		fmt.Printf("Group: %s\n", k)
		for _, v := range v {
			fmt.Printf("\t: %s\n", v)
		}
	}
}
