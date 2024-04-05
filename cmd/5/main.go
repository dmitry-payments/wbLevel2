package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Флаги командной строки
	after := flag.Int("A", 0, "печатать +N строк после совпадения")
	before := flag.Int("B", 0, "печатать +N строк до совпадения")
	context := flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	count := flag.Bool("c", false, "количество строк")
	ignoreCase := flag.Bool("i", false, "игнорировать регистр")
	invert := flag.Bool("v", false, "вместо совпадения, исключать")
	fixed := flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum := flag.Bool("n", false, "номер строки")
	flag.Parse()

	if (*after > 0 || *before > 0) && *context > 0 {
		fmt.Println("Необходимо указать либо after и before либо context")
		os.Exit(1)
	}

	if *context > 0 {
		*after = *context
		*before = *context
	}

	// Аргументы командной строки
	pattern := flag.Arg(0)

	// Проверка наличия аргумента
	if pattern == "" {
		fmt.Println("Необходимо указать паттерн для фильтрации.")
		os.Exit(1)
	}

	if *ignoreCase {
		pattern = strings.ToLower(pattern)
	}

	// Сканер для чтения ввода
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
		os.Exit(1)
	}

	matched := false

	var matchedLinesCount int
	printLine := func(s string, idx int) {
		if *lineNum {
			fmt.Printf("%d: ", idx)
		}
		fmt.Println(s)
	}

	for idx, line := range lines {
		l := line
		if *ignoreCase {
			l = strings.ToLower(l)
		}
		if *fixed {
			matched = l == pattern
		} else {
			matched = strings.Contains(l, pattern)
		}
		if *invert {
			matched = !matched
		}
		if matched {
			matchedLinesCount++
			for i := 1; i < *before+1; i++ {
				if idx-i >= 0 {
					printLine(lines[idx-i], idx-i)
				}
			}
			printLine(line, idx)
			for i := 1; i < *after+1; i++ {
				if idx+i < len(lines) {
					printLine(lines[idx+i], idx+i)
				}
			}
		}
	}
	if *count {
		fmt.Println("количество найденных совпадений:", matchedLinesCount)
	}
}
