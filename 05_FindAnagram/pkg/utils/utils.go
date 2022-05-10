package utils

import "strings"

//func CombCalc(a int) int {
//	b := a
//	for i := 0; i < b; i++ {
//		a *= b
//	}
//	return a
//}

func toMap(raw []string) map[string]struct{} {
	res := make(map[string]struct{}, len(raw))
	for _, v := range raw {
		if len(v) < 1 {
			continue
		}
		res[strings.ToLower(v)] = struct{}{}
	}
	return res
}

func HasCycleCompeted(arri []int) bool {
	counter := 0
	for i, v := range arri {
		if i == v {
			counter++
		}
	}
	return counter == len(arri)
}

func HasEqual(arri []int) bool {
	for i := range arri {
		for j := range arri {
			if i == j {
				continue
			} else if arri[i] == arri[j] {
				return true
			}
		}
	}
	return false
}
