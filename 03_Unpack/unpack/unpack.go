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

func isBlock(rune []byte, i int) bool {
	return (!isNum(rune[i]) &&
		(i == 0 || rune[i-1] != '\\')) ||
		(i == 0 || rune[i-1] == '\\')
}

func getBlock(rune []byte) (block, int) {
	b := block{rune[0], 1}
	i := 1
	for ; i < len(rune); i++ {
		if !isNum(rune[i]) {
			break
		}
	}
	if i != 1 {
		n, _ := strconv.Atoi(string(rune[1:i]))
		b.count = n
		return b, i - 1
	}
	return b, 0
}

func getBlocks(rune []byte) []block {
	blocks := make([]block, 0)

	if isNum(rune[0]) {
		return nil
	} //			"a4bc2d5e", "aaaabccddddde"
	//			"\\1\0", "\0"
	for i := 0; i < len(rune); i++ {
		if rune[i] == '\\' {
			i++
		}
		if isBlock(rune, i) {
			block, n := getBlock(rune[i:])
			blocks = append(blocks, block)
			i += n
		}
	}
	return blocks
}

func Unpack(rune *string) string {
	if *rune == "" {
		return ""
	}
	blocks := getBlocks([]byte(*rune))
	sb := strings.Builder{}
	for _, v := range blocks {
		sb.WriteString(strings.Repeat(string(v.char), v.count))
	}
	return sb.String()
}
