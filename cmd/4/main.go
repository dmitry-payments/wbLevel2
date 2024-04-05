package main

import (
	"fmt"
	"sort"
)

func findAnagramSets(words []string) map[string][]string {
	anagramMap := make(map[string][]string)

	sort.Strings(words)

	for _, word := range words {
		sortedWord := sortString(word)
		anagramMap[sortedWord] = append(anagramMap[sortedWord], word)
	}

	resultMap := make(map[string][]string)

	for _, anagrams := range anagramMap {
		if len(anagrams) > 1 {
			resultMap[anagrams[0]] = anagrams
		}
	}

	return resultMap
}

func sortString(s string) string {
	chars := []rune(s)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	return string(chars)
}

func main() {
	words := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "столяр"}
	anagramSets := findAnagramSets(words)
	fmt.Println(anagramSets)
}
