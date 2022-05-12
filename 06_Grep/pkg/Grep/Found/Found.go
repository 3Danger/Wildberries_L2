package Found

import (
	color "github.com/fatih/color"
	"grep/pkg/Config"
	"sort"
)

type PointIndex struct {
	startString, endString int
	indexes                []int
}

func (p PointIndex) isSelected(i int) bool {
	for _, v := range p.indexes {
		if v == i {
			return true
		}
	}
	return false
}

func (p PointIndex) GetSize() int {
	return p.endString - p.startString
}

func NewPointIndex(index, startStr, endStr int) *PointIndex {
	return &PointIndex{
		startString: index - startStr,
		endString:   index + endStr,
		indexes:     []int{index},
	}
}

func MixPoints(pointIndexes ...*PointIndex) []*PointIndex {
	if pointIndexes == nil || len(pointIndexes) == 0 {
		return nil
	}
	var result = make([]*PointIndex, 0, len(pointIndexes))
	sort.SliceIsSorted(pointIndexes, func(i, j int) bool {
		return pointIndexes[i].indexes[0] > pointIndexes[j].indexes[0]
	})
	var i int
	result = append(result, pointIndexes[0])
	for _, v := range pointIndexes[1:] {
		if result[i].endString >= v.startString {
			result[i].endString = v.endString
			result[i].indexes = append(result[i].indexes, v.indexes...)
		} else {
			i++
			result = append(result, v)
		}
	}
	return result
}

type Found struct {
	Conf Config.Conf
	data []string

	indexes *PointIndex
}

func NewFound(conf Config.Conf, data []string, indexes *PointIndex) *Found {
	return &Found{conf, data, indexes}
}

func (f Found) GetData() []string {
	var result []string
	var row string
	var start = f.indexes.startString
	var end = f.indexes.endString

	for ; start <= end; start++ {
		row = prepareResult(&f, start, f.data[start])
		result = append(result, row)
	}
	return result
}

func prepareResult(f *Found, numRow int, row string) string {
	var prefix string
	isSelected := f.indexes.isSelected(numRow)
	if f.Conf.Keyn {
		prefix = color.GreenString("%d", numRow)
		if isSelected {
			prefix += color.BlueString("%c", ':')
		} else {
			prefix += color.BlueString("%c", '-')
		}
	}
	return prefix + row + "\n"

}
