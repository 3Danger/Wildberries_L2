package mysort

import "strings"

type Indexes [][2]int

func (idx Indexes) Get(i int) [2]int {
	if i < 0 || i >= len(idx) {
		return [2]int{0, len(idx)}
	}
	return idx[i]
}

type StringerElem struct {
	Data           string
	IndexesOfSplit Indexes
}

func (s *StringerElem) GetByIdx(Idx int) string {
	i := s.IndexesOfSplit.Get(Idx)
	return s.Data[i[0]:]
}

func (s *StringerElem) GetWord(idx int) string {
	indexes := s.IndexesOfSplit.Get(idx)
	return s.Data[indexes[0]:indexes[1]]
}

func getIndexes(data, delim string) (indexes [][2]int) {
	indexes = make([][2]int, 0, 10)
	ab := [2]int{0, 0}
	for i := 0; i < len(data); {
		if strings.HasPrefix(data[i:], delim) {
			for ; i < len(data) && strings.HasPrefix(data[i:], delim); i++ {
			}
			ab[0] = i
		} else {
			for ; i < len(data) && !strings.HasPrefix(data[i:], delim); i++ {
			}
			ab[1] = i
			indexes = append(indexes, ab)
		}
	}
	return indexes
}

func SliceToStringers(data []string, delim string) (result StringerElements) {
	for i := range data {
		indexes := getIndexes(data[i], delim)
		result = append(result, StringerElem{
			Data:           data[i],
			IndexesOfSplit: indexes})
	}
	return result
}

type StringerElements []StringerElem

func (s StringerElements) ToString() string {
	sb := strings.Builder{}
	for i := range s {
		sb.WriteString(s[i].Data + "\n")
	}
	return sb.String()
}
