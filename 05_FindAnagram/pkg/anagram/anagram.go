package anagram

import (
	"reflect"
	"strings"
)

// Sort сортирует буквы по алфавиту для дальнейшего использования в качестве ключа
func Sort[T string | rune](data []T) []T {
	size := len(data)
	swapper := reflect.Swapper(data)
	for j := 0; j < size; j++ {
		for k := j + 1; k < size; k++ {
			if data[j] > data[k] {
				swapper(j, k)
			}
		}
	}
	return data
}

// Приводим все в единое состояние
func getSiglegram(raw []string) *map[string]string {
	abcmap := make(map[string]string, len(raw))
	for i := range raw {
		// Все слова должны быть приведены к нижнему регистру.
		raw[i] = strings.ToLower(raw[i])
		// Так же сортируем буквы по алфавиту для дальнейшего использования в качестве ключа
		abcmap[raw[i]] = string(Sort([]rune(raw[i])))
	}
	return &abcmap
}

func makeAnagram(rawsigle *map[string]string) *map[string][]string {
	mediaresult := make(map[string][]string)
	result := make(map[string][]string)
	for k, v := range *rawsigle {
		mediaresult[v] = append(mediaresult[v], k)
	}
	for _, v := range mediaresult {
		// Множества из одного элемента не должны попасть в результат.
		if len(v) > 1 {
			//Sort(v)
			//Массив должен быть отсортирован по возрастанию.
			result[v[0]] = v[1:]
		}
	}
	return &result
}

// FindAnagram находит все анаграммы и возвращает результат
func FindAnagram(raw []string) *map[string][]string {
	rawsingle := getSiglegram(raw)
	anag := makeAnagram(rawsingle)
	return anag
}
