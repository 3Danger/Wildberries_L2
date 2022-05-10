package anagram

import (
	ut "FindAnagram/pkg/utils"
)

func GetCombinations(word []rune) [][]rune {
	var (
		size    = len(word)
		setSize = size * size
		words   = make([][]rune, 0, setSize)
		arri    = make([]int, size)
	)

	for j := range arri {
		arri[j] = j
	}

	for w := 0; w < setSize; w++ {
		words = append(words, make([]rune, size))
		for j := range arri {
			words[w][j] = word[arri[j]]
		}
		nextComb(arri, len(arri), true)
		if ut.HasCycleCompeted(arri) {
			break
		}
	}
	return words
}

func nextComb(arri []int, size int, rec bool) {
	lastIndex := len(arri) - 1
	arri[lastIndex]++
	if arri[lastIndex] == size {
		arri[lastIndex] = 0
		if lastIndex != 0 {
			nextComb(arri[:lastIndex], size, false)
		}
	}
	if rec && ut.HasEqual(arri) {
		nextComb(arri, size, true)
	}
}

func hasCycle(arri []int) bool {
	size := len(arri)
	if arri[size-1] == size-2 {
		for i := range arri[:size-1] {
			if arri[i] == i-1 {
				return true
			}
			return false
		}
	}
	return false
}

//func anagram(m *map[string]struct{}, word []rune) []string {
//
//	return nil
//}
//
//func FindAnagram(raw []string) *map[string][]string {
//	tmp := toMap(raw)
//}
