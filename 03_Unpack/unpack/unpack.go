package unpack

import (
	"strconv"
	"strings"
)

type block struct {
	char  byte
	count int
}

func isNum(r byte) bool {
	return r >= '0' && r <= '9'
}

func isBlock(data []byte, i int) bool {
	return (!isNum(data[i]) &&
		(i == 0 || data[i-1] != '\\')) ||
		(i == 0 || data[i-1] == '\\')
}

func getBlock(data []byte) (block, int) {
	b := block{data[0], 1}
	i := 1
	for ; i < len(data); i++ {
		if !isNum(data[i]) {
			break
		}
	}
	if i != 1 {
		n, _ := strconv.Atoi(string(data[1:i]))
		b.count = n
		return b, i - 1
	}
	return b, 0
}

func getBlocks(data []byte) []block {
	blocks := make([]block, 0)

	if isNum(data[0]) {
		return nil
	} //			"a4bc2d5e", "aaaabccddddde"
	//			"\\1\0", "\0"
	for i := 0; i < len(data); i++ {
		if data[i] == '\\' {
			i++
		}
		if isBlock(data, i) {
			block, n := getBlock(data[i:])
			blocks = append(blocks, block)
			i += n
		}
	}
	return blocks
}

func Unpack(data *string) string {
	if *data == "" {
		return ""
	}
	blocks := getBlocks([]byte(*data))
	sb := strings.Builder{}
	for _, v := range blocks {
		sb.WriteString(strings.Repeat(string(v.char), v.count))
	}
	return sb.String()
}
