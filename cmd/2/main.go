package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func findLetterIndex(r []rune) int {
	for i, char := range r {
		if unicode.IsLetter(char) {
			return i
		}
	}
	return -1
}

func unpackLetter(r []rune, builder *strings.Builder) {
	if len(r) == 1 {
		builder.WriteRune(r[0])
		return
	}
	num, _ := strconv.Atoi(string(r[1:]))
	for i := 0; i < num; i++ {
		builder.WriteRune(r[0])
	}
}

func splitString(str string) [][]rune {
	var result [][]rune
	runes := []rune(str)

	idx := findLetterIndex(runes)
	if idx == -1 {
		return result
	}
	for {
		nextR := runes[idx+1:]

		nextIdx := findLetterIndex(nextR)

		if nextIdx == -1 {
			result = append(result, runes)
			break
		}
		result = append(result, runes[idx:nextIdx+idx+1])
		runes = runes[nextIdx+idx+1:]
		idx = 0
	}
	return result
}

func unpackString(str string) string {
	runes := splitString(str)
	sb := strings.Builder{}
	for _, ele := range runes {
		unpackLetter(ele, &sb)
	}
	return sb.String()
}

func main() {
	fmt.Println(unpackString("a4bc2d5e"))
	fmt.Println(unpackString("abcd"))
	fmt.Println(unpackString("45"))
	fmt.Println(unpackString(""))
	fmt.Println(unpackString("45a"))
}
