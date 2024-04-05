package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func reverseSlice(lines []string) {
	for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}
}

func uniqueLines(lines []string) []string {
	uniqueMap := make(map[string]bool)
	result := make([]string, 0)

	for _, line := range lines {
		if !uniqueMap[line] {
			uniqueMap[line] = true
			result = append(result, line)
		}
	}

	return result
}

func atoiSlice(s []string) ([]int, error) {
	var numbers []int
	for _, str := range s {
		n, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}
	return numbers, nil
}

func isSorted(data []string) bool {
	for i := 1; i < len(data); i++ {
		if data[i-1] > data[i] {
			return false
		}
	}
	return true
}

func printLines(data []string) {
	for _, ele := range data {
		fmt.Println(ele)
	}
}

func doSort(lines []string, reverse, unique, countSort, sorted bool, columns []int) {
	if sorted {
		if !isSorted(lines) {
			fmt.Println("Данные не отсортированы")
		} else {
			fmt.Println("Данные отсортированы")
		}
	}

	if len(columns) > 0 {
		var lineColumns [][]string
		for _, ele := range lines {
			cc := strings.Split(ele, " ")
			lineColumns = append(lineColumns, cc)
		}
		sort.Slice(lines, func(i, j int) bool {
			for _, c := range columns {
				if countSort {
					iv, err := strconv.Atoi(lineColumns[i][c])
					if err != nil {
						goto str
					}
					jv, err := strconv.Atoi(lineColumns[j][c])
					if err != nil {
						goto str
					}
					if iv == jv {
						continue
					}
					return iv < jv
				}
			str:
				if lineColumns[i][c] == lineColumns[j][c] {
					continue
				}
				return lineColumns[i][c] < lineColumns[j][c]
			}
			return false
		})
	} else {
		sort.Slice(lines, func(i, j int) bool {
			if countSort {
				iv, err := strconv.Atoi(lines[i])
				if err != nil {
					goto str
				}
				jv, err := strconv.Atoi(lines[j])
				if err != nil {
					goto str
				}
				return iv < jv
			}
		str:
			return lines[i] < lines[j]
		})
	}

	if reverse {
		reverseSlice(lines)
	}

	if unique {
		lines = uniqueLines(lines)
	}

	printLines(lines)
}

func main() {
	// Флаги командной строки
	reverse := flag.Bool("r", false, "сортировать в обратном порядке")
	unique := flag.Bool("u", false, "выводить только уникальные строки")
	countSort := flag.Bool("n", false, "сортировать по числовому значению")
	column := flag.String("k", "", "указать колонки для сортировки")
	sorted := flag.Bool("c", false, "проверять отсортированы ли данные")
	trim := flag.Bool("b", false, "игнорировать хвостовые пробелы")

	flag.Parse()

	var columns []int
	if *column != "" {
		for _, c := range strings.Split(*column, ",") {
			i, err := strconv.Atoi(c)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			if i <= 0 {
				fmt.Println("i должно быть > 1")
			}
			columns = append(columns, i-1)
		}
	}

	// Сканер для чтения ввода
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		if *trim {
			line = strings.TrimSpace(line)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка при чтении ввода:", err)
		os.Exit(1)
	}

	doSort(lines, *reverse, *unique, *countSort, *sorted, columns)
}
